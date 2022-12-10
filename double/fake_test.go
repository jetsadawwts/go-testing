package double

import "testing"

//Fake คือ การ implement logic บางอย่าง ทำให้เหมือนกับบางอย่าง เช่น Database api data


type FakeSearcher struct {}

func (fs FakeSearcher) Search(people []*Person, firstName string, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}
	
	return people[0]
}


func TestFindCallsSearchAndReturnsEmptyStringForNoPerson(t *testing.T) {
	
	fake := &FakeSearcher{} 
	phonebook := &Phonebook{}
	

	phone, _ := phonebook.Find(fake, "Jane", "Doe")

	if phone != "" {
		t.Errorf("Want '', Got '%s'", phone)
	}
}
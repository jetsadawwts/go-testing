package double

import "testing"

//Stub คือ เราต้องการ retrun value บางอย่าง คืนไป


type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName: lastName,
		Phone: ss.phone,
	}
}

func TestFindReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	phone, _ := phonebook.Find(StubSearcher{fakePhone}, "Jane", "Doe")

	if phone != fakePhone {
		t.Errorf("Want '%s', Got '%s'", fakePhone, phone)
	}
}
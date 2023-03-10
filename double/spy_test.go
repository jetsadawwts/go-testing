package double

import "testing"

//ความสามารถเหมือน Stub เเต้สามารถเขียน Logic เพิ่มได้

type SpySearcher struct {
	phone string
	searchWasCalled bool
	whatIsFirstName string
}

func (ss *SpySearcher) Search(people []*Person, firstName string, lastName string) *Person {
	ss.searchWasCalled = true
	ss.whatIsFirstName = firstName
	return &Person{
		FirstName: firstName,
		LastName: lastName,
		Phone: ss.phone,
	}
}

func TestFindCallSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	spy := &SpySearcher{phone: fakePhone} 

	phone, _ := phonebook.Find(spy, "Jane", "Doe")

	if !spy.searchWasCalled {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't.")
	}

	if spy.whatIsFirstName != "Jane" {
		t.Errorf("Expected to call 'Search' with 'Jane', as first name, but it wasn't.")
	}

	if phone != fakePhone {
		t.Errorf("Want '%s', Got '%s'", fakePhone, phone)
	}
}

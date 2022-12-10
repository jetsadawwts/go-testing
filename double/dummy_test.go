package double

import (
	"testing"
)

//Dummy คิอ ข้อมูล Object ทีไม่ต้องผ่านกระบวนการใดๆ สามารถใช้งานได้เลย

type DummySearcher struct {}

func (ds DummySearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{}
}

func TestFindItShouldReturnsErrorWhenFirstNameOrLastNameEmpty(t *testing.T) {
	phonebook := &Phonebook{}
	want := ErrMissingArgs

	_, got := phonebook.Find(DummySearcher{}, "", "")

	if got != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}
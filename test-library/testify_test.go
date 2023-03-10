package testlibrary

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

func TestSomething(t *testing.T) {
	t.Run("Run nil", func(t *testing.T) {
		pp := &Person{FirstName: "Jetsadawwts"}

		if(assert.NotNil(t, pp)) {
			assert.Equal(t, "Jetsadawwts", pp.FirstName)
		}
	})

	t.Run("equal", func(t *testing.T) {
		want := 555
		got := 555
		
		assert.Equal(t, want, got, "they should be not equal")
	})

	t.Run("not equal", func(t *testing.T) {
		want := 888
		got := 999
		
		assert.NotEqual(t, want, got, "they should be equal")
	})

	t.Run("nil", func(t *testing.T) {
		var p *Person

		assert.Nil(t, p)
	})

	t.Run("not nil", func(t *testing.T) {
		pp := &Person{FirstName: "Jetsadawwts", LastName: "Wongwit",Phone: "0643544759"}
		
		assert.NotNil(t, pp)
	})
}
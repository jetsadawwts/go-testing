package ticket

import "testing"

func TestTicketPrice(t *testing.T) {
	t.Run("should return 0 when age is 0", func(t *testing.T) {
		want := 0.0

		got := Price(0)

		if got != want {
			t.Errorf("Price(0) = %f; want %f", got, want)
		}

	})

	t.Run("Free Ticket when age under 3", func(t *testing.T) {
		want := 0.0

		got := Price(3)

		if got != want {
			t.Errorf("Price(3) = %f; want %f", got, want)
		}

	})

	t.Run("Ticket $15 when age at 4 year old", func(t *testing.T) {
		want := 15.0

		got := Price(4)

		if got != want {
			t.Errorf("Price(4) = %f; want %f", got, want)
		}

	})

	t.Run("Ticket $15 when age at 15 year old", func(t *testing.T) {
		want := 15.0

		got := Price(15)

		if got != want {
			t.Errorf("Price(15) = %f; want %f", got, want)
		}

	})

	t.Run("Ticket $30 when age at 16 year old", func(t *testing.T) {
		want := 30.0
		// age := 16
		var age uint = 16

		got := Price(age)

		if got != want {
			t.Errorf("Price(16) = %f; want %f", got, want)
		}

	})

	t.Run("Ticket $30 when age at 50 year old", func(t *testing.T) {
		want := 30.0
		// age := 50
		var age uint = 50
		got := Price(age)

		if got != want {
			t.Errorf("Price(50) = %f; want %f", got, want)
		}

	})

	t.Run("Ticket $5 when age at 51 year old", func(t *testing.T) {
		want := 5.0
		// age := 51
		var age uint = 51

		got := Price(age)

		if got != want {
			t.Errorf("Price(51) = %f; want %f", got, want)
		}

	})

	
}

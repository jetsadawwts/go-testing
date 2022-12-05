package ticket_test

import "testing"
import "github.com/jetsadawwts/ticket"


//Black box testing

func TestTicket(t *testing.T) {
	t.Run("should return 0 when age is 3", func(t *testing.T) {
			// Arrange
			want := 0.00
			var age uint = 3
			// Act
			got := ticket.Price(age)
			
			//Assert
			if got != want {
				t.Errorf("Price(3) = %f; want %f", got, want)
			}
		})

}
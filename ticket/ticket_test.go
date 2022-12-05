package ticket

import "testing"


type Case struct {
	name string
	age int
	want float64
}

func TestTicketPrice(t *testing.T) {
	// Arrange
	// tests := []struct {
	// 	name string
	// 	age int
	// 	want float64
	// } {
	// 	{name: "Free Ticket when age is 0",age: 0,want:  0.0},
	// 	{name: "Free Ticket when age under 3",age: 3,want:  0.0},
	// 	{name: "Ticket $15 when age at 4 year old",age: 4,want:  15.0},
	// 	{name: "Ticket $15 when age at 15 year old",age: 15,want:  15.0},
	// 	{name: "Ticket $30 when age at 16 year old",age: 16,want:  30.0},
	// 	{name: "Ticket $30 when age at 50 year old",age: 50,want:  30.0},
	// 	{name: "Ticket $5 when age at 51 year old",age: 51,want:  5.0},
	// }

	t1 := Case{name: "Free Ticket when age is 0",age: 0,want:  0.0}
	t2 := Case{name: "Free Ticket when age under 3",age: 3,want:  0.0}
	t3 := Case{name: "Ticket $15 when age at 4 year old",age: 4,want:  15.0}
	t4 := Case{name: "Ticket $15 when age at 15 year old",age: 15,want:  15.0}
	t5 := Case{name: "Ticket $30 when age at 16 year old",age: 16,want:  30.0}
	t6 := Case{name: "Ticket $30 when age at 50 year old",age: 50,want:  30.0}
	t7 := Case{name: "Ticket $5 when age at 51 year old",age: 51,want:  5.0}

	tests := []Case {
		t1,
		t2,
		t3,
		t4,
		t5,
		t6,
		t7,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got := Price(uint(tt.age))
			
			//Assert
			if got != tt.want {
				t.Errorf("Price(%d) = %f; want %f", tt.age, got, tt.want)
			}
		})
	}

}


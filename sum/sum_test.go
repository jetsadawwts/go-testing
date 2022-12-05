package sum

import "testing"

func TestSum(t *testing.T) {
	// t.Log("testing sum")
	// t.Error("testing sum")

	t.Run("should null parameter", func(t *testing.T) {
		// Arrange
		want := 0
		

		//Act
		got := sum([]int{}...)


		//Assert
		if got != want {
			t.Errorf("sum(1, 2) = %d; want 3", got)
		}
	})

	t.Run("should multi parameter", func(t *testing.T) {
		// Arrange
		want := 7
		

		//Act
		got := sum([]int{2, 3, 3, -1}...)


		//Assert
		if got != want {
			t.Error("Expected", 8, "Got", got)
		}
	})

	t.Run("should return 3 when 1 and 2", func(t *testing.T) {
		// Arrange
		want := 3
		

		//Act
		got := sum(1, 2)


		//Assert
		if got != want {
			t.Errorf("sum(1, 2) = %d; want 3", got)
		}
	})

	t.Run("should return -2 when -1 and -1", func(t *testing.T) {
		
		// Arrange
		want := -2

		//Act
		got := sum(-1, -1)

		//Assert
		if got != want {
			t.Errorf("sum(-1, -1) = %d; want -2", got)
		}
	})
	
}

// func TestSumShouldReturn1WhenInput1And0(t *testing.T) {
// 	got := sum(1, 0)

// 	if got != 1 {
// 		t.Errorf("sum(1, 0) = %d; want 1", got)
// 	}
// }

// func TestSumShouldReturnMinus2WhenInputMinus1AndMinus2(t *testing.T) {
// 	got := sum(-1, -1)

// 	if got != -2 {
// 		t.Errorf("sum(-1, -1) = %d; want -2", got)
// 	}
// }
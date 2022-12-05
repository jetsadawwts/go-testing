package teardown

import "testing"


//Setup / Teardown = กรณีที่มีการต่อ Database หรือ Api

func setup(t *testing.T) func() {
	t.Log("setup")

	return func() {
		t.Log("teardown")
	}
}


func TestTeardown(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	//t.Cleanup(teardown)

	//test...
	t.Run("subtest", func(t *testing.T) {
		t.Log("subtest")
	})
}
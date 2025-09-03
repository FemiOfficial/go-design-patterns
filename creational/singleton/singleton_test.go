package creational

import "testing"

func TestInstance(t *testing.T) {
	singleton := GetInstance()

	expectedCounter := singleton

	count := singleton.AddOne()
	if count != 1 {
		t.Error("Count is not 1")
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Error("Counter2 is not the same as expectedCounter")
	}

	count2 := counter2.AddOne()
	if (count2 != 2) {
		t.Error("Count is not 2")
	}
}
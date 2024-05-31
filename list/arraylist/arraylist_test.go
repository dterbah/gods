package arraylist

import (
	"fmt"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	list := New[int]()

	list.Add(1, 2, 3)
	expectedSize := 3

	if list.size != 3 {
		testError := fmt.Sprintf("Incorrect list size. Expected %d, Received %d", expectedSize, list.size)
		t.Error(testError)
	}

	for i, element := range list.elements[:list.size] {
		expectedValue := i + 1

		if expectedValue != element {
			testError := fmt.Sprintf("Incorrect value at the index %d. Expected %d, Recieved %d", i, expectedValue, element)
			t.Error(testError)
		}
	}
}

func TestArrayList_At(t *testing.T) {
	list := New[int]()

	list.Add(1, 2, 3)

	expectedValue := 2
	value, err := list.At(1)

	if err != nil {
		t.Error("The list shouldn't return any error")
	}

	if value != expectedValue {
		testError := fmt.Sprintf("Value at index %d not correct. Expected %d, Recieved %d", 1, expectedValue, value)
		t.Error(testError)
	}

	// Test a value out of bounds
	_, err = list.At(3)

	if err == nil {
		t.Error("The list should return an error because the index is out of bounds")
	}

	// Test negative index
	_, err = list.At(-1)

	if err == nil {
		t.Error("The list should return an error because the index is negative")
	}

}

func TestArrayList_Size(t *testing.T) {
	list := New[int]()

	list.Add(1, 2, 3)

	expectedSize := 3

	if expectedSize != list.Size() {
		testError := fmt.Sprintf("Incorrect list size. Expected %d, Received %d", expectedSize, list.Size())
		t.Error(testError)
	}
}

package algorithm

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	// Positive test case
	arr := []int{5, 3, 1, 4, 2}
	expected := []int{1, 2, 3, 4, 5}
	quickSort(0, len(arr)-1, arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("QuickSort failed, got: %v, want: %v", arr, expected)
	}

	// Negative test case
	arr = []int{5, 3, 1, 4, 2}
	expected = []int{1, 3, 2, 4, 5}
	quickSort(0, len(arr)-1, arr)
	if reflect.DeepEqual(arr, expected) {
		t.Errorf("QuickSort failed, got: %v, want: %v", arr, expected)
	}
}

func TestBubbleSort(t *testing.T) {
	// Positive test case
	arr := []int{5, 3, 1, 4, 2}
	expected := []int{1, 2, 3, 4, 5}
	result := bubbleSort(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BubbleSort failed, got: %v, want: %v", result, expected)
	}

	// Negative test case
	arr = []int{5, 3, 1, 4, 2}
	expected = []int{1, 3, 2, 4, 5}
	result = bubbleSort(arr)
	if reflect.DeepEqual(result, expected) {
		t.Errorf("BubbleSort failed, got: %v, want: %v", result, expected)
	}
}

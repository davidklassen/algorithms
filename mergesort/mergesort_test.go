package mergesort_test

import (
	"github.com/davidklassen/algorithms/mergesort"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 6, 4, 3, 8, 7, 3, 1, 2, 1, 9}
	exp := []int{1, 1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	mergesort.MergeSort(arr)
	if !reflect.DeepEqual(arr, exp) {
		t.Errorf("expected %+v to equal %+v", arr, exp)
	}
}

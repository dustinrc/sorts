package sorts_test

import (
	"testing"

	"github.com/dustinrc/sorts"
)

func TestInsertion(t *testing.T) {
	for _, tt := range integerCases {
		toSort := copyIntSlice(tt.given)
		isc := &sorts.IntSliceCounts{IntSlice: toSort}
		sorts.Insertion(isc)
		for i := range toSort {
			if toSort[i] != tt.expected[i] {
				// not fatal because we want to catch all failed cases in the test table
				t.Errorf("Insertion(%v): expected %v, actual %v", tt.given, tt.expected, toSort)
				break
			}
		}
		t.Log(isc)
	}
}

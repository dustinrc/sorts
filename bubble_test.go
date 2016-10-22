package sorts_test

import (
	"testing"

	"github.com/dustinrc/sorts"
)

func TestBubble(t *testing.T) {
	for _, tt := range integerCases {
		toSort := copyIntSlice(tt.given)
		isc := &sorts.IntSliceCounts{Slice: toSort}
		sorts.Bubble(isc)
		for i := range toSort {
			if toSort[i] != tt.expected[i] {
				// not fatal because we want to catch all failed cases in the test table
				t.Errorf("Bubble(%v): expected %v, actual %v", tt.given, tt.expected, toSort)
				break
			}
		}
		t.Log(isc)
	}
}

func TestBubble2(t *testing.T) {
	for _, tt := range integerCases {
		toSort := copyIntSlice(tt.given)
		isc := &sorts.IntSliceCounts{Slice: toSort}
		sorts.Bubble2(isc)
		for i := range toSort {
			if toSort[i] != tt.expected[i] {
				// not fatal because we want to catch all failed cases in the test table
				t.Errorf("Bubble2(%v): expected %v, actual %v", tt.given, tt.expected, toSort)
				break
			}
		}
		t.Log(isc)
	}
}

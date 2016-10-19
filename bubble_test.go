package sorts_test

import (
	"testing"

	"github.com/dustinrc/sorts"
)

type byInteger []int

func (n byInteger) Len() int           { return len(n) }
func (n byInteger) Less(i, j int) bool { return n[i] < n[j] }
func (n byInteger) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func TestBubble(t *testing.T) {
	for _, tt := range integerCases {
		toSort := make([]int, len(tt.given))
		copy(toSort, tt.given)
		sorts.Bubble(byInteger(toSort))
		for i := range toSort {
			if toSort[i] != tt.expected[i] {
				// not fatal because we want to catch all failed cases in the test table
				t.Errorf("Bubble(%v): expected %v, actual %v", tt.given, tt.expected, toSort)
				break
			}
		}
	}
}

func TestBubble2(t *testing.T) {
	for _, tt := range integerCases {
		toSort := make([]int, len(tt.given))
		copy(toSort, tt.given)
		sorts.Bubble2(byInteger(toSort))
		for i := range toSort {
			if toSort[i] != tt.expected[i] {
				// not fatal because we want to catch all failed cases in the test table
				t.Errorf("Bubble2(%v): expected %v, actual %v", tt.given, tt.expected, toSort)
				break
			}
		}
	}
}

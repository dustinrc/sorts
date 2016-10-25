package sorts

import "fmt"

// IntSlice implements the sorting methods for []int
type IntSlice []int

func (n IntSlice) Len() int                { return len(n) }
func (n IntSlice) Less(i, j int) bool      { return n[i] < n[j] }
func (n IntSlice) LessEqual(i, j int) bool { return n[i] <= n[j] }
func (n IntSlice) Swap(i, j int)           { n[i], n[j] = n[j], n[i] }

// IntSliceCounts implements the sorting methods for []int, also providing counts for comparisons and sorts
type IntSliceCounts struct {
	IntSlice
	comparisons, swaps int
}

func (nc *IntSliceCounts) Less(i, j int) bool {
	nc.comparisons++
	return nc.IntSlice.Less(i, j)
}

func (nc *IntSliceCounts) LessEqual(i, j int) bool {
	nc.comparisons++
	return nc.IntSlice.LessEqual(i, j)
}

func (nc *IntSliceCounts) Swap(i, j int) {
	nc.swaps++
	nc.IntSlice.Swap(i, j)
}

// String returns a string displaying the number of comparisons and swaps
func (nc IntSliceCounts) String() string {
	return fmt.Sprintf("Comparisons: %d; Swaps: %d", nc.comparisons, nc.swaps)
}

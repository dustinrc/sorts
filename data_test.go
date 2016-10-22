package sorts_test

var (
	randomInts          = []int{14, 3, 18, 8, 10, 15, 5, 19, 1, 11, 9, 13, 2, 17, 6, 16, 7, 0, 12, 4}
	nearlySortedInts    = []int{1, 2, 3, 4, 0, 6, 7, 8, 9, 5, 11, 12, 13, 14, 10, 16, 17, 18, 19, 15}
	reversedInts        = []int{19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fewUniqueInts       = []int{3, 0, 2, 1, 0, 2, 4, 3, 1, 4, 4, 1, 3, 0, 2, 4, 2, 1, 3, 0}
	sortedInts          = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fewUniqueSortedInts = []int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4}

	integerCases = []struct {
		given, expected []int
	}{
		{randomInts, sortedInts},
		{nearlySortedInts, sortedInts},
		{reversedInts, sortedInts},
		{sortedInts, sortedInts},
		{fewUniqueInts, fewUniqueSortedInts},
	}
)

func copyIntSlice(original []int) []int {
	copied := make([]int, len(original))
	copy(copied, original)
	return copied
}

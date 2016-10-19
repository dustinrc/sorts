package sorts

import "fmt"

type InsertionInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Insertion(in InsertionInterface) {
	var comparisons, swaps int
	for i := 1; i < in.Len(); i++ {
		for j := i; j > 0 && in.Less(j, j-1); j-- {
			comparisons++
			in.Swap(j, j-1)
			swaps++
		}
	}
	fmt.Println("Comparisons:", comparisons, "; Swaps:", swaps)
}

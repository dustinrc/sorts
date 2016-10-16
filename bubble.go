package sorts

import "fmt"

type BubbleInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Bubble(b BubbleInterface) {
	var comparisons, swaps int
	for i := 0; i < b.Len(); i++ {
		swapped := false
		for j := b.Len() - 1; j > i; j-- {
			comparisons++
			if b.Less(j, j-1) {
				swaps++
				b.Swap(j, j-1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Println("Comparisons:", comparisons, "; Swaps:", swaps)
}

func Bubble2(b BubbleInterface) {
	var comparisons, swaps int
	for n := b.Len(); n != 0; {
		newN := 0
		for i := 1; i <= n-1; i++ {
			comparisons++
			if b.Less(i, i-1) {
				swaps++
				b.Swap(i, i-1)
				newN = i
			}
		}
		n = newN
	}
	fmt.Println("Comparisons:", comparisons, "; Swaps:", swaps)
}

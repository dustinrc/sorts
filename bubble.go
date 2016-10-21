package sorts

type BubbleInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Bubble(b BubbleInterface) {
	for i := 0; i < b.Len(); i++ {
		swapped := false
		for j := b.Len() - 1; j > i; j-- {
			if b.Less(j, j-1) {
				b.Swap(j, j-1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func Bubble2(b BubbleInterface) {
	for n := b.Len(); n != 0; {
		newN := 0
		for i := 1; i <= n-1; i++ {
			if b.Less(i, i-1) {
				b.Swap(i, i-1)
				newN = i
			}
		}
		n = newN
	}
}

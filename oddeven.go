package sorts

type OddEvenInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func OddEven(oe OddEvenInterface) {
	for sorted := false; !sorted; {
		sorted = true
		for i := 1; i < oe.Len()-1; i += 2 {
			if oe.Less(i+1, i) {
				oe.Swap(i+1, i)
				sorted = false
			}
		}
		for i := 0; i < oe.Len()-1; i += 2 {
			if oe.Less(i+1, i) {
				oe.Swap(i+1, i)
				sorted = false
			}
		}
	}
}

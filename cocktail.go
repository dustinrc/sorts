package sorts

type CocktailInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Cocktail(c CocktailInterface) {
	start, end := 0, c.Len()-2
	var swapped bool
	for start <= end {
		newStart := end
		newEnd := start
		for i := start; i <= end; i++ {
			if c.Less(i+1, i) {
				c.Swap(i+1, i)
				swapped = true
				newEnd = i
			}
		}
		if !swapped {
			break
		}
		end = newEnd - 1
		for i := end; i >= start; i-- {
			if c.Less(i+1, i) {
				c.Swap(i+1, i)
				swapped = true
				newStart = i
			}
		}
		if !swapped {
			break
		}
		start = newStart + 1
	}
}

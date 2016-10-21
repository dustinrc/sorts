package sorts

type InsertionInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Insertion(in InsertionInterface) {
	for i := 1; i < in.Len(); i++ {
		for j := i; j > 0 && in.Less(j, j-1); j-- {
			in.Swap(j, j-1)
		}
	}
}

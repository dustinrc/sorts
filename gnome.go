package sorts

type GnomeInterface interface {
	Len() int
	LessEqual(i, j int) bool
	Swap(i, j int)
}

func Gnome(g GnomeInterface) {
	var pos int
	for pos < g.Len() {
		if pos == 0 || g.LessEqual(pos-1, pos) {
			pos++
		} else {
			g.Swap(pos-1, pos)
			pos--
		}
	}
}

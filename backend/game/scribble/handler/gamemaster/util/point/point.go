package point

type Handler interface {
	GetPoints() int
	ResetPoints()
}

type handler struct {
	WordPointBonuses []int
	WordPointIndex   int
}

func Get() *handler {
	return &handler{
		WordPointBonuses: []int{50, 20, 10},
		WordPointIndex:   0,
	}
}

func (h *handler) GetPoints() int {
	retPoints := h.WordPointBonuses[h.WordPointIndex]
	if h.WordPointIndex < len(h.WordPointBonuses)-1 {
		h.WordPointIndex++
	}
	return retPoints
}

// ResetPoints will reset how the point system hands out points
func (h *handler) ResetPoints() {
	h.WordPointIndex = 0
}

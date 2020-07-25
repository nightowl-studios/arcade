package point

type Handler struct {
	WordPointBonuses []int
	WordPointIndex   int
}

func Get() *Handler {
	return &Handler{
		WordPointBonuses: []int{50, 20, 10},
		WordPointIndex:   0,
	}
}

func (h *Handler) GetPoints() int {
	retPoints := h.WordPointBonuses[h.WordPointIndex]
	if h.WordPointIndex < len(h.WordPointBonuses)-1 {
		h.WordPointIndex++
	}
	return retPoints
}

func (h *Handler) ResetPoints() {
	h.WordPointIndex = 0
}

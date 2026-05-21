package emergence

import (
	"nospace/quantum"
)

// InteractionHistory хранит счётчики взаимодействий между объектами.
type InteractionHistory struct {
	Counts map[string]map[string]int
}

// NewInteractionHistory создаёт новую историю.
func NewInteractionHistory() *InteractionHistory {
	return &InteractionHistory{Counts: make(map[string]map[string]int)}
}

// Record регистрирует взаимодействие obj1 и obj2.
func (h *InteractionHistory) Record(obj1, obj2 *quantum.QuantumObject) {
	n1, n2 := obj1.Name, obj2.Name
	if h.Counts[n1] == nil {
		h.Counts[n1] = make(map[string]int)
	}
	if h.Counts[n2] == nil {
		h.Counts[n2] = make(map[string]int)
	}
	h.Counts[n1][n2]++
	h.Counts[n2][n1]++
}

// PerceivedDistance возвращает воспринимаемое расстояние между объектами,
// обратно пропорциональное числу взаимодействий: d = 1 / (1 + count).
func (h *InteractionHistory) PerceivedDistance(obj1, obj2 *quantum.QuantumObject) float64 {
	n1, n2 := obj1.Name, obj2.Name
	count := 0
	if h.Counts[n1] != nil {
		count = h.Counts[n1][n2]
	}
	return 1.0 / (1.0 + float64(count))
}

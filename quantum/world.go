package quantum

import (
	"fmt"
	"math/rand"
)

// QuantumObject хранит распределение вероятностей координат,
// флаг коллапса и финальную координату.
type QuantumObject struct {
	Name        string
	CoordDist   map[[2]int]float64 // (x,y) -> вес (вероятность до нормировки)
	IsCollapsed bool
	FinalCoord  [2]int
}

// NewQuantumObject создаёт новый квантовый объект с заданным распределением.
func NewQuantumObject(name string, dist map[[2]int]float64) *QuantumObject {
	return &QuantumObject{
		Name:      name,
		CoordDist: dist,
	}
}

// NormalizeDistribution нормирует распределение так, чтобы сумма вероятностей стала 1.
func (q *QuantumObject) NormalizeDistribution() {
	total := 0.0
	for _, w := range q.CoordDist {
		total += w
	}
	if total > 0 {
		for k, w := range q.CoordDist {
			q.CoordDist[k] = w / total
		}
	}
}

// Collapse выполняет коллапс волновой функции: выбирает случайную координату
// согласно распределению вероятностей. Если объект уже коллапсирован, ничего не делает.
func (q *QuantumObject) Collapse() {
	if q.IsCollapsed {
		return
	}
	q.NormalizeDistribution()
	r := rand.Float64()
	cumulative := 0.0
	for coord, prob := range q.CoordDist {
		cumulative += prob
		if r <= cumulative {
			q.FinalCoord = coord
			q.IsCollapsed = true
			// заменяем распределение на дельта-функцию
			q.CoordDist = map[[2]int]float64{coord: 1.0}
			break
		}
	}
}

func (q *QuantumObject) String() string {
	if q.IsCollapsed {
		return fmt.Sprintf("<%s collapsed at (%d, %d)>",
			q.Name, q.FinalCoord[0], q.FinalCoord[1])
	}
	return fmt.Sprintf("<%s in superposition (uncollapsed)>", q.Name)
}

// World — дискретное пространство размером Width×Height, содержащее объекты.
type World struct {
	Width   int
	Height  int
	Objects []*QuantumObject
}

// NewWorld создаёт новый мир заданного размера.
func NewWorld(width, height int) *World {
	return &World{Width: width, Height: height}
}

// AddQuantumObject добавляет объект в мир.
func (w *World) AddQuantumObject(obj *QuantumObject) {
	w.Objects = append(w.Objects, obj)
}

// MeasureInteraction выполняет взаимодействие между двумя объектами.
// Взаимодействие происходит только в точках совпадения координат.
func (w *World) MeasureInteraction(obj1, obj2 *QuantumObject) {
	if obj1.IsCollapsed && obj2.IsCollapsed {
		return
	}
	obj1.NormalizeDistribution()
	obj2.NormalizeDistribution()

	newDist1 := make(map[[2]int]float64)
	newDist2 := make(map[[2]int]float64)

	for c1, p1 := range obj1.CoordDist {
		for c2, p2 := range obj2.CoordDist {
			if c1 == c2 && p1 > 0 && p2 > 0 {
				w := p1 * p2
				if w > 0 {
					newDist1[c1] += w
					newDist2[c2] += w
				}
			}
		}
	}

	// Если нет общих точек, взаимодействие не происходит.
	if len(newDist1) == 0 || len(newDist2) == 0 {
		return
	}

	obj1.CoordDist = newDist1
	obj2.CoordDist = newDist2
	obj1.Collapse()
	obj2.Collapse()
}

// CollapseAll коллапсирует все объекты в мире.
func (w *World) CollapseAll() {
	for _, obj := range w.Objects {
		obj.Collapse()
	}
}

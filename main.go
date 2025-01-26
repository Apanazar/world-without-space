package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// QuantumObject хранит распределение вероятностей координат объекта,
// а также признак, был ли коллапс (isCollapsed) и финальную координату (finalCoord).
type QuantumObject struct {
	Name        string
	CoordDist   map[[2]int]float64 // распределение: (x,y) -> "вес" (вероятность до нормировки)
	IsCollapsed bool
	FinalCoord  [2]int
}

// NewQuantumObject конструктор, принимает имя и словарь координат.
func NewQuantumObject(name string, dist map[[2]int]float64) *QuantumObject {
	return &QuantumObject{
		Name:      name,
		CoordDist: dist,
	}
}

// NormalizeDistribution нормирует распределение до суммарной вероятности 1.0
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

// Collapse выбирает случайную координату по весам и «фиксирует» объект.
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
			break
		}
	}
}

func (q *QuantumObject) String() string {
	if q.IsCollapsed {
		return fmt.Sprintf("<%s collapsed at (%d, %d)>",
			q.Name, q.FinalCoord[0], q.FinalCoord[1])
	} else {
		return fmt.Sprintf("<%s in superposition (uncollapsed)>", q.Name)
	}
}

// World хранит «пространство» (допустим, дискретное в width×height) и набор объектов.
type World struct {
	Width   int
	Height  int
	Objects []*QuantumObject
}

// NewWorld создаёт мир с указанными размерами.
func NewWorld(width, height int) *World {
	return &World{
		Width:  width,
		Height: height,
	}
}

// AddQuantumObject добавляет объект в мир.
func (w *World) AddQuantumObject(obj *QuantumObject) {
	w.Objects = append(w.Objects, obj)
}

// MeasureInteraction упрощённо моделирует «совместное измерение» (например, "человек увидел дерево").
//
// Идея: если объекты ещё не коллапсированы, мы пересекаем их распределения так,
// чтобы они «могли встретиться». В данном упрощении считаем, что встретиться можно
// только в одной и той же точке (x,y). Затем делаем коллапс обоих в эту точку.
func (w *World) MeasureInteraction(obj1, obj2 *QuantumObject) {
	// Если оба уже коллапсированы - ничего не делаем
	if obj1.IsCollapsed && obj2.IsCollapsed {
		return
	}

	// Нормируем исходные распределения
	obj1.NormalizeDistribution()
	obj2.NormalizeDistribution()

	newDist1 := make(map[[2]int]float64)
	newDist2 := make(map[[2]int]float64)

	// Перебираем все пары (x1,y1) из obj1 и (x2,y2) из obj2
	// "Визуальный контакт" => (x1,y1) == (x2,y2)
	// Совместный вес = p1 * p2
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

	// Если пересечение пустое, значит нет координат, где они могут «взаимодействовать».
	// Тогда в данной модели просто ничего не делаем.
	if len(newDist1) == 0 || len(newDist2) == 0 {
		return
	}

	// Обновляем распределения
	obj1.CoordDist = newDist1
	obj2.CoordDist = newDist2

	// Коллапсируем оба
	obj1.Collapse()
	obj2.Collapse()
}

// CollapseAll коллапсирует все объекты.
func (w *World) CollapseAll() {
	for _, obj := range w.Objects {
		obj.Collapse()
	}
}

// uniformDistribution создаёт равномерное распределение по всем (x,y).
func uniformDistribution(width, height int) map[[2]int]float64 {
	dist := make(map[[2]int]float64)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dist[[2]int{x, y}] = 1.0
		}
	}
	return dist
}

// gaussFactor возвращает exp(-0.5*r^2), где r^2 = (dx^2 + dy^2).
func gaussFactor(x, y, cx, cy int) float64 {
	dx := float64(x - cx)
	dy := float64(y - cy)
	distSq := dx*dx + dy*dy
	return math.Exp(-0.5 * distSq)
}

func exampleScenario() {
	// Создаем мир 10×10 (дискретная сетка).
	world := NewWorld(10, 10)

	// Создаем равномерные распределения для всех.
	treeDist := uniformDistribution(world.Width, world.Height)
	johnDist := uniformDistribution(world.Width, world.Height)
	observerDist := uniformDistribution(world.Width, world.Height)

	// Создаем сами объекты.
	tree := NewQuantumObject("Tree", treeDist)
	john := NewQuantumObject("John", johnDist)
	observer := NewQuantumObject("Observer1", observerDist)

	// "Посадка дерева":
	// Усилим амплитуды дерева вокруг (3,3).
	for coord, weight := range tree.CoordDist {
		x, y := coord[0], coord[1]
		tree.CoordDist[coord] = weight * gaussFactor(x, y, 3, 3)
	}
	// Усилим амплитуды Джона вокруг (2,3).
	for coord, weight := range john.CoordDist {
		x, y := coord[0], coord[1]
		john.CoordDist[coord] = weight * gaussFactor(x, y, 2, 3)
	}

	// Добавляем объекты в мир.
	world.AddQuantumObject(tree)
	world.AddQuantumObject(john)
	world.AddQuantumObject(observer)

	// "Джон сажает дерево" => совместное измерение (John, Tree)
	world.MeasureInteraction(john, tree)

	// "Прошло несколько лет, человек (observer) видел дерево 3 раза":
	for i := 0; i < 3; i++ {
		world.MeasureInteraction(observer, tree)
	}

	// Коллапсируем все объекты
	world.CollapseAll()

	for _, obj := range world.Objects {
		fmt.Println(obj)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	exampleScenario()
}

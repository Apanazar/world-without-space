package nonlocal

import (
	"math"
	"nospace/quantum"
)

// InfoDistance — функция информационной близости координат.
// Использует гауссово ядро: exp(-((dx²+dy²)/(2σ²))) с σ=1.
func InfoDistance(c1, c2 [2]int, width, height int) float64 {
	dx := float64(c1[0] - c2[0])
	dy := float64(c1[1] - c2[1])
	return math.Exp(-(dx*dx + dy*dy) / 2.0)
}

// MeasureNonlocalInteraction выполняет нелокальное измерение:
// вес пары = p1 * p2 * InfoDistance(c1,c2).
// После расчёта выбирается пара с максимальным весом, и объекты коллапсируют.
func MeasureNonlocalInteraction(obj1, obj2 *quantum.QuantumObject, width, height int) {
	if obj1.IsCollapsed && obj2.IsCollapsed {
		return
	}
	obj1.NormalizeDistribution()
	obj2.NormalizeDistribution()

	type pair struct{ c1, c2 [2]int }
	weights := make(map[pair]float64)

	for c1, p1 := range obj1.CoordDist {
		for c2, p2 := range obj2.CoordDist {
			w := p1 * p2 * InfoDistance(c1, c2, width, height)
			if w > 0 {
				weights[pair{c1, c2}] = w
			}
		}
	}

	if len(weights) == 0 {
		return
	}

	// Находим пару с максимальным весом
	var best pair
	maxW := 0.0
	for p, w := range weights {
		if w > maxW {
			maxW = w
			best = p
		}
	}

	// Коллапсируем объекты в выбранные координаты
	obj1.CoordDist = map[[2]int]float64{best.c1: 1.0}
	obj1.IsCollapsed = true
	obj1.FinalCoord = best.c1

	obj2.CoordDist = map[[2]int]float64{best.c2: 1.0}
	obj2.IsCollapsed = true
	obj2.FinalCoord = best.c2
}

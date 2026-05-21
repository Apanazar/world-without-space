package dynamics

import "nospace/quantum"

// Diffuse выполняет один шаг диффузии распределения вероятностей.
// Параметр rate ∈ [0,1] определяет долю вероятности, передаваемую соседям.
func Diffuse(obj *quantum.QuantumObject, width, height int, rate float64) {
	if obj.IsCollapsed {
		return
	}
	obj.NormalizeDistribution()
	newDist := make(map[[2]int]float64)

	for coord, prob := range obj.CoordDist {
		x, y := coord[0], coord[1]
		// остаток в текущей клетке
		stay := prob * (1 - rate)
		newDist[coord] += stay

		// равномерная передача четырём ортогональным соседям
		share := prob * rate / 4.0
		neighbors := [][2]int{
			{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1},
		}
		for _, nb := range neighbors {
			nx, ny := nb[0], nb[1]
			if nx >= 0 && nx < width && ny >= 0 && ny < height {
				newDist[[2]int{nx, ny}] += share
			} else {
				// возврат вышедшей за границу доли в исходную клетку
				newDist[coord] += share
			}
		}
	}
	obj.CoordDist = newDist
}

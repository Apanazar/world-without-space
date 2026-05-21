package interactions

import "nospace/quantum"

// PlantSeed моделирует "посадку семечка": распределение planter прибавляется
// к распределению seed с последующей нормировкой.
func PlantSeed(planter, seed *quantum.QuantumObject) {
	planter.NormalizeDistribution()
	seed.NormalizeDistribution()

	// добавляем вес planter в распределение seed
	for coord, p := range planter.CoordDist {
		if existing, ok := seed.CoordDist[coord]; ok {
			seed.CoordDist[coord] = existing + p
		} else {
			seed.CoordDist[coord] = p
		}
	}
	seed.NormalizeDistribution()
}

// Remember моделирует "вспоминание": observer копирует распределение target.
func Remember(observer, target *quantum.QuantumObject) {
	observer.CoordDist = make(map[[2]int]float64)
	for k, v := range target.CoordDist {
		observer.CoordDist[k] = v
	}
}

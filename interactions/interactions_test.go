package interactions

import (
	"nospace/quantum"
	"testing"
)

func TestPlantSeed(t *testing.T) {
	john := quantum.NewQuantumObject("John", map[[2]int]float64{{1, 1}: 1.0})
	seed := quantum.NewQuantumObject("Seed", map[[2]int]float64{{2, 2}: 1.0})
	PlantSeed(john, seed)

	// seed должен иметь вес и в (1,1), и в (2,2)
	if seed.CoordDist[[2]int{1, 1}] <= 0 {
		t.Error("seed should have weight at (1,1)")
	}
	if seed.CoordDist[[2]int{2, 2}] <= 0 {
		t.Error("seed should retain weight at (2,2)")
	}
}

func TestRemember(t *testing.T) {
	tree := quantum.NewQuantumObject("Tree", map[[2]int]float64{{3, 3}: 0.7, {4, 4}: 0.3})
	observer := quantum.NewQuantumObject("Observer", map[[2]int]float64{{0, 0}: 1.0})
	Remember(observer, tree)

	if observer.CoordDist[[2]int{3, 3}] != 0.7 {
		t.Error("observer should copy tree's distribution")
	}
}

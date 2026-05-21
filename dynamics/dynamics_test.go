package dynamics

import (
	"nospace/quantum"
	"testing"
)

func TestDiffuse(t *testing.T) {
	obj := quantum.NewQuantumObject("X", map[[2]int]float64{{2, 2}: 1.0})
	Diffuse(obj, 5, 5, 0.5)

	// центральная вероятность должна уменьшиться
	if obj.CoordDist[[2]int{2, 2}] == 1.0 {
		t.Error("center probability should have decreased")
	}

	total := 0.0
	for _, v := range obj.CoordDist {
		total += v
	}
	if total < 0.99 || total > 1.01 {
		t.Errorf("probability not conserved, total=%f", total)
	}
}

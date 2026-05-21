package quantum

import (
	"testing"
)

func TestNormalizationAfterInteraction(t *testing.T) {
	world := NewWorld(5, 5)
	obj1 := NewQuantumObject("Obj1", map[[2]int]float64{{0, 0}: 1, {1, 1}: 1})
	obj2 := NewQuantumObject("Obj2", map[[2]int]float64{{0, 0}: 1, {2, 2}: 1})
	world.AddQuantumObject(obj1)
	world.AddQuantumObject(obj2)
	world.MeasureInteraction(obj1, obj2)

	total1, total2 := 0.0, 0.0
	for _, w := range obj1.CoordDist {
		total1 += w
	}
	for _, w := range obj2.CoordDist {
		total2 += w
	}
	if total1 != 1.0 || total2 != 1.0 {
		t.Errorf("normalization failed: obj1=%f, obj2=%f", total1, total2)
	}
}

func TestInteractionDependenceOnDistance(t *testing.T) {
	world := NewWorld(5, 5)
	obj1 := NewQuantumObject("Obj1", map[[2]int]float64{{0, 0}: 1})
	obj2 := NewQuantumObject("Obj2", map[[2]int]float64{{4, 4}: 1})
	world.AddQuantumObject(obj1)
	world.AddQuantumObject(obj2)
	world.MeasureInteraction(obj1, obj2)
	if obj1.IsCollapsed || obj2.IsCollapsed {
		t.Error("objects should not collapse when coordinates do not match")
	}

	obj3 := NewQuantumObject("Obj3", map[[2]int]float64{{2, 2}: 1})
	obj4 := NewQuantumObject("Obj4", map[[2]int]float64{{2, 2}: 1})
	world.AddQuantumObject(obj3)
	world.AddQuantumObject(obj4)
	world.MeasureInteraction(obj3, obj4)
	if !obj3.IsCollapsed || !obj4.IsCollapsed {
		t.Error("objects should collapse when coordinates match")
	}
}

func TestSystemStability(t *testing.T) {
	world := NewWorld(5, 5)
	obj1 := NewQuantumObject("Obj1", map[[2]int]float64{{1, 1}: 1})
	obj2 := NewQuantumObject("Obj2", map[[2]int]float64{{1, 1}: 1})
	world.AddQuantumObject(obj1)
	world.AddQuantumObject(obj2)
	world.MeasureInteraction(obj1, obj2) // оба коллапсируют
	coord1 := obj1.FinalCoord

	obj3 := NewQuantumObject("Obj3", map[[2]int]float64{{1, 1}: 1})
	world.AddQuantumObject(obj3)
	world.MeasureInteraction(obj1, obj3)
	if obj1.FinalCoord != coord1 {
		t.Error("collapsed object should not change coordinate")
	}
}

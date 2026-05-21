package nonlocal

import (
	"nospace/quantum"
	"testing"
)

func TestNonlocalCollapse(t *testing.T) {
	obj1 := quantum.NewQuantumObject("A", map[[2]int]float64{{0, 0}: 1})
	obj2 := quantum.NewQuantumObject("B", map[[2]int]float64{{4, 4}: 1})
	MeasureNonlocalInteraction(obj1, obj2, 5, 5)
	if !obj1.IsCollapsed || !obj2.IsCollapsed {
		t.Error("nonlocal interaction should collapse both objects")
	}
	if obj1.FinalCoord != [2]int{0, 0} {
		t.Errorf("expected (0,0), got %v", obj1.FinalCoord)
	}
	if obj2.FinalCoord != [2]int{4, 4} {
		t.Errorf("expected (4,4), got %v", obj2.FinalCoord)
	}
}

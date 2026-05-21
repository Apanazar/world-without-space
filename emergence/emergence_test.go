package emergence

import (
	"nospace/quantum"
	"testing"
)

func TestPerceivedDistance(t *testing.T) {
	hist := NewInteractionHistory()
	a := quantum.NewQuantumObject("A", nil)
	b := quantum.NewQuantumObject("B", nil)

	d0 := hist.PerceivedDistance(a, b)
	if d0 != 1.0 {
		t.Errorf("expected initial distance 1.0, got %f", d0)
	}

	hist.Record(a, b)
	d1 := hist.PerceivedDistance(a, b)
	if d1 >= 1.0 {
		t.Error("distance should decrease after interaction")
	}

	hist.Record(a, b)
	d2 := hist.PerceivedDistance(a, b)
	if d2 >= d1 {
		t.Error("distance should further decrease")
	}
}

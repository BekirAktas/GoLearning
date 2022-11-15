package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	expected := Repeat("burak", 6)
	repeated := "burakburakburakburakburakburak"

	if expected != repeated {
		t.Errorf("expected %s but got %s", expected, repeated)
	}
}
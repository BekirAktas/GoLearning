package integers

import "testing"


func TestSum(t *testing.T) {
	want := 4
	got := Sum(2,2)

	if want != got {
		t.Errorf("expected %d but got %d", want,got)		
	}
}
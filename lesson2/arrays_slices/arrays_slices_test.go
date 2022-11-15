package arraysSlices

import (
	"reflect"
	"testing"
)

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	bekirArr := []int{2,3,5,10};
	expected := Sum(bekirArr)
	got := 20;

	if expected != got {
		t.Errorf("expected %d but got %d",expected, got)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3,9}

	checkSums(t,got,want)
}


func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2,5,3}, []int{5,9,1,2})
		want := []int{10,12}

		checkSums(t,got,want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5,9,1,2})
		want := []int{0,12}

		checkSums(t,got,want)
	})
}
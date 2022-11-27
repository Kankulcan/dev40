package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection off any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}

	})
	
}

func TestSumall(t *testing.T){
	got := sumall([]int{1,2}, []int{0,9})
	want := []int{3,9}

	if !reflect.DeepEqual(got, want)  {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumalltaisl(t *testing.T){

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := sumalltails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
		
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := sumalltails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t,got, want)
	})



}

package main

import (
	"reflect"
	"testing"
)

func TestIfRandom(t *testing.T) {
	a, b := newNumbers(10, 100), newNumbers(10, 100)

	if reflect.DeepEqual(a, b) == true {
		a.print()
		b.print()
		t.Errorf("Numbers are not random, check seed")
	}
}

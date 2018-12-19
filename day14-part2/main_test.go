package main

import (
	"testing"
)

func TestChocolateChart01245(t *testing.T) {
	cc := NewChocolateChart("01245", []int{3, 7})
	actual := cc.Generations()
	expected := 5
	if actual != expected {
		t.Errorf("Generations were incorrect.\nActual:   %d\nExpected: %d", actual, expected)
	}
}

func TestChocolateChart51589(t *testing.T) {
	cc := NewChocolateChart("51589", []int{3, 7})
	actual := cc.Generations()
	expected := 9
	if actual != expected {
		t.Errorf("Generations were incorrect.\nActual:   %d\nExpected: %d", actual, expected)
	}
}

func TestChocolateChart92510(t *testing.T) {
	cc := NewChocolateChart("92510", []int{3, 7})
	actual := cc.Generations()
	expected := 18
	if actual != expected {
		t.Errorf("Generations were incorrect.\nActual:   %d\nExpected: %d", actual, expected)
	}
}

func TestChocolateChart59414(t *testing.T) {
	cc := NewChocolateChart("59414", []int{3, 7})
	actual := cc.Generations()
	expected := 2018
	if actual != expected {
		t.Errorf("Generations were incorrect.\nActual:   %d\nExpected: %d", actual, expected)
	}
}

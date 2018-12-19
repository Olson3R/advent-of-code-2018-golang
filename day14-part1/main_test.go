package main

import (
	"testing"
)

func TestChocolateChart9_10(t *testing.T) {
	cc := NewChocolateChart(9, 10, []int{3, 7})
	scores := cc.Scores()
	expected := "5158916779"
	if scores != expected {
		t.Errorf("Scores were incorrect.\nActual:   %s\nExpected: %s", scores, expected)
	}
}

func TestChocolateChart5_10(t *testing.T) {
	cc := NewChocolateChart(5, 10, []int{3, 7})
	scores := cc.Scores()
	expected := "0124515891"
	if scores != expected {
		t.Errorf("Scores were incorrect.\nActual:   %s\nExpected: %s", scores, expected)
	}
}

func TestChocolateChart18_10(t *testing.T) {
	cc := NewChocolateChart(18, 10, []int{3, 7})
	scores := cc.Scores()
	expected := "9251071085"
	if scores != expected {
		t.Errorf("Scores were incorrect.\nActual:   %s\nExpected: %s", scores, expected)
	}
}

func TestChocolateChart2018_10(t *testing.T) {
	cc := NewChocolateChart(2018, 10, []int{3, 7})
	scores := cc.Scores()
	expected := "5941429882"
	if scores != expected {
		t.Errorf("Scores were incorrect.\nActual:   %s\nExpected: %s", scores, expected)
	}
}

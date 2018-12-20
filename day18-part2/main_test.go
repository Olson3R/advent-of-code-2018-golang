package main

import (
	"io/ioutil"
	"testing"
)

func TestLoggingZone(t *testing.T) {
	data, err := ioutil.ReadFile("../input/d18-input-test.txt")
	check(err)

	lz := LoggingZone{}
	lz.ParseMap(string(data))
	for i := 1; i <= 10; i++ {
		lz.Next()
	}
	actual := lz.Value()
	expected := 1147
	if actual != expected {
		t.Errorf("Value was incorrect.\nActual:   %d\nExpected: %d", actual, expected)
	}
}

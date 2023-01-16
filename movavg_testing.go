package main

import (
	"bytes"
	"testing"
)

func TestCalculateAvg(t *testing.T) {
	input := bytes.NewBufferString(`{"timestamp":"2022-01-01 00:00:00.000000", "duration": 10}
    {"timestamp":"2022-01-01 00:01:00.000000", "duration": 20}
    {"timestamp":"2022-01-01 00:02:00.000000", "duration": 30}`)
	expectedOutput := `{"date":"2022-01-01 00:00:00","average_delivery_time":10}
    {"date":"2022-01-01 00:01:00","average_delivery_time":15}
    {"date":"2022-01-01 00:02:00","average_delivery_time":20}`
	var buf bytes.Buffer
	windowSize := uint(1)
	if err := calculateAvg(input, &buf, windowSize); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if buf.String() != expectedOutput {
		t.Errorf("Expected: %s, got: %s", expectedOutput, buf.String())
	}
}

package main

import (
	"os"
	"testing"
)

func TestEnv_reads_environment_variables(t *testing.T) {
	expected := "testvalue"
	os.Setenv("testkey", "testvalue")

	actual := env("testkey")
	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

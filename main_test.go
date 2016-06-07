package main

import (
	"flag"
	"log"
	"os"
	"testing"
)

var testVer = "1.0"

/* Setup & Teardown */
func TestMain(m *testing.M) {
	flag.Parse()

	// Get current VERSION
	curVer := os.Getenv("VERSION")

	// Set VERSION to test value
	if err := os.Setenv("VERSION", testVer); err != nil {
		log.Fatal(err)
	}

	runCode := m.Run()

	// Set VERSION to its original value
	if err := os.Setenv("VERSION", curVer); err != nil {
		log.Fatal(err)
	}

	os.Exit(runCode)
}

func TestGetVersion(t *testing.T) {
	v := GetVersion()
	if v != "1.0" {
		t.Errorf("expected 1.0; got %s", v)
	}
}

func TestGreetingMessage(t *testing.T) {
	m := GreetingMessage()
	if m != "Running version 1.0" {
		t.Errorf("expected 'Running version 1.0'; got '%s'", m)
	}
}

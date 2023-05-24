package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	code := m.Run()

	postRun()

	os.Exit(code)
}

/*
Copyright © 2025 mozduh
*/

package utils

import (
	"testing"
)

func TestCheckProjectDir(t *testing.T) {
	temp := t.TempDir() // makes a temp directory and cleans up after the test

	found := CheckForJitDir(temp)
	if found {
		t.Fatalf("Found a directory when it shouldn't have.")
	}
}

func TestInitProjectDir(t *testing.T) {
	temp := t.TempDir() // makes a temp directory and cleans up after the test

	found := CheckForJitDir(temp)
	if found {
		t.Fatalf("Found a directory when it shouldn't have.")
	}

	err := InitJitProject(temp)
	if err != nil {
		t.Fatal("got error when trying to init jit project directory:", err)
	}
}

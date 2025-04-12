/*
Copyright © 2025 mozduh
*/

package cmd

import (
	"bytes"
	"testing"
)

func TestInitCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	temp := t.TempDir() // makes a temp directory and cleans up after the test
	initCmd.SetOut(buf)
	initCmd.SetArgs([]string{"--dir", temp})

	err := initCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

}

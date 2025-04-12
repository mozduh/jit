/*
Copyright © 2025 mozduh
*/

package task

import (
	"bytes"
	"testing"
)

func TestTaskCreateCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	TaskCreateCmd.SetOut(buf)

	err := TaskCreateCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

}

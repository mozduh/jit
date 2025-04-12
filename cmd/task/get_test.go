/*
Copyright © 2025 mozduh
*/

package task

import (
	"bytes"
	"testing"
)

func TestTaskGetCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	TaskGetCmd.SetOut(buf)
	TaskGetCmd.SetArgs([]string{"123"})

	err := TaskGetCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

}

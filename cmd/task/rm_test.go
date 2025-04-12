/*
Copyright © 2025 mozduh
*/

package task

import (
	"bytes"
	"testing"
)

func TestTaskRMCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	TaskRmCmd.SetOut(buf)

	err := TaskRmCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

}

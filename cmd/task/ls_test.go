/*
Copyright © 2025 mozduh
*/

package task

import (
	"bytes"
	"testing"
)

func TestTaskLSCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	TaskLsCmd.SetOut(buf)

	err := TaskLsCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

}

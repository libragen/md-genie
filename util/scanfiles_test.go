package util

import "testing"

func TestLsMdHacknews(t *testing.T) {
	if files, err := LsArchivesMdFiles("archives"); err != nil {
		t.Log(files)
	} else {
		t.Error(err)
	}
}

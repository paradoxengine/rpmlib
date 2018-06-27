package rpmlib

import (
	"testing"
	//	"fmt"
)

func TestReadPackage(t *testing.T) {
	ts, err := NewTransactionSet()
	if err != nil {
		t.Fatal("err")
	}

	// Note our test package has an untursted signature
	h, err := ts.ReadPackageFile("testdata/sample.rpm", false)
	if err != nil {
		t.Fatal(err)
	}

	_, err = h.GetString(RPMTAG_SUMMARY)
	if err != nil {
		t.Fatal(err)
	}

	_, err = h.GetStringArray(RPMTAG_BASENAMES)
	if err != nil {
		t.Fatal(err)
	}

	h.Free()
	ts.Free()
}

// IMPORTANT: this test will only succeed if you have a local RPM DB.
func TestTransactionIterator(t *testing.T) {
	ts, err := NewTransactionSet()
	if err != nil {
		t.Fatal(err)
	}

	iter, err := ts.SequencialIterator()
	if err != nil {
		t.Fatal(err)
	}

	h, err := iter.Next()
	if err != nil {
		t.Fatal(err)
	}

	_, err = h.GetString(RPMTAG_SUMMARY)
	if err != nil {
		t.Fatal(err)
	}

	iter.Free()

	ts.Free()
}

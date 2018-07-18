// Copyright 2018 necomeshi, All Rights Reserved.
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

	h, err := ts.ReadPackageFile("testdata/sample.rpm")
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


func TestDatabase(t *testing.T) {

	db, err := OpenDatabase()
	if err != nil {
		t.Fatal(err)
	}

	iter, err := db.SequencialIterator()
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

	err = db.Close()
	if err != nil {
		t.Fatal(err)
	}

}

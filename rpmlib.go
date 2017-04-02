package rpmlib

/*
#cgo LDFLAGS: -lrpm -lrpmio -lpopt

#include <stdlib.h>
#include <stdint.h>
#include <fcntl.h>
#include <rpm/header.h>
#include <rpm/rpmlib.h>
#include <rpm/rpmdb.h>
#include <rpm/rpmts.h>
#include <rpm/rpmtag.h>
*/
import "C"
import "unsafe"
import "fmt"
import "io"

//////////////////////////////////////////
// Header
//////////////////////////////////////////
type Header struct {
	rpmheader C.Header
}

func (h *Header) GetString(tag RpmTagVal) (s string, err error) {

	cstring := C.headerGetString(h.rpmheader, C.rpmTag(tag))

	if cstring == nil {
		err = fmt.Errorf("C.headerGetString: Cannot get tag value")
		return
	}

	s = C.GoString(cstring)

	// cstring is 'const char*'. Do not C.free it.

	return
}

func (h *Header) GetStringArray(tag RpmTagVal) (values []string, err error) {

	var sa C.struct_rpmtd_s

	if C.headerGet(h.rpmheader, C.rpmTag(tag), &sa, C.HEADERGET_MINMEM) == 0 {
		err = fmt.Errorf("C.headerGet: Cannot get tag value")
		return
	}

	if C.rpmtdCount(&sa) == 0 {
		C.rpmtdFreeData(&sa)
		err = fmt.Errorf("C.rpmtdCount: Tag not contain data")
		return
	}

	for {
		v := C.rpmtdNextString(&sa)
		if v == nil {
			break
		}

		gv := C.GoString(v)
		values = append(values, gv)
	}

	C.rpmtdFreeData(&sa)

	return
}

func (h *Header) GetUint16(tag RpmTagVal) (value uint16, err error) {

	var sa C.struct_rpmtd_s

	if C.headerGet(h.rpmheader, C.rpmTag(tag), &sa, C.HEADERGET_MINMEM) == 0 {
		err = fmt.Errorf("C.headerGet: Cannot get tag value")
		return
	}

	v := C.rpmtdGetUint16(&sa)
	if v == nil {
		C.rpmtdFreeData(&sa)
		err = fmt.Errorf("C.rpmtdGetUint16: Tag not contain data")
		return
	}

	value = uint16(*v)
	C.free(unsafe.Pointer(v))

	C.rpmtdFreeData(&sa)
	return
}

func (h *Header) GetUint32(tag RpmTagVal) (value uint32, err error) {

	var sa C.struct_rpmtd_s

	if C.headerGet(h.rpmheader, C.rpmTag(tag), &sa, C.HEADERGET_MINMEM) == 0 {
		err = fmt.Errorf("C.headerGet: Cannot get tag value")
		return
	}

	v := C.rpmtdGetUint32(&sa)
	if v == nil {
		C.rpmtdFreeData(&sa)
		err = fmt.Errorf("C.rpmtdGetUint32: Tag not contain data")
		return
	}

	value = uint32(*v)
	C.free(unsafe.Pointer(v))

	C.rpmtdFreeData(&sa)
	return
}

func (h *Header) GetUint32Array(tag RpmTagVal) (values []uint32, err error) {

	var sa C.struct_rpmtd_s

	if C.headerGet(h.rpmheader, C.rpmTag(tag), &sa, C.HEADERGET_MINMEM) == 0 {
		err = fmt.Errorf("C.headerGet: Cannot get tag value")
		return
	}

	if C.rpmtdCount(&sa) == 0 {
		C.rpmtdFreeData(&sa)
		err = fmt.Errorf("C.rpmtdCount: Tag not contain data")
		return
	}

	for {
		v := C.rpmtdNextUint32(&sa)
		if v == nil {
			break
		}

		gv := uint32(*v)
		values = append(values, gv)

		C.free(unsafe.Pointer(v))
	}

	C.rpmtdFreeData(&sa)

	return
}


func (h *Header) GetUint64(tag RpmTagVal) (value uint64, err error) {

	var sa C.struct_rpmtd_s

	if C.headerGet(h.rpmheader, C.rpmTag(tag), &sa, C.HEADERGET_MINMEM) == 0 {
		err = fmt.Errorf("C.headerGet: Cannot get tag value")
		return
	}

	v := C.rpmtdGetUint64(&sa)
	if v == nil {
		C.rpmtdFreeData(&sa)
		err = fmt.Errorf("C.rpmtdGetUint64: Tag not contain data")
		return
	}

	value = uint64(*v)
	C.free(unsafe.Pointer(v))

	C.rpmtdFreeData(&sa)
	return
}

func (h *Header) GetUint64Array(tag RpmTagVal) (values []uint64, err error) {

	var sa C.struct_rpmtd_s

	if C.headerGet(h.rpmheader, C.rpmTag(tag), &sa, C.HEADERGET_MINMEM) == 0 {
		err = fmt.Errorf("C.headerGet: Cannot get tag value")
		return
	}

	if C.rpmtdCount(&sa) == 0 {
		C.rpmtdFreeData(&sa)
		err = fmt.Errorf("C.rpmtdCount: Tag not contain data")
		return
	}

	for {
		v := C.rpmtdNextUint64(&sa)
		if v == nil {
			break
		}

		gv := uint64(*v)
		values = append(values, gv)

		C.free(unsafe.Pointer(v))
	}

	C.rpmtdFreeData(&sa)

	return
}

func (h *Header) IsSource() (isSource bool) {
	if C.headerIsSource(h.rpmheader) == 1 {
		isSource = true
	} else {
		isSource = false
	}
	return
}

func (h *Header) Free() {
	C.headerFree(h.rpmheader)
}

//////////////////////////////////////////
// Database iterator
//////////////////////////////////////////
type Iterator struct {
	mi C.rpmdbMatchIterator
}

func (iter *Iterator) Next() (h *Header, err error) {
	h = new(Header)
	rpmheader := C.rpmdbNextIterator(iter.mi)

	if rpmheader == nil {
		err = io.EOF
	}

	h.rpmheader = C.headerLink(rpmheader)

	return
}

func (iter *Iterator) Free() {
	C.rpmdbFreeIterator(iter.mi)
}

//////////////////////////////////////////
// Database
//////////////////////////////////////////
type Database struct {
	rpmdb C.rpmdb
}

func OpenDatabase() (db *Database, err error) {
	db = new(Database)

	prefix := C.CString("")

	if C.rpmdbOpen(prefix, &db.rpmdb, C.O_RDWR, 0666) != 0 {
		err = fmt.Errorf("C.rpmdbOpen: Cannot open database")
	}

	C.free(unsafe.Pointer(prefix))

	return
}

func (db *Database) SequencialIterator() (iter *Iterator, err error) {
	iter = new(Iterator)

	iter.mi = C.rpmdbInitIterator(db.rpmdb, C.RPMDBI_PACKAGES, nil, 0)
	if iter.mi == nil {
		err = fmt.Errorf("Cannot get iterator")
	}

	return
}

func (db *Database) Close() (err error) {
	if C.rpmdbClose(db.rpmdb) != 0 {
		err = fmt.Errorf("C.rpmdbClose: Cannot close database")
	}
	return
}

//////////////////////////////////////////
// Transaction
//////////////////////////////////////////

type TransactionSet struct {
	ts C.rpmts
}

func NewTransactionSet() (ts *TransactionSet, err error) {

	ts = new(TransactionSet)

	ts.ts = C.rpmtsCreate()

	C.rpmtsSetRootDir(ts.ts, C.CString("/"))

	return
}

func (ts *TransactionSet) SequencialIterator() (iter *Iterator, err error) {
	iter = new(Iterator)

	iter.mi = C.rpmtsInitIterator(ts.ts, C.RPMDBI_PACKAGES, nil, 0)
	if iter.mi == nil {
		err = fmt.Errorf("C.rpmtsInitIterator: Cannot get iterator")
	}

	return
}

func (ts *TransactionSet) ReadPackageFile(name string) (header *Header, err error) {
	cname := C.CString(name)
	cmode := C.CString("r.ufdio")

	fd := C.Fopen(cname, cmode)

	if fd == nil {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(cmode))
		return nil, fmt.Errorf("C.Fopen: Error")
	}

	header = new(Header)
	ret := C.rpmReadPackageFile(ts.ts, fd, cname, &header.rpmheader)
	if ret != C.RPMRC_OK {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(cmode))
		return nil, fmt.Errorf("C.rpmReadPackageFile: Error")
	}

	C.Fclose(fd)
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(cmode))

	return
}

func (ts *TransactionSet) Free() {
	// Opened database will be closed in here
	C.rpmtsFree(ts.ts)
}

func init() {
	C.rpmReadConfigFiles(nil, nil)
}
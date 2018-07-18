rpmlib
===

RPM API language bindings for Go language

## Description
RPM API language bindings for Go language.

## Installation

This package depends on RPM C language API libraries. 
On RedHat/CentOS environment, you can get RPM development 
libraries by hitting following command.

```bash
$ yum -y install rpm-devel
$ rpm -qi rpm-devel
```

Please check the rpm-devel package version. 
Currently, my language bindings support RPM version 4.8.0 only!

Then, import to your project

```bash
$ export GOROOT=<Your Go language tools path>
$ export GOPATH=<Your Go project directory>
$ go get github.com/necomeshi/rpmlib
```

## Usages

Currently, following APIs are available. 

* Searcing and getting installed package information

```c
package main

import (
	"os"
	"fmt"
	"io"
	"github.com/necomeshi/rpmlib"
)

func main() {
	ts, err := rpmlib.NewTransactionSet()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot get transaction set: %s\n", err)
		os.Exit(1)
	}

	defer ts.Free()

	iter, err := ts.SequencialIterator()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot get iterator: %s\n", err)
		os.Exit(1)
	}

	defer iter.Free()

	for {
		h, itr_err := iter.Next()

		if itr_err == io.EOF {
			break	
		}
		
		if itr_err != nil {
			fmt.Fprintf(os.Stderr, "Cannot get next iterator: %s\n", itr_err)
			break
		}

		defer h.Free()

		name, h_err := h.GetString(rpmlib.RPMTAG_NAME)	
		if h_err != nil {
			fmt.Fprintf(os.Stderr, "Cannot get name from rpm header: %s\n", h_err)
			break
		}

		fmt.Println(name)

	}
}
```

* Get a package file information

```c
package main

import (
	"os"
	"fmt"
	"github.com/necomeshi/rpmlib"
)


func main() {

	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stderr, "No package name is specified")
		os.Exit(1)
	}

	ts, err := rpmlib.NewTransactionSet()	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot get transaction set: %s\n", err)
		os.Exit(1)
	}

	defer ts.Free()

	h, err := ts.ReadPackageFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot get package rpm header: %s\n", err)
		os.Exit(1)
	}

	defer h.Free()

	name, err := h.GetStringArray(rpmlib.RPMTAG_NAME)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot get package name: %s\n", err)
	}

	fmt.Println(name)

}
```

## Author
Necomeshi

## Copyright

Copyright (c) 2018 necomeshi https://github.com/necomesh

see ./LICENSE

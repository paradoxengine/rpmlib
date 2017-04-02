gorpm
===

RPM package management system API bindings for Go language

## Description
RPM commmand line utility implemented in Go language.
This 'gorpm' and other commands aim to get RPM package's 
information at such as Windows, debian,which does not have yum/rpm command. 
So this command is not for package management on RHEL or CentOS.

This is my first 'decent' Go language project for practicing.
If you found any mistakes or bugs in command, please tell me.

## Installation

This package depents on RPM C language API libraries.
In RedHat/CentOS enviroment, you can get RPM libraries

```bash
$ yum -y install rpm-devel
$ rpm -qi rpm-devel
```

Then, import your project

```bash
$ export GOROOT=<Your Go language tools path>
$ export GOPATH=<Your Go project directory>
$ go get github.com/necomeshi/rpmlib
```

## Usages

Currently, API for searching installed packages and for getting 
package information are available.

* Searcing installed package

* Get a package information

## FAQ
1. Why xx option has not been implemented ? When will you implement it ?
 Sometime when I need it. Or sometime when others give me an early Xmas present.


## Author
Necomeshi

go-xslt
=====

[![Go Reference](https://pkg.go.dev/badge/github.com/wamuir/go-xslt.svg)](https://pkg.go.dev/github.com/wamuir/go-xslt)
[![Build Status](https://travis-ci.com/wamuir/go-xslt.svg?branch=master)](https://travis-ci.com/wamuir/go-xslt)
[![codecov](https://codecov.io/gh/wamuir/go-xslt/branch/master/graph/badge.svg)](https://codecov.io/gh/wamuir/go-xslt)
[![Go Report Card](https://goreportcard.com/badge/github.com/wamuir/go-xslt)](https://goreportcard.com/report/github.com/wamuir/go-xslt)

# Description

`go-xslt` is a Go module that performs basic XSLT 1.0 transformations via Libxslt.

This fork has been extended to take optional URLs as parameters and reads the files from there. Quick and Dirty

# Installation

You'll need the development libraries for libxml2 and libxslt, along with those for liblzma and zlib.  Install these via your package manager. For instance, if using `apt` then:

    sudo apt install libxml2-dev libxslt1-dev liblzma-dev zlib1g-dev

This module can be installed with the `go get` command:

    go get -u github.com/uvwild/go-xslt


# Usage

```go

  // style is an XSLT 1.0 stylesheet, as []byte.
  xs, err := xslt.NewStylesheet(style)
  if err != nil {
      panic(err)
  }
  defer xs.Close()

  // doc is an XML document to be transformed and res is the result of
  // the XSL transformation, both as []byte. 
  res, err := xs.Transform(doc)
  if err != nil {
      panic(err)
  }

```

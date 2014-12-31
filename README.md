[![GoDoc](https://godoc.org/github.com/lunixbochs/matter?status.svg)](http://godoc.org/github.com/lunixbochs/matter)
[![Build Status](https://travis-ci.org/lunixbochs/matter.svg?branch=master)](https://travis-ci.org/lunixbochs/matter)

matter
===

A Go library for easily reading/writing files with frontmatter.

What is frontmatter?
---

Files with YAML frontmatter attached usually look like this:

```
---
key: value
key: value
---
regular file contents.
```

Installation
---

    go get github.com/lunixbochs/matter

Example
---

```Go
package main

import (
    "fmt"
    "log"

    "github.com/lunixbochs/matter"
)

type TestStruct struct {
    Example string
    KeyTwo  string
}

var dataTest = []byte("Example data.")
var structTest = &TestStruct{
    Example: "Eggs",
    KeyTwo:  "Ham",
}

func main() {
    err := matter.WriteYAML("test.yml", structTest, dataTest, 0600)
    if err != nil {
        log.Fatal(err)
    }

    front := &TestStruct{}
    data, err := matter.ReadYAML("test.yml", front)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("  frontmatter:", front)
    fmt.Println("  data:", string(data))
}
```

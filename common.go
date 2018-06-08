package main

import (
	"os"
	"path"
)

var baseDir, err = os.Getwd()
var tmplDir = path.Join(baseDir, "templates")


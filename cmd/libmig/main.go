package main

import (
	"github.com/mazrean/libmig"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(libmig.Analyzer) }

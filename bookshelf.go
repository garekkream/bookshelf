package main

import (
	parser "gopkg.in/alecthomas/kingpin.v2"
)

var (
	ver  = "none"
	date = "note"
)

func main() {
	parser.Version(date + ver)
	parser.Parse()
}

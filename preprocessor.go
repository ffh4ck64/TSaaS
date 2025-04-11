package main

import "sync"

var (
	templateStore = make(map[string]template)
	tmplMutex     sync.RWMutex
	templateDir   = "/data"
	documentsDir  = "/generates"
)

func put_imag() {}

func put_ra() {}

func preprocessor() {}

package main

import (
	"encoding/json"
	"strings"
)

func printVarnames(varnames []string) {
	switch *varnamesFormat {
	case "json":
		jsonBytes, err := json.Marshal(varnames)
		app.FatalIfError(err, "")
		print(string(jsonBytes))
	case "list":
		print(strings.Join(varnames, "\n"))
	default:
		app.FatalUsage("unknown format %s", *varnamesFormat)
	}
}

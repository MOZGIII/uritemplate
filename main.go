package main

import (
	"fmt"
	"os"

	"github.com/yosida95/uritemplate"
	"gopkg.in/alecthomas/kingpin.v2"
)

type kingpinArgCallable interface {
	Arg(name, help string) *kingpin.ArgClause
}

func templateArg(c kingpinArgCallable) *string {
	return c.Arg("template", "A URI template.").Required().String()
}

var (
	app     = kingpin.New("uritemplate", "A command-line RFC6570 (URI Template) expander.")
	newline = app.Flag("newline", "Print newline at the end of output (use --no-newline to ommit newlines).").Default("true").Bool()

	expand         = app.Command("expand", "Print the expanded URI template.")
	expandTemplate = templateArg(expand)
	expandStrings  = expand.Flag("var", "Variables to substitute as key-value pairs.").PlaceHolder("KEY=VAL").Short('v').StringMap()
	expandJSONs    = JSONValues(expand.Flag("json", "Variables to substitute as JSON.").Short('j'))

	varnames         = app.Command("varnames", "Print variable names found in the template.")
	varnamesTemplate = templateArg(varnames)
	varnamesFormat   = varnames.Flag("format", "Format to use.").Default("list").Short('f').Enum("list", "json")

	regexp         = app.Command("regexp", "Print a regexp that matches the template.")
	regexpTemplate = templateArg(regexp)
)

func print(s ...interface{}) {
	if *newline {
		fmt.Println(s...)
	} else {
		fmt.Print(s...)
	}
}

func addStrings(values uritemplate.Values) error {
	for key, val := range *expandStrings {
		if values.Get(key) != nil {
			return fmt.Errorf("variable '%v' specified twice", key)
		}
		values.Set(key, uritemplate.String(val))
	}
	return nil
}

func addJSONs(values uritemplate.Values) error {
	for key, val := range *expandJSONs {
		if values.Get(key) != nil {
			return fmt.Errorf("variable '%v' specified twice", key)
		}
		values.Set(key, val)
	}
	return nil
}

func main() {
	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch command {
	case expand.FullCommand():
		template, err := uritemplate.New(*expandTemplate)
		app.FatalIfError(err, "")

		values := make(uritemplate.Values)
		app.FatalIfError(addStrings(values), "")
		app.FatalIfError(addJSONs(values), "")

		uri, err := template.Expand(values)
		app.FatalIfError(err, "")

		print(uri)
	case varnames.FullCommand():
		template, err := uritemplate.New(*varnamesTemplate)
		app.FatalIfError(err, "")
		printVarnames(template.Varnames())
	case regexp.FullCommand():
		template, err := uritemplate.New(*regexpTemplate)
		app.FatalIfError(err, "")
		print(template.Regexp())
	}
}

// version.go : supporting library to display the version information
// for an application
//
// Copyright 2023 Simon Rowe <simon@wiremoons.com>.
// All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package lib

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"runtime"
	"strings"
	"text/template"
)

// Version function collects details of the program being run and
// displays it on stdout
func Version(appname string, appversion string) {
	// set up a caser to create titlecase words
	caser := cases.Title(language.English)
	// define a template for display on screen with place holders for data
	const appInfoTmpl = `
'{{.appname}}' is version {{.appversion}}.
Built with Go Complier '{{.compiler}}' using Golang version '{{.version}}'.
Executing on operating system '{{.operSys}}' with CPU architecture '{{.runArch}}'.
Source: https://github.com/wiremoons/go-check/
MIT License. Copyright © 2023 Simon Rowe.
`
	// build a map with keys set to match the template names used
	// and the data fields to be used in the template as values
	verData := map[string]interface{}{
		"appname":    appname,
		"appversion": appversion,
		"operSys":    caser.String(runtime.GOOS),
		"runArch":    strings.ToUpper(runtime.GOARCH),
		"compiler":   runtime.Compiler,
		"version":    runtime.Version(),
	}
	// check and build the template so the data field values are added
	// and the final output is displayed. Check for any error, and
	// abort if one is found.
	t := template.Must(template.New("appinf").Parse(appInfoTmpl))

	if err := t.Execute(os.Stdout, verData); err != nil {
		fmt.Printf("FATAL ERROR: in function 'Version()' when building template with err: %v", err)
	}
}

/*
 * Compilation: go build
 * Execution: ./autoescape > autoescape.html
 *
 * suppress auto-escaping behavior for fields that contain trusted HTML data by
 * using the named string type template.HTML instead of string. Similar named
 * types exist for trusted Java Script, CSS, and URLs.
 *
 * $ ./autoescape > autoescape.html
 * $ cat autoescape.html
 * <p>A: &lt;b&gt;Hello!&lt;/b&gt;</p><p>B: <b>Hello!</b></p>
 * $ /usr/bin/google-chrome autoescape.html
 * A: <b>Hello!</b>
 *
 * B: Hello!
 *
 */
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

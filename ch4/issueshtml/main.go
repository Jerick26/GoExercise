/*
 * Compilation: go build
 * Execution: ./issueshtml repo:golang/go commenter:gopherbot json encoder >issues.html
 *
 * More elaborate format. Separate the format from the code more
 * completely. generate a html format file.
 *
 * $ ./issueshtml repo:golang/go commenter:gopherbot json encoder >issues.html
 * $ /usr/bin/google-chrome issues.html
 *
 * 18 issues
 *
 * #	State	User	Title
 * 7872	open	extemporalgenome	encoding/json: Encoder internally buffers full output
 * 6492	closed	kisielk	encoding/json: Add Encoder.Indent
 * 11939	open	joeshaw	proposal: encoding/json, encoding/xml: support zero values of structs with omitempty
 * 14749	closed	cyberphone	encoding/json: add Encoder option to control escaping behavior
 * 12529	closed	allaud	encoding/json: support int map keys in encoding/json
 * 20206	open	markdryan	encoding/base64: encoding is slow
 * 4146	open	bradfitz	reflect: MakeInterface
 * 12146	closed	augustoroman	encoding/json: allow non-string map keys that implement encoding.TextMarshaler/TextUnmarshaler
 * 15808	open	randall77	cmd/compile: loads/constants not lifted out of loop
 * 11940	closed	lukescott	proposal: encoding/json: Encode channel as array
 * 17255	closed	xoba	encoding/json docs reference an unknown entity "DisableHTMLEscaping"
 * 7767	closed	gopherbot	encoding/json: Encoder adds trailing newlines
 * 4474	closed	gopherbot	encoding/json: json encoder fails for embedded non-struct fields
 * 17704	closed	dsnet	encoding/json: regression in handling of nil RawMessage
 * 14493	closed	niemeyer	encoding/json: RawMessage marshals as base64
 * 4747	closed	gopherbot	encoding/json Added tag options to ignore fields of struct for encoder/decoder separately
 * 6901	closed	lukescott	encoding/json, encoding/xml: option to treat unknown fields as an error
 * 11508	closed	josharian	cmd/go: trace http viewer: "http: multiple response.WriteHeader calls"
 */
package main

import (
	"html/template"
	"log"
	"os"

	"../github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-algin:left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

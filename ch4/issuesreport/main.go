/*
 * Compilation: go build
 * Execution: ./issuesreport repo:golang/go is:open json decoder
 *
 * More elaborate format. Separate the format from the code more
 * completely. generate a text format report.
 *
 * $ ./issuesreport repo:golang/go is:open json decoder
 * 19 issues:
 * ----------------
 * Number: 20567
 * User: schmichael
 * Title: json: reduce allocations by Decoder for \uXXXX
 * Age: 53 days
 * ----------------
 * Number: 20814
 * User: epelc
 * Title: proposal: image: support decoder instance with specific formats
 * Age: 29 days
 * ----------------
 * ......
 *
 */
package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"../github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func noMust() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

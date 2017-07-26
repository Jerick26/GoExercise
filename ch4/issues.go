/******************************************************************************
 *  Compilation:  go build issues.go
 *  Execution:    ./issues repo:golang/go is:open json decoder
 *
 *  The SearchIssues func tion makes an HTTP request and decodes the result as
 *  JSON.
 *
 *  % ./issues repo:golang/go is:open json decoder
 *  19 issues:
 *  #20567 schmichae json: reduce allocations by Decoder for \uXXXX
 *  #20814     epelc proposal: image: support decoder instance with specific
 *  #11046     kurin encoding/json: Decoder internally buffers full input
 *  #15314    okdave proposal: some way to reject unknown fields in encoding
 *  #21026  shurcooL encoding/json: indentation of raw strings in examples l
 *  #16212 josharian encoding/json: do all reflect work before decoding
 *  #12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
 *  #7872  extempora encoding/json: Encoder internally buffers full output
 *  #5901        rsc encoding/json: allow override type marshaling
 *  #17609 nathanjsw encoding/json: ambiguous fields are marshalled
 *  #14750 cyberphon encoding/json: parser ignores the case of member names
 *  #20528  jvshahid net/http: connection reuse does not work happily with n
 *  #19469 chengzhic runtime: temporary object is not garbage collected
 *  #7213  davechene cmd/compile: escape analysis oddity
 *  #20754       rsc encoding/xml: unmarshal only processes first XML elemen
 *  #15808 randall77 cmd/compile: loads/constants not lifted out of loop
 *  #20206 markdryan encoding/base64: encoding is slow
 *  #19029    cynecx runtime: fatal error: sweep increased allocation count
 *  #17244       adg proposal: decide policy for sub-repositories
 *
 *  (details: https://api.github.com/search/issues?q=repo:golang/go+is:open+json+decoder)
 *
 ******************************************************************************/
package main

import (
	"./github"
	"fmt"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

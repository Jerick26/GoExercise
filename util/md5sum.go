package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

func Md5Base64(input string) string {
	m := md5.New()
	io.WriteString(m, input)
	md5Result := m.Sum(nil)
	return base64.StdEncoding.EncodeToString(md5Result)
}

func main() {
	raw := "2f7caaf310da3dcb24bacdc7944456210Suggestionsdfxalldisk"
	fmt.Println(Md5Base64(raw))
}

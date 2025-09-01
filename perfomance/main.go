// xml2csv.go
package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type User struct {
	ID    string `xml:"id"`
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	// 見出し
	if err := w.Write([]string{"id", "name", "email"}); err != nil {
		fmt.Fprintln(os.Stderr, "header write:", err)
		os.Exit(1)
	}

	// <user>要素を順次拾ってCSVへ
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "token error:", err)
			os.Exit(1)
		}
		start, ok := tok.(xml.StartElement)
		if !ok || start.Name.Local != "user" {
			continue
		}
		var u User
		if err := dec.DecodeElement(&u, &start); err != nil {
			fmt.Fprintln(os.Stderr, "decode error:", err)
			os.Exit(1)
		}
		if err := w.Write([]string{u.ID, u.Name, u.Email}); err != nil {
			fmt.Fprintln(os.Stderr, "csv write:", err)
			os.Exit(1)
		}
	}
}

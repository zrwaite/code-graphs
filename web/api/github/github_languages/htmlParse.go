package github_languages

import (
	"fmt"
	"strings"

	"github.com/zrwaite/github-graphs/utils"
	"golang.org/x/net/html"
)

func GetSvg(text string) (data string) {
	tkn := html.NewTokenizer(strings.NewReader(text))

	var inSvg bool

	svg := ""

	for {

		tt := tkn.Next()
		t := tkn.Token()

		switch {
		case tt == html.ErrorToken:
			fmt.Println("error")
			return ""

		case tt == html.StartTagToken:

			if t.Data == "svg" {
				inSvg = true
				// utils.SetAttribute(&t, "xmlns", "http://www.w3.org/2000/svg")
				// utils.SetAttribute(&t, "width", "843")
				// utils.SetAttribute(&t, "height", "148")
				// utils.SetAttribute(&t, "viewBox", "0 0 843 148")
			}

			// if class="header" then change inside text
			header, found := utils.GetAttribute(&t, "class")
			if found && header == "header" {
				t.Data = "text"
			}

		case tt == html.TextToken:
		case tt == html.EndTagToken:
			if inSvg && t.Data == "svg" {
				svg += t.String()
				return svg
			}
		}
		if inSvg {
			svg += t.String()
		}

	}
}

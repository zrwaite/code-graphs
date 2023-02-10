package github_languages

import (
	"fmt"
	"strings"

	"github.com/zrwaite/github-graphs/utils"
	"golang.org/x/net/html"
)

func GetSvg(text string) (data string) {
	tkn := html.NewTokenizer(strings.NewReader(text))

	var svgDepth int
	var inHeader bool

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
				svgDepth += 1
			}

			header, found := utils.GetAttribute(&t, "class")
			if found && header == "header" {
				inHeader = true
			}

		case tt == html.TextToken:
			if inHeader {
				t.Data = "Top Github Languages"
				inHeader = false
			}
		case tt == html.EndTagToken:
			if t.Data == "svg" {
				if svgDepth == 1 {
					svg += t.String()
					return svg
				} else {
					svgDepth -= 1
				}
			}
		}
		if svgDepth > 0 {
			svg += t.String()
		}

	}
}

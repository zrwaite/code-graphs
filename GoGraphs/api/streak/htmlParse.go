package streak

import (
	"fmt"
	"strings"

	"github.com/zrwaite/github-graphs/utils"
	"golang.org/x/net/html"
)

func GetSvg(text string) (data string) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var inSvg bool
	var firstSvgTag bool
	var firstGroupFound bool

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
				firstSvgTag = true
				utils.SetAttribute(&t, "xmlns", "http://www.w3.org/2000/svg")
				utils.SetAttribute(&t, "width", "843")
				utils.SetAttribute(&t, "height", "148")
				utils.SetAttribute(&t, "viewBox", "0 0 843 148")
			}

			if !firstGroupFound && t.Data == "g" {
				firstGroupFound = true
				utils.SetAttribute(&t, "transform", "translate(25, 30)")
				utils.SetAttribute(&t, "fill", "black")
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
		if firstSvgTag {
			svg += GetStyleSheet()
			svg += GetBackground()
			firstSvgTag = false
		}
	}
}

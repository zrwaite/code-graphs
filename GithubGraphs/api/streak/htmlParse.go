package streak

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

//	func getAttribute(n *html.Token, key string) (string, bool) {
//		for _, attr := range n.Attr {
//			if attr.Key == key {
//				return attr.Val, true
//			}
//		}
//		return "", false
//	}

func removeAttribute(t *html.Token, index int) {
	t.Attr = append(t.Attr[:index], t.Attr[index+1:]...)
}

func setAttribute(t *html.Token, key string, value string) {
	index := -1
	for i, attr := range t.Attr {
		if attr.Key == key {
			attr.Val = value
			index = i
		}
	}
	if index != -1 {
		removeAttribute(t, index)
	}

	t.Attr = append(t.Attr, html.Attribute{
		Key: key,
		Val: value,
	})
}

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
				setAttribute(&t, "xmlns", "http://www.w3.org/2000/svg")
				setAttribute(&t, "width", "843")
				setAttribute(&t, "height", "148")
				setAttribute(&t, "viewBox", "0 0 843 148")
			}

			if !firstGroupFound && t.Data == "g" {
				firstGroupFound = true
				setAttribute(&t, "transform", "translate(25, 30)")
				setAttribute(&t, "fill", "black")
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

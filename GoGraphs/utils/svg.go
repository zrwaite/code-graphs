package utils

import "golang.org/x/net/html"

func GetAttribute(n *html.Token, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func RemoveAttribute(t *html.Token, index int) {
	t.Attr = append(t.Attr[:index], t.Attr[index+1:]...)
}

func SetAttribute(t *html.Token, key string, value string) {
	index := -1
	for i, attr := range t.Attr {
		if attr.Key == key {
			attr.Val = value
			index = i
		}
	}
	if index != -1 {
		RemoveAttribute(t, index)
	}

	t.Attr = append(t.Attr, html.Attribute{
		Key: key,
		Val: value,
	})
}

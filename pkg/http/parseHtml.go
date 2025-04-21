package http

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

func ParseHtml(content string) *html.Node {
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func Traverse(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("Element: %s\n", n.Data) // n.Data 是标签名，如 "div", "p" 等
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Traverse(c)
	}
}

func FindElements(n *html.Node, key string, value string) []*html.Node {
	var result []*html.Node
	if n.Type == html.ElementNode && n.Data == "div" {
		//fmt.Println(n.Attr)
		for _, a := range n.Attr {
			if a.Key == key && a.Val == value {
				return []*html.Node{n}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, FindElements(c, key, value)...)
	}
	return result
}

func GetText(n *html.Node, beforeTypeData string, beforeKey string, beforeValue string) string {
	var text string

	if n.Type == html.ElementNode && n.Data == beforeTypeData && len(n.Attr) > 0 {
		for _, a := range n.Attr {
			if a.Key == beforeKey && a.Val == beforeValue {
				text = GetOneText(n.NextSibling)
				return text
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t := GetText(c, beforeTypeData, beforeKey, beforeValue)
		if t != "" {
			text = t
		}
	}

	return text
}

func GetOneText(n *html.Node) string {
	var text string
	if n.Type == html.TextNode {
		text = n.Data
		return text
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			text = GetOneText(c)
		}
	}
	return text
}

func GetHrefWithPreFix(n *html.Node, prefix string) string {
	var href string
	if n.Data == "a" {
		if len(n.Attr) > 0 {
			if strings.HasPrefix(n.Attr[0].Val, prefix) {
				href = n.Attr[0].Val
				return href
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f := GetHrefWithPreFix(c, prefix)
		if f != "" {
			href = f
			break
		}
	}

	return href
}

func GetAttrWithAttrs(n *html.Node, typeData string, key string, value string, fetchKey string) string {
	var result string
	if n.Data == typeData {
		if len(n.Attr) > 0 {
			find := false
			v := ""
			for _, a := range n.Attr {
				if a.Key == key && a.Val == value {
					find = true
				}
				if a.Key == fetchKey {
					v = a.Val
				}
			}
			if find && v != "" {
				return v
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f := GetAttrWithAttrs(c, typeData, key, value, fetchKey)
		if f != "" {
			result = f
			break
		}
	}

	return result
}

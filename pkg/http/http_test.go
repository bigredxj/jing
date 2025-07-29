package http

import (
	"fmt"
	"testing"
)

func TestGetAttrWithAttrs(t *testing.T) {
	url := "https://cl.vxn75q.info/htm_data/2504/25/6766229.html"
	headers := map[string]string{
		"cookie": "ismob=1",
	}

	context := Get(url, headers)
	node := ParseHtml(context)
	result := GetAttrWithAttrs(node, "a", "id", "rmlink", "href")

	fmt.Println(result)
}

func TestDownLoad(t *testing.T) {
	url := "https://www.rmdown.com/download.php?action=magnet&ref=251ce1c66f4959499c333a58ed5d0c87ffa8edb1f9c&reff=9"
	headers := map[string]string{}

	context := Get(url, headers)
	fmt.Println(context)
}

func TestGet(t *testing.T) {
	url := "https://cl.vxn75q.info/htm_data/2504/25/6767204.html"
	headers := map[string]string{
		"cookie": "ismob=1",
	}

	context := Get(url, headers)
	fmt.Println(context)

}

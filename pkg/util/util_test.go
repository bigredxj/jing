package util

import (
	"fmt"
	"testing"
)

func TestParsePrefixNum(t *testing.T) {
	fmt.Println(GetPrefixNum("3300()") == "3300")
	fmt.Println(GetPrefixNum("3300(式基66)") == "3300")
	fmt.Println(GetPrefixNum("3300") == "3300")
	fmt.Println(GetPrefixNum("") == "")
	fmt.Println(GetPrefixNum("工3300") == "")

}

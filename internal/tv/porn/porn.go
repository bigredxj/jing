package porn

import (
	"fmt"
	"jing/internal/config"
	"jing/internal/register"
	"jing/pkg/http"
	"jing/pkg/util"
	"sort"
	"strconv"
	"strings"
	"time"
)

type PornSearch struct {
}

var p = PornSearch{}

func init() {
	register.RegisterTv(config.PornKind, p.DoSearch)
}

type sortUrl struct {
	url  string
	sort int
}

func (s PornSearch) DoSearch(tv config.Tv) []string {
	var result []string

	switch tv.Name {
	case "caoliu":
		result = CaoliuSearch(tv)
		break
	default:

	}

	return result
}

func CaoliuSearch(tv config.Tv) []string {
	result := make([]string, 0, 100)
	sortUrls := make([]sortUrl, 0, 100)
	for i := 1; i <= tv.MaxPage; i++ {
		rs := SearchOnePage(tv, i)
		sortUrls = append(sortUrls, rs...)
	}

	sort.Slice(sortUrls, func(i, j int) bool {
		return sortUrls[i].sort > sortUrls[j].sort
	})

	for _, s := range sortUrls {
		fmt.Println(s)
	}

	for i := 0; i < tv.DownloadSize && i < len(sortUrls); i++ {
		d := getDownLoad(sortUrls[i].url, tv.Heads)
		result = append(result, d)
	}

	return result
}

func SearchOnePage(tv config.Tv, page int) []sortUrl {
	sortUrls := make([]sortUrl, 0, 100)
	url := strings.Replace(tv.Url, "<page>", strconv.Itoa(page), 1)
	html := http.Get(url, tv.Heads)

	doc := http.ParseHtml(html)
	nodes := http.FindElements(doc, "class", "list t_one")
	for _, node := range nodes {
		h := http.GetHrefWithPreFix(node, tv.HomePrefix)
		if h != "" {
			text := http.GetText(node, "i", "class", "icon-dl")
			downNum, _ := strconv.Atoi(util.GetPrefixNum(text))
			if downNum > tv.DownloadNumLimit {
				h = tv.Domain + strings.Replace(h, "htm_mob", "htm_data", 1)
				sortUrl := sortUrl{
					h,
					downNum,
				}
				sortUrls = append(sortUrls, sortUrl)
			}

		}
	}
	return sortUrls
}

func getDownLoad(url string, headers map[string]string) string {
	result := ""
	context := http.Get(url, headers)
	node := http.ParseHtml(context)
	downPage := http.GetAttrWithAttrs(node, "a", "id", "rmlink", "href")
	hash := strings.Split(downPage, "hash=")[1]
	downUrl := "https://www.rmdown.com/download.php?action=magnet&ref=" +
		hash +
		"&reff=9"
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(downUrl)
	result = http.Get(downUrl, headers)
	if strings.Contains(result, "error") {
		time.Sleep(10 * time.Second)
		result = http.Get(downUrl, headers)
		if strings.Contains(result, "error") {
			time.Sleep(60 * time.Second)
			result = http.Get(downUrl, headers)
		}
	}
	fmt.Println(result)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	return result
}

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

func (s PornSearch) DoSearch(tv config.Tv) {
	switch tv.Name {
	case "caoliu":
		CaoliuSearch(tv)
		break
	default:

	}

}

func CaoliuSearch(tv config.Tv) {
	sortUrls := make([]sortUrl, 0, 100)
	for i := 1; i <= tv.MaxPage; i++ {
		fmt.Println(i)
		rs := SearchOnePage(tv, i)
		sortUrls = append(sortUrls, rs...)
	}

	fmt.Println("total urls>>>>>>>>>>>>>>>> " + strconv.Itoa(len(sortUrls)))
	sort.Slice(sortUrls, func(i, j int) bool {
		return sortUrls[i].sort > sortUrls[j].sort
	})

	downloadUrls := make([]string, 0, 100)
	hashs := make([]string, 0, 100)
	for i := 0; i < tv.DownloadSize && i < len(sortUrls); i++ {
		tmp := getDownLoadHash(sortUrls[i].url, tv.Heads)
		if tmp != "" {
			hashs = append(hashs, tmp[3:])
			downUrl := "https://www.rmdown.com/download.php?action=magnet&ref=" +
				tmp +
				"&reff=9"
			downloadUrls = append(downloadUrls, downUrl)
		}
	}
	path := util.GetWorkDir() + "/tmp/urls.txt"
	hashPath := util.GetWorkDir() + "/tmp/hash.txt"
	util.WriteOutput(hashPath, hashs)
	util.WriteOutput(path, downloadUrls)
	/*
		for _, url := range downloadUrls {
			getDownLoad(url, tv.Heads)
		}
	*/

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
			if downNum > tv.DownloadNumLimit*page {
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

func getDownLoadHash(url string, headers map[string]string) string {
	fmt.Println(url)
	context := http.Get(url, headers)
	hash := ""

	if context != "" {
		node := http.ParseHtml(context)
		downPage := http.GetAttrWithAttrs(node, "a", "id", "rmlink", "href")
		hash = strings.Split(downPage, "hash=")[1]
	}
	return hash

}

func getDownLoad(downUrl string, headers map[string]string) {
	result := ""

	//fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(downUrl)
	result = http.Get(downUrl, headers)
	if !strings.HasPrefix(result, "magnet") {
		time.Sleep(10 * time.Second)
		result = http.Get(downUrl, headers)
		for !strings.HasPrefix(result, "magnet") {
			fmt.Println(downUrl)
			fmt.Println("wait for operate mannually")
			time.Sleep(30 * time.Second)
			result = http.Get(downUrl, headers)

		}
	}
	path := util.GetWorkDir() + "/tmp/result.txt"
	util.AppendToFile(path, result)
	//fmt.Println(result)
	//fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
}

func GetDownLoadFromFile() {
	tv := config.ListTv(config.PornKind, "caoliu")
	path := util.GetWorkDir() + "/tmp/urls.txt"
	urls := util.ReadLinesFromFile(path)
	for _, u := range urls {
		getDownLoad(u, tv.Heads)
	}
	//util.PrintArrString(urls)

}

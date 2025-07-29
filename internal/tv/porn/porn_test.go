package porn

import (
	"jing/internal/config"
	"testing"
)

func TestCaoliuHttp(t *testing.T) {
	headers := map[string]string{
		"cookie": "ismob=1",
	}
	tv := config.Tv{
		Kind:             "porn",
		Name:             "caoliu",
		Url:              "https://cl.vxn75q.info/thread0806.php?fid=25&search=&page=1",
		Domain:           "https://cl.vxn75q.info",
		HomePrefix:       "/htm_mob",
		Heads:            headers,
		DownloadNumLimit: 1000,
	}
	CaoliuSearch(tv)

}

func TestGetDownLoadFromFile(t *testing.T) {
	GetDownLoadFromFile()
}

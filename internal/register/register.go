package register

import "jing/internal/config"

var tvMap = make(map[string]func(tv config.Tv) []string)

func RegisterTv(key string, f func(tv config.Tv) []string) {
	tvMap[key] = f
}

func GetTvFunc(key string) func(tv config.Tv) []string {
	f, _ := tvMap[key]
	return f
}

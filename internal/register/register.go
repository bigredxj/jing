package register

import "jing/internal/config"

var tvMap = make(map[string]func(tv config.Tv))

func RegisterTv(key string, f func(tv config.Tv)) {
	tvMap[key] = f
}

func GetTvFunc(key string) func(tv config.Tv) {
	f, _ := tvMap[key]
	return f
}

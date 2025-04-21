package tv

import (
	"jing/internal/config"
	"jing/internal/register"
)

type Search interface {
	DoSearch(tv config.Tv) []string
}

func DoSearch(kind string, tv config.Tv) []string {
	f := register.GetTvFunc(kind)
	return f(tv)
}

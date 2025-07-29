package tv

import (
	"jing/internal/config"
	"jing/internal/register"
)

type Search interface {
	DoSearch(tv config.Tv)
}

func DoSearch(kind string, tv config.Tv) {
	f := register.GetTvFunc(kind)
	f(tv)
}

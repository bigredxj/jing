package config

import (
	"fmt"
	"testing"
)

func TestListTvs(t *testing.T) {
	tv := ListTv(PornKind, "caoliu")

	fmt.Println(tv)

}

package examples

import (
	"fmt"
	"github.com/15732631053/go_md5"
	"testing"
)

func TestMd5(t *testing.T) {
	str := go_md5.StringMd5("aaac")
	fmt.Println(str)
}

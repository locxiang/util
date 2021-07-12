package util

import (
	"fmt"
	"gopkg.in/ffmt.v1"
	"testing"
)

func TestPutStringSlice(t *testing.T) {

	arr := []string{"1", "2", "3", "4"}

	str := "2"

	PutStringSlice(&arr, str)

	fmt.Printf("%s", arr)
}

func TestSliceStringRemoveEmpty(t *testing.T) {
	data := []string{
		"http://photo1.baihe.com/2019/07/21/1024_1024/467C59D1522F3F41EEED30A8BD02EC0B.jpg",
		"http://photo1.baihe.com/2019/07/21/1024_1024/7F2F0AF01D2BC4DE76FB3D53AB34D64A.jpg",
		"http://photo1.baihe.com/2019/07/21/1024_1024/552720DD20823F8EAE1D5AC3AF5BFADD.jpg",
		"",
		" ",
	}

	empty := SliceStringRemoveEmpty(data)

	ffmt.Pjson(empty)
}

package util

import (
	"fmt"
	"testing"
)

func TestTimeStr(t *testing.T) {

	str, err := TimeStr("2017-03-04 03:30:32")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%s", str)
}

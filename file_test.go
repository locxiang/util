package util

import (
	"bytes"
	"fmt"
	"gopkg.in/ffmt.v1"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestImageHash(t *testing.T) {

	open, err := os.Open("../dat/wong.jpg")
	if err != nil {
		t.Error(err)
	}
	defer open.Close()

	if s, err := ImageHash(open); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%s", s)
	}
}

func BenchmarkImageHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		open, err := os.Open("./1.jpg")
		if err != nil {
			b.Error(err)
		}
		if _, err := ImageHash(open); err != nil {
			b.Error(err)
		}
		open.Close()

	}
}

func TestDownloadFile(t *testing.T) {

	url := "http://wx1.sinaimg.cn/mw690/005S7cGCly1gohe2n7vrnj30ku0u2wgh.jpg"
	file, s, i, err := DownloadFile(url, url)
	if err != nil {
		panic(err)
	}
	ffmt.Pjson(file, s, i)

	time.Sleep(2 * time.Second)
}

func TestMemoryByte(t *testing.T) {

	url := "http://wx1.sinaimg.cn/mw690/005S7cGCly1gohe2n7vrnj30ku0u2wgh.jpg"
	body, _, _, err := DownloadFile(url, url)
	if err != nil {
		panic(err)
	}

	d1, _ := ioutil.ReadAll(bytes.NewBuffer(body))
	d2, _ := ioutil.ReadAll(bytes.NewBuffer(body))
	fmt.Printf("%d,%d", len(d1), len(d2))

}

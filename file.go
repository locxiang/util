package util

import (
	"bytes"
	"github.com/corona10/goimagehash"
	"github.com/pkg/errors"
	"image"
	_ "image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var defaultHttpClient *http.Client

func init() {

	defaultHttpClient = &http.Client{
		Timeout: 5 * time.Minute,
	}

}

// 下载文件
func DownloadFile(uri, referer string) (buff []byte, contentType string, size int64, err error) {

	if strings.HasPrefix(uri, `//`) {
		uri = "http:" + uri
	}
	if _, err := url.ParseRequestURI(uri); err != nil {
		return nil, "", 0, errors.Wrap(err, "网址错误")
	}
	// 1. 获取远程文件数据

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36`)
	req.Header.Add("Referer", referer)
	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		return nil, "", 0, errors.Wrap(err, "下载失败")
	}

	if resp.Body == nil {
		return nil, "", 0, errors.New("body is null")
	}
	defer resp.Body.Close()

	reader := MemoryByte(resp.Body)

	// 2. 读取服务器返回的文件大小
	size, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)

	contentType = resp.Header.Get("Content-Type")

	return reader, contentType, size, err
}

func ImageHash(r io.Reader) (string, error) {
	img1, _, err := image.Decode(r)
	if err != nil {
		return "", errors.Wrap(err, "解析图片结构失败")
	}
	hash1, err := goimagehash.DifferenceHash(img1)
	if err != nil {
		return "", errors.Wrap(err, "计算图片aHash失败")
	}

	if hash1.ToString() == "a:ffffff0000000000" {
		return "", errors.New("错误的hash计算")
	}

	return hash1.ToString(), nil
}

func MemoryByte(data io.ReadCloser) []byte {
	defer data.Close()
	buffer := bytes.NewBuffer(make([]byte, 0, 65536))
	io.Copy(buffer, data)
	temp := buffer.Bytes()
	length := len(temp)
	var body []byte
	//are we wasting more than 10% space?
	if cap(temp) > (length + length/10) {
		body = make([]byte, length)
		copy(body, temp)
	} else {
		body = temp
	}

	return body
}

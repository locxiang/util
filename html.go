package util

import (
	"fmt"
	"regexp"
	"strings"
)

// src提取
func RepHTML(typeA string, htmlData string) []string {
	reg := fmt.Sprintf("<%s[^>]+\\bsrc=[\"']([^\"']+)[\"']", typeA)
	var imgRE = regexp.MustCompile(reg)
	imgs := imgRE.FindAllStringSubmatch(htmlData, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}

// 过滤html多余标签
func RemoveHTML(src string) string {

	if src == "" {
		return ""
	}

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除iframe
	re, _ = regexp.Compile("\\<iframe[\\S\\s]+?\\</iframe\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "<br/>")

	src = strings.ReplaceAll(src, "\n", " <br/>")
	return strings.TrimSpace(src)

}

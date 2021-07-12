package util

import "strings"

func PutStringSlice(arr *[]string, str string) {
	if IsContain(*arr, str) == false {
		*arr = append(*arr, str)
	}
}

func SliceStringRemove(slice []string, elem string) []string {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return SliceStringRemove(slice, elem)
		}
	}
	return slice
}

func SliceStringRemoveRepeat(slc []string) []string {
	result := make([]string, 0)  //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

func SliceStringRemoveEmpty(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if strings.TrimSpace(v) == "" {
			slice = append(slice[:i], slice[i+1:]...)
			return SliceStringRemoveEmpty(slice)
		}
	}
	return slice
}

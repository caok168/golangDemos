package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	res, err := TagsConvertArrStr("{123,456}")

	res, err = TagsConvertArrStr("{}")
	fmt.Println("res:", res)
	fmt.Println("err:", err)
}

func TagsConvertArrStr(tagsInfo string) (tagsArrStr string, err error) {

	if tagsInfo == "" {
		return "", nil
	}

	tagsInfo = tagsInfo[1 : len(tagsInfo)-1]

	tagsArr := strings.Split(tagsInfo, ",")

	tagByte, err := json.Marshal(tagsArr)
	tagsArrStr = string(tagByte)

	return
}

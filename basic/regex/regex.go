package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse@gmail.com"

const texts = `
My email is ccmouse@gmail.com
email1 is abc@def.org
email2 is   kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	//re,err := regexp.Compile("ccmouse@gmail.com")
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)

	//match := re.FindString(text)
	match := re.FindAllString(texts, -1)
	fmt.Println(match)

	for _, mat := range match {
		fmt.Println(mat)
	}
}

package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com
emasd is  adff sdkl@org.com
asdfaev adfvae  sk@qq.com.cn.edu
qwerq@qqeq`

func main() {
	re, err := regexp.Compile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9]+[.[a-zA-Z0-9]+]*)`)
	if err != nil {
		panic(err)
	}
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}

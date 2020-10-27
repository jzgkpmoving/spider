package parser

import (
	"regexp"
	"spider/crawler/engine"
)

func ParseCity(contents []byte) engine.ParseResult {
	rex := regexp.MustCompile(`(http://album.zhenai.com/u/[0-9a-zA-Z]+)[^>]*>([^<]+)</a>`)
	all := rex.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range all {
		//fmt.Printf("city : %s , url : %s\n",m[2],m[1])
		result.Items = append(result.Items, "User :"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{string(m[1]), engine.NilParser})
	}
	//fmt.Println("match num",len(all))
	return result
}

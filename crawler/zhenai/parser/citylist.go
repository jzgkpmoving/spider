package parser

import (
	"regexp"
	"spider/crawler/engine"
)

func ParseCityList(contents []byte) engine.ParseResult {
	rex := regexp.MustCompile(`(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)[^>]*>([^<]+)</a>`)
	all := rex.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range all {
		//fmt.Printf("city : %s , url : %s\n",m[2],m[1])
		result.Items = append(result.Items, "City : "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{string(m[1]), ParseCity})
	}
	//fmt.Println("match num",len(all))
	return result
}

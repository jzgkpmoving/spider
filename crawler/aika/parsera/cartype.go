package parsera

import (
	"regexp"
	"spider/crawler/engine"
)

func ParseCarType(contents []byte) engine.ParseResult {
	rex := regexp.MustCompile(`href="(http://newcar.xcar.com.cn/car[^0]+0-[^"]+)"[^t]+title="([^"]+)"`)
	all := rex.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range all {
		//brandname := strings.Replace(string(m[2]), " ", "", -1)

		//result.Items = append(result.Items, "Type :"+brandname)
		result.Requests = append(result.Requests, engine.Request{string(m[1]), ParseBrand})
	}
	return result
}

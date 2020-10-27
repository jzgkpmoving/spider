package parsera

import (
	"regexp"
	"spider/crawler/engine"
	"strings"
)

func ParseCar(contents []byte) engine.ParseResult {
	rex := regexp.MustCompile(`href="(/[0-9]+/)[^>]+>([^<]+)</a>`)
	all := rex.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range all {
		brandname := strings.Replace(string(m[2]), " ", "", -1)
		//fmt.Printf("city : %s , url : %s\n",brandname,m[1])
		result.Items = append(result.Items, "Carname :"+brandname)
		result.Requests = append(result.Requests, engine.Request{"https://newcar.xcar.com.cn" + string(m[1]), ParseCarInfo})
	}
	//fmt.Println("match num",len(all))
	return result
}

package parsera

import (
	"regexp"
	"spider/crawler/engine"
)

func ParseBrand(contents []byte, url string) engine.ParseResult {
	//<a href="/car/0-0-0-0-251-0-0-0-0-0-0-0/" data-id="251">                            <span class="sign"> <img src="//img1.xcarimg.com/PicLib/logo/pl251_160s.png-40x30.jpg"></span>                            ARCFOX极狐                                                    </a>
	rex := regexp.MustCompile(`href="([^"]+)"[^>]+>[^>]+>[^>]+>[^>]+>([^<]+)</a>`)
	all := rex.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for i, m := range all {
		if i < 2 {
			continue
		}
		//brandname := strings.Replace(string(m[2]), " ", "", -1)
		//fmt.Printf("city : %s , url : %s\n",brandname,m[1])
		//result.Items = append(result.Items, "Brand :"+brandname)
		result.Requests = append(result.Requests, engine.Request{"https://newcar.xcar.com.cn" + string(m[1]), engine.NewFuncParser(ParseCar, "ParseCar")})
	}
	return result
}

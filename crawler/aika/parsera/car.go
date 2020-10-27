package parsera

import (
	"fmt"
	"regexp"
	"spider/crawler/engine"
	"strings"
)

type CarInfo struct {
	name             string
	price            string
	level            string
	structure        string
	oilconsumption   string
	displacement     string
	warranty         string
	transmissioncase string
}

func (car CarInfo) printInfo() {
	fmt.Printf("名字 ：%s\n价格：%s\n等级：%s   结构：%s\n油耗：%s   排量：%s\n质保：%s   变速箱：%s\n", car.name, car.price, car.level, car.structure, car.oilconsumption,
		car.displacement, car.warranty, car.transmissioncase)
}

const findname = `lt_f1">([^<]+)<[^>]+>[^>]+>([^<]+)<a`
const findprice = `ref_gd[^>]+>[^>]+>[^>]+>([^<]+)</a>`
const findlevel = `级　别：[^>]+>([^<]+)<`
const findoil = `油　耗：[^>]+>[^>]+>([^<]+)</a>`
const findwar = `质　保：[^>]+>([^<]+)<`
const findstruct = `结　构：[^>]+>([^<]+)<`
const findtran = `变速箱：[^>]+>[^>]+>[^>]+>([^<]+)</a>`
const finddisplace = `排　量：[^>]+>[^>]+>([^<]+)</a>`

func ParseCarInfo(contents []byte) engine.ParseResult {
	var rightcar CarInfo
	rexname := regexp.MustCompile(findname)
	all := rexname.FindAllSubmatch(contents, -1)
	rightcar.name = string(all[0][1]) + string(all[0][2])
	rexprice := regexp.MustCompile(findprice)
	all = rexprice.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.price = "未上市"
	} else {
		rightcar.price = string(all[0][1]) + "万"
	}
	rexstruct := regexp.MustCompile(findstruct)
	all = rexstruct.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.structure = "暂无"
	} else {
		temp := strings.Replace(string(all[0][1]), " ", "", -1)

		rightcar.structure = temp[2:]
	}

	rextran := regexp.MustCompile(findtran)
	all = rextran.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.transmissioncase = "暂无"
	} else {
		rightcar.transmissioncase = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	rexdis := regexp.MustCompile(finddisplace)
	all = rexdis.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.displacement = "暂无"
	} else {
		rightcar.displacement = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	rexlevel := regexp.MustCompile(findlevel)
	all = rexlevel.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.level = "暂无"
	} else {
		rightcar.level = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	rexoil := regexp.MustCompile(findoil)
	all = rexoil.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.oilconsumption = "暂无"
	} else {
		rightcar.oilconsumption = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	rexwar := regexp.MustCompile(findwar)
	all = rexwar.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.warranty = "暂无"
	} else {
		rightcar.warranty = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	result := engine.ParseResult{}
	result.Items = append(result.Items, rightcar)
	result.Requests = append(result.Requests, engine.Request{"", engine.NilParser})
	return result
}

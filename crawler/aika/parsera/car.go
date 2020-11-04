package parsera

import (
	"regexp"
	"spider/crawler/engine"
	"strings"
)

const findname = `lt_f1">([^<]+)<[^>]+>[^>]+>([^<]+)<a`
const findprice = `ref_gd[^>]+>[^>]+>[^>]+>([^<]+)</a>`
const findlevel = `级　别：[^>]+>([^<]+)<`
const findoil = `油　耗：[^>]+>[^>]+>([^<]+)</a>`
const findwar = `质　保：[^>]+>([^<]+)<`
const findstruct = `结　构：[^>]+>([^<]+)<`
const findtran = `变速箱：[^>]+>[^>]+>[^>]+>([^<]+)</a>`
const finddisplace = `排　量：[^>]+>[^>]+>([^<]+)</a>`

func ParseCarInfo(contents []byte, url string) engine.ParseResult {
	var withUrlCar engine.Item
	var rightcar engine.CarInfo
	rexname := regexp.MustCompile(findname)
	all := rexname.FindAllSubmatch(contents, -1)
	rightcar.Name = string(all[0][1]) + string(all[0][2])
	rexprice := regexp.MustCompile(findprice)
	all = rexprice.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.Price = "未上市"
	} else {
		rightcar.Price = string(all[0][1]) + "万"
	}
	rexstruct := regexp.MustCompile(findstruct)
	all = rexstruct.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.Struct = "暂无"
	} else {
		temp := strings.Replace(string(all[0][1]), " ", "", -1)

		rightcar.Struct = temp[2:]
	}

	rextran := regexp.MustCompile(findtran)
	all = rextran.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.Tran = "暂无"
	} else {
		rightcar.Tran = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	rexdis := regexp.MustCompile(finddisplace)
	all = rexdis.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.Disapp = "暂无"
	} else {
		rightcar.Disapp = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	rexlevel := regexp.MustCompile(findlevel)
	all = rexlevel.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.Level = "暂无"
	} else {
		rightcar.Level = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	rexoil := regexp.MustCompile(findoil)
	all = rexoil.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.Oilconsumption = "暂无"
	} else {
		rightcar.Oilconsumption = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	rexwar := regexp.MustCompile(findwar)
	all = rexwar.FindAllSubmatch(contents, -1)
	if all == nil {
		rightcar.Warryary = "暂无"
	} else {
		rightcar.Warryary = strings.Replace(string(all[0][1]), " ", "", -1)
	}

	result := engine.ParseResult{}
	withUrlCar.Payload = rightcar
	withUrlCar.Url = url
	withUrlCar.Type = "aika"
	idUrlRe := regexp.MustCompile("/([0-9]+)/")
	id := idUrlRe.FindAllSubmatch([]byte(url), -1)
	withUrlCar.Id = string(id[0][1])
	result.Items = append(result.Items, withUrlCar)
	result.Requests = append(result.Requests, engine.Request{"", engine.NilParser})
	return result
}

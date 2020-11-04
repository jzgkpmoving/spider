package engine

import (
	"fmt"
)

type Request struct {
	Url        string
	ParserFunc func([]byte, string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

type CarInfo struct {
	Name           string
	Price          string
	Level          string
	Struct         string
	Oilconsumption string
	Disapp         string
	Warryary       string
	Tran           string
}

func (car CarInfo) printInfo() {
	fmt.Printf("名字 ：%s\n价格：%s\n等级：%s   结构：%s\n油耗：%s   排量：%s\n质保：%s   变速箱：%s\n", car.Name, car.Price, car.Level, car.Struct, car.Oilconsumption,
		car.Disapp, car.Warryary, car.Tran)
}

func NilParser([]byte, string) ParseResult {
	return ParseResult{}
}

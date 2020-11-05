package engine

import (
	"fmt"
)

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url    string
	Parser Parser
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
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

type NilParser struct {
}

func (n NilParser) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		p, name,
	}
}

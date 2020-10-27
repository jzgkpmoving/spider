package parser

import (
	"spider/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	result := ParseCityList(contents)

	if len(result.Requests) != 494 {
		t.Errorf("result show have 494 requers buf had %d", len(result.Requests))
	}

}

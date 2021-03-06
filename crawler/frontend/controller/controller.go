package controller

import (
	"context"
	"github.com/olivere/elastic/v7"
	"net/http"
	"reflect"
	"spider/crawler/engine"
	"spider/crawler/frontend/model"
	"spider/crawler/frontend/view"
	"strconv"
	"strings"
)

type SearchResultHandle struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandle(template string) SearchResultHandle {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}
	return SearchResultHandle{view.CreateSearchResultView(template), client}
}
func (h SearchResultHandle) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).From(from).Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func (s SearchResultHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	q := strings.TrimSpace(request.FormValue("q"))
	from, err := strconv.Atoi(
		request.FormValue("from"))

	if err != nil {
		from = 0
	}

	var page model.SearchResult
	page, err = s.getSearchResult(q, from)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	err = s.view.Render(writer, page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

//func rewriteQueryString(q string) string {
//	re := regexp.MustCompile(`[A-Z][a-z]*:`)
//	re.ReplaceAllString(q , "Pa")
//}

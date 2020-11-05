package persiser

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"spider/crawler/engine"
	"testing"
)

func TestItemSaver(t *testing.T) {
	profile := engine.Item{
		"www.baidu/com",
		"1",
		"aika",
		engine.CarInfo{
			"东风雷诺-ARKANA",
			"未上市",
			"紧凑型SUV",
			"暂无",
			"暂无",
			"暂无",
			"暂无",
			"暂无",
		},
	}

	err := Save(profile)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type(profile.Type).Id(profile.Id).Do(context.Background())

	t.Logf("%+v", resp)

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	if actual != profile {
		t.Errorf("actual %+v !=profile %+v", actual, profile)
	}
	t.Logf("actual %+v !=profile %+v", actual, profile)
}

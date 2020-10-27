package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func Fetch(url string) ([]byte, error) {
	res, err := http.NewRequest(http.MethodGet, url, nil)
	res.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("rediretc", req)
			return nil
		},
	}

	resp, err := client.Do(res)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("get html body error")
		return nil, fmt.Errorf("wrong status code :%d", resp.StatusCode)
	}
	bodyreader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyreader)
	utf8Reader := transform.NewReader(bodyreader, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	//fmt.Printf("neirong:%s\n",all)
	return all, err
}

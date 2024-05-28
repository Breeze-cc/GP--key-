package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const (
	APIKEY = "MTY1MzQyMTY5YzViZmU0ODRhZjcyZDkwODUxNzAwYTNhNDA4MzQ0MDZmNGZjODc0Yjg0ZGZkOTY1NmM0ZTUxZHZhbGlkVW50aWw9MTcxNjk2NTQxOCZ1c2VyVG9rZW49ZWFlMzc3NmFiOGYwODYxNTIxNjVlZmFlNTdhYjIwZmY="
	ID     = "94HE6YATEI"
)

func getName(word string) (results []string) {
	url := "https://94he6yatei-dsn.algolia.net/1/indexes/steamdb/query?x-algolia-agent=SteamDB%20Autocompletion"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"hitsPerPage":10,"attributesToSnippet":null,"attributesToHighlight":"name","attributesToRetrieve":"objectID,lastUpdated","query":"%s"}`, word))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Referer", "https://steamdb.info/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	req.Header.Add("x-algolia-api-key", "MTY1MzQyMTY5YzViZmU0ODRhZjcyZDkwODUxNzAwYTNhNDA4MzQ0MDZmNGZjODc0Yjg0ZGZkOTY1NmM0ZTUxZHZhbGlkVW50aWw9MTcxNjk2NTQxOCZ1c2VyVG9rZW49ZWFlMzc3NmFiOGYwODYxNTIxNjVlZmFlNTdhYjIwZmY=")
	req.Header.Add("x-algolia-application-id", "94HE6YATEI")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var result SteamDB
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
	}
	for i, hit := range result.Hits {
		results = append(results, removeHTMLTags(hit.HighlightResult.Name.Value))
		fmt.Println(i, ":", removeHTMLTags(hit.HighlightResult.Name.Value))
	}
	return
}

// 去掉搜索结果中的标签字符串
func removeHTMLTags(input string) string {
	// 定义正则表达式来匹配HTML标签
	re := regexp.MustCompile(`(?i)<\/?[^>]+>`)
	// 替换匹配的HTML标签为空字符串
	result := re.ReplaceAllString(input, "")
	return result
}

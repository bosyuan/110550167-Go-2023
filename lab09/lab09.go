// lab09.go
package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

type Comment struct {
	Time     string
	Username string
	Content  string
}

func main() {
	// 解析命令行參數
	maxComments := flag.Int("max", 10, "Max number of comments to show")
	flag.Parse()

	// 檢查是否有未定義的 flag 或使用了錯誤的 flag
	if flag.NFlag() > 1 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(2)
	}

	// 創建一個新的 Collector
	c := colly.NewCollector()

	// // 設定 PTT 的 cookie，以繞過成年驗證
	// c.OnRequest(func(r *colly.Request) {
	// 	r.Headers.Set("Cookie", "over18=1")
	// })

	var comments []Comment

	c.OnHTML(".push", func(e *colly.HTMLElement) {
		// 提取留言時間
		time := e.ChildText(".push-ipdatetime")
		// 提取使用者名稱
		username := e.ChildText(".push-userid")
		// 提取留言內容
		content := e.ChildText(".push-content")

			comments = append(comments, Comment{
				Time:     time,
				Username: username,
				Content:  content,
			})
	})

	// 訪問 PTT 文章頁面
	err := c.Visit("https://www.ptt.cc/bbs/joke/M.1481217639.A.4DF.html")
	if err != nil {
		fmt.Println("Error visiting the page:", err)
		return
	}

	// 印出指定數量的留言
	for i, comment := range comments {
		if i >= *maxComments {
			break
		}
		fmt.Printf("%d. 名字：%s，留言%s，時間： %s\n", i+1, comment.Username, comment.Content, comment.Time)
	}
}

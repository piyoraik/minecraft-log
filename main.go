package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"minecraft-log/webhook"

	"github.com/joho/godotenv"
)

var times = []int{0, 0, 0, 0, 0}

func main() {
	err := godotenv.Load()
	if err != nil {
		// TODO: .env読めなかった場合の処理
	}

	// TODO: ５人固定なので汎用性のあるような形にしてゆく
	var names = []string{os.Getenv("USER1"), os.Getenv("USER2"), os.Getenv("USER3"), os.Getenv("USER4"), os.Getenv("USER5")}
	var viewNames = []string{os.Getenv("USER_NAME1"), os.Getenv("USER_NAME2"), os.Getenv("USER_NAME3"), os.Getenv("USER_NAME4"), os.Getenv("USER_NAME5")}
	var start, end string
	var msg []string

	var files []string
	files = dirWalk("./logs")

	for cnt, n := range names {
		for _, f := range files {
			start, end = fileLine(f, n, cnt, start, end)
		}
		start, end = "", ""
	}

	for i, t := range times {
		h := t / 3600
		m := (t / 60) % 60
		s := t % 60
		msg = append(msg, fmt.Sprintf("%s: %d時間%02d分%02d秒\n", viewNames[i], h, m, s))
	}

	dw := &webhook.DiscordWebhook{UserName: "PlayTime", Content: strings.Join(msg, "")}
	webhook.SendWebhook(os.Getenv("WEBHOOK"), dw)
}

// dirWalk 指定ディレクトリ内のファイル名を配列で返す
func dirWalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var paths []string
	for _, file := range files {
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}

// fileLine 指定ファイルの内容を
func fileLine(fileName string, name string, cnt int, start string, end string) (string, string) {
	date := filepath.Base(fileName)[:10]

	b, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(b)

	for i := 1; s.Scan(); i++ {
		line := s.Text()
		if strings.Contains(line, name+" joined the game") {
			start = date + " " + line[1:9]
		}
		if strings.Contains(line, name+" left the game") {
			end = date + " " + line[1:9]
		}

		if start != "" && end != "" {
			layout := "2006-01-02 15:04:05"
			st, _ := time.Parse(layout, start)
			et, _ := time.Parse(layout, end)

			times[cnt] = times[cnt] + int(et.Sub(st).Seconds())
			start, end = "", ""
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return start, end
}

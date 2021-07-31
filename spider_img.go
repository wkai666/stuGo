package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	chanImgUrls chan string
	wg2 sync.WaitGroup
	chanTask chan string
	reImg2 = `https?://[^'']+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
	// url3 = "https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/"
)

func main() {
	chanImgUrls = make(chan string, 1000000)
	chanTask = make(chan string, 26)

	for i := 1; i < 27; i++ {
		wg2.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}

	wg2.Add(1)
	go CheckOK()

	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go DownloadImg()
	}
	wg2.Wait()
}

func getImgUrls(url string)  {
	urls := getImgs(url)
	for _, urlStr := range urls {
		chanImgUrls <- urlStr
	}

	chanTask <- url
	wg2.Done()
}

func getImgs(url string) (urls []string) {
	pageStr := GetPageStrV3(url)
	re := regexp.MustCompile(reImg2)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到 %d 条结果\n", len(results))
	for _, result := range results {
		urlStr := result[0]
		urls = append(urls, urlStr)
	}
	return
}

func GetPageStrV3(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleErrorV3(err, "http.Get url")

	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleErrorV3(err, "ioutil.ReadAll")

	pageStr = string(pageBytes)
	return
}

func CheckOK()  {
	var count int
	for {
		urlStr := <- chanTask
		fmt.Printf("%s 完成了爬取任务\n", urlStr)
		count++
		if count == 26 {
			close(chanImgUrls)
			break
		}
	}
	wg2.Done()
}

func DownloadImg()  {
	for url1 := range chanImgUrls {
		filename := GetFilenameFromUrl(url1)
		ok := DownloadFile(url1, filename)
		if ok {
			fmt.Printf("%s download success\n", filename)
		} else {
			fmt.Printf("%s download failure\n", filename)
		}
	}
	wg2.Done()
}

func GetFilenameFromUrl(url string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex+1:]
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleErrorV3(err, "http.Get.url")

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	HandleErrorV3(err, "resp.Body")

	filename = "" + filename
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

func HandleErrorV3(err error, why string)  {
	if err != nil {
		fmt.Println(why, err)
	}
}
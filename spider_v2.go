package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reEmail = `\w+@\w+.\w`		// \w 代表打小写字母+数字+下划线
	reLinke = `href="(https://[\s\S]+?)"`	// s? 有或者没有 s
	rePhone = `1[3456789]\d\s?\d{4}\s?\d{4}`	// + 代表出现一次或者多次
	// \s\S 代表各种字符
	// +？ 代表贪婪模式
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12\d])|(3[01]))\d{3}[\dXx]`
	reImg = `https?://[^'']+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func HandleErrorV2(err error, why string)  {
	if err != nil {
		fmt.Println(why, err)
	}
}

func GetEmailV2(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleErrorV2(err, "http.Get url")

	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleErrorV2(err, "ioutail.ReadAll")

	pageStr = string(pageBytes)
	return pageStr
}

func GetIdCard(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reIdcard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetLink(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[1])
	}
}

func GetPhone(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetImg(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}

func main()  {
	// 爬取邮箱
	GetEmailV2(url)

}
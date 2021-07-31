package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var reQQEmail = `(\+d)@qq.com`
var url = "https://tieba.baidu.com/p/6051076813?red_tag=1573533731"

func GetEmail(url string)  {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")

	defer resp.Body.Close()
	fmt.Println(resp)

	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")

	pageStr := string(pageBytes)

	re := regexp.MustCompile(reQQEmail)

	results := re.FindAllStringSubmatch(pageStr, -1)

	for _, result := range results {
		fmt.Println("mail: ", result[0])
		fmt.Println("qq: ", result[1])
	}
}

func HandleError(err error, why string)  {
	if err != nil {
		fmt.Println(why, err)
	}
}

func main()  {
	GetEmail(url)
}

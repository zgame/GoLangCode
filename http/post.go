
package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func main() {


	url := "http://localhost:5000/user/login"

	postTest(url)
	HttpPostForm(url)
	HttpGet()


}

func postTest(url string)  {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("username=zsw&password=123"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

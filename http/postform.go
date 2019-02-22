package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"net/url"
)



func HttpPostForm(urlP string) {
	resp, err := http.PostForm(urlP,
		url.Values{"username": {"zsw"}, "password": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

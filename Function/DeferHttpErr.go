package main

import "net/http"

func do() error {

	res, err := http.Get("https://www.baidu.com")
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

func main() {
	println(do())
}

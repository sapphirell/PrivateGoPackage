package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ping_fantuanpu()

}
func ping_fantuanpu() {
	resp, err := http.Get("http://www.fantuanpu.com/ping")
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

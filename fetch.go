package main

import (
	"os"
	"net/http"

	"fmt"
	"io/ioutil"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if (err != nil) {
			fmt.Printf("fetch:%v\n", err)
			os.Exit(1)

		}
		fmt.Println(resp.Status)

		io.Copy(qwer,resp.Body)

	}
}

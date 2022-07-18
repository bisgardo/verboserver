package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	listenPort, err := strconv.Atoi(os.Getenv("LISTEN_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	proxyUrl := os.Getenv("PROXY_URL")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		client := &http.Client{}

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}

		proxyReq, err := http.NewRequest(req.Method, proxyUrl+req.URL.Path, bytes.NewReader(reqBody))
		if err != nil {
			log.Println(err)
		}
		proxyReq.Header = req.Header

		proxyRes, err := client.Do(proxyReq)
		if err != nil {
			log.Println(err)
		}
		defer proxyRes.Body.Close()

		resBody, err := ioutil.ReadAll(proxyRes.Body)
		if err != nil {
			log.Println(err)
		}
		if _, err := w.Write(resBody); err != nil {
			log.Println(err)
		}

		fmt.Printf("%s: %s\t=>\t%s\n", req.URL.Path, strings.TrimSpace(string(reqBody)), strings.TrimSpace(string(resBody)))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
}

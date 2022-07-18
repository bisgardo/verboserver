package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	listenPort := 8028
	proxyUrl := "http://localhost:8080"

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		client := &http.Client{}

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}

		proxyReq, err := http.NewRequest(req.Method, proxyUrl+req.URL.Path, bytes.NewReader(reqBody))
		if err != nil {
			log.Fatalln(err)
		}
		proxyReq.Header = req.Header

		proxyRes, err := client.Do(proxyReq)
		if err != nil {
			log.Fatalln(err)
		}
		defer proxyRes.Body.Close()

		resBody, err := ioutil.ReadAll(proxyRes.Body)
		if err != nil {
			log.Fatalln(err)
		}
		if _, err := w.Write(resBody); err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s: %s\t=>\t%s\n", req.URL.Path, reqBody, resBody)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
}

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	res *http.Response
	err error
}

func process() {
	ctx, cancle := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancle()
	tr := &http.Transport{}
	client := http.Client{Transport: tr}
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		fmt.Println("new request failed:", err)
	}
	c := make(chan Result, 1)
	go func() {
		res, err := client.Do(req)
		pack := Result{res: res, err: err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		r := <-c
		fmt.Print("Timeout! err:", r.err)
	case r := <-c:
		defer r.res.Body.Close()
		out, _ := ioutil.ReadAll(r.res.Body)
		fmt.Printf("Server Response:%s", out)
	}
}

func main() {
	process()
}

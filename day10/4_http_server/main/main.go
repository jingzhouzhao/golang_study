package main

import (
	"io"
	"net/http"
	"fmt"
)


func SimpleServer(w http.ResponseWriter,r *http.Request){
	io.WriteString(w, "hello world")
}

var s string = `
	<form method="post">
		<input name="input">
		<input name="input1">
		<input type="submit"/>
	</form>
`

func logPanics(h http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request){
		defer func(){
			if err:=recover();err!=nil{
				fmt.Println("[%v] caught panic: %v",r.RemoteAddr,err)
			}
		}()
		h(w,r)
	}
}

func FormServer(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case http.MethodGet:
		io.WriteString(w, s)
	case http.MethodPost:
		r.ParseForm()
		io.WriteString(w, r.FormValue("input"))
		io.WriteString(w, "<br>")
		io.WriteString(w, r.Form["input2"][0])
	}
}

func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	http.ListenAndServe("0.0.0.0:8888", nil)
}
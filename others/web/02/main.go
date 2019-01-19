package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello golang")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // 仅仅能解析get方法

	fmt.Println(r.Form)

//	打印请求方法
	fmt.Println("请求方法：", r.Method)

	// postman 必须选择 x-www-form-urlencoded !!!!!
	if r.Method == "POST" {
		username := r.Form.Get("username")
		age := r.Form.Get("age")
		fmt.Println("用户名:", username)
		fmt.Println("年龄：", age)
	} else {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("服务器错误:", err.Error())
	}
}

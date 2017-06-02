package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	//"net/url"
	"log"
	"os"
	"io"
	"crypto/md5"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, req * http.Request){
	fmt.Println("method:", req.Method) //获取请求的方法
	if req.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		//fmt.Println(strings.SplitN(handler.Filename, "/", 1)[len(strings.SplitN(handler.Filename, "/", 1)) - 1])
		f, err := os.OpenFile("../test/" + strings.Split(handler.Filename, "/")[len(strings.Split(handler.Filename, "/")) - 1], os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	//解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

//func index(rw http.ResponseWriter, req * http.Request){
//	t, _ = template.ParseFiles("index.tpl")
//	t.Execute(rw, nil)
//}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./index.tpl")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		if checkPasswd(r.Form["password"][0]) {
			t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
			t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")

			//template.HTMLEscape(w, []byte("<script>alert()</script>"))
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
		} else {
			fmt.Printf("密码长度不符合")
			t, _ := template.ParseFiles("/home/zhuxinquan/IdeaProjects/GoLanguage/GoFormManage/main/index.tpl")
			t.Execute(w, "密码长度不符合")
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func checkPasswd(s string) bool{
	if len(s) > 16 || len(s) < 6 {
		return false
	}
	return true
}

func main() {
	//v := url.Values{}
	//v.Set("name", "name")
	//v.Add("friend", "friend1")
	//v.Add("friend", "friend2")
	//fmt.Println(v.Encode())
	//fmt.Println(v["friend"][1])
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

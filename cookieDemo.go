package main
 
import (
    "io"
    "net/http"
    "log"
    "fmt"
    "time"
)
 
func DrawMenu(w http.ResponseWriter){
    w.Header().Set("Content-Type", "text/html")
//	io.WriteString(w, "<script src='http://h5.bwgame.com.cn/Scripts/jquery-2.1.4.js' type='text/javascript'></script>" + "\n")
//	io.WriteString(w,"<script type='text/javascript'>jQuery(document).ready(function(){$.ajax({beforeSend: function (xhr) {xhr.setRequestHeader('Mycookie', '110');},type: 'POST',url: 'http://test.com:3000/writecookie',success: function(output, status, xhr) {console.log('##jquery.cookie:##      \n' +jQuery.cookie('mycookie'));console.log('##allHeaders:##       \n' +xhr.getAllResponseHeaders());},cache: false});});</script>"+"\n")
    io.WriteString(w, "<a href='/'>HOME <ba><br/>" + "\n")
    io.WriteString(w, "<a href='/readcookie'>Read Cookie <ba><br/>" + "\n")
    io.WriteString(w, "<a href='/writecookie'>Write Cookie <ba><br/>" + "\n")
    io.WriteString(w, "<a href='/deletecookie'>Delete Cookie <ba><br/>" + "\n")
 
}
 
func IndexServer(w http.ResponseWriter, req *http.Request) {
    // draw menu
    DrawMenu(w)
}
 
 
func ReadCookieServer(w http.ResponseWriter, req *http.Request) {
 
    // draw menu
    DrawMenu(w)
 
    // read cookie
    var cookie,err = req.Cookie("testcookiename")
    if err == nil {
        var cookievalue = cookie.Value
        io.WriteString(w, "<b>get cookie value is " + cookievalue + "</b>\n")
    }
 
}
 
func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
	if(req.Method == "OPTIONS"){
		w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Tocookie")
		w.Header().Add("Access-Control-Max-Age", "30")
	}else if (req.Method == "POST"){
		if(req.Header.Get("Origin") == "http://test.com:9001"){
			w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			expire := time.Now().AddDate(0, 0, 1)
    		cookie := http.Cookie{Name: "testcookiename", Value: "test  cooki eva lue", Path: "/", Expires: expire, MaxAge: 86400}
			http.SetCookie(w, &cookie)		
		}
	}else{
		fmt.Println("请求失败")
	}
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Tocookie,Mycookie")



	//fmt.Println(req.Header.Get("Mycookie"))
//	fmt.Println(req.Cookies())
    // set cookies.
//    expire := time.Now().AddDate(0, 0, 1)
//    cookie := http.Cookie{Name: "testcookiename", Value: "test  cooki eva lue", Path: "/", Expires: expire, MaxAge: 86400}
//	w.Header().Set("Access-Control-Allow-Origin", "http://test.com:9001")//允许访问所有域
//	w.Header().Set("Access-Control-Allow-Origin", "*")//允许访问所有域
//	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Tocookie,Mycookie")//header的类型
//	w.Header().Set("tocookie", "this is a coo kfds")
//	w.Header().Set("content-type", "application/json")//返回数据格式是json
//    http.SetCookie(w, &cookie)
 
    //
    // we can not set cookie after writing something to ResponseWriter
    // if so ,we cannot set cookie succefully.
    //
    // so we have draw menu after set cookie
    DrawMenu(w)
 
}
 
 
func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
 
    // set cookies.
    cookie := http.Cookie{Name: "testcookiename", Path: "/", MaxAge: -1}
    http.SetCookie(w, &cookie)
 
    // ABOUT MaxAge
    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
 
    // draw menu
    DrawMenu(w)
 
}
 
func main() {
 
    http.HandleFunc("/", IndexServer)
    http.HandleFunc("/readcookie", ReadCookieServer)
    http.HandleFunc("/writecookie", WriteCookieServer)
    http.HandleFunc("/deletecookie", DeleteCookieServer)
 
    fmt.Println("listen on 3000")
    err := http.ListenAndServe(":3000", nil)
 
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
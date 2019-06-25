package main

import (
	"fmt"
    "io/ioutil"
    // "strings"
    "github.com/xxtea/xxtea-go/xxtea"
)

func decrypt_img(){
	key := "kaile_game12345a"
    fileName :="encrypt_img.png"
	if contents,err := ioutil.ReadFile(fileName);err == nil {
        //因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
        // result := strings.Replace(string(contents),"\n","",1)
        // fmt.Println(result)
        decrypt_data := xxtea.Decrypt(contents, []byte(key))
        if ioutil.WriteFile("decrypt_img.png",decrypt_data,0644) == nil {
            fmt.Println("解密文件成功")
        }
	}
}

func encrypt_img(){
    key := "kaile_game12345a"
    fileName :="img.png"
    if contents,err := ioutil.ReadFile(fileName);err == nil {
        //因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
        // result := strings.Replace(string(contents),"\n","",1)
        // fmt.Println(result)
        encrypt_data := xxtea.Encrypt(contents, []byte(key))
        if ioutil.WriteFile("encrypt_img.png",encrypt_data,0644) == nil {
            fmt.Println("加密文件成功")
        }
	}
}

func test_encrypt_str(){
    str := "Hello World! 你好，中国！"
    key := "kaile_game12345a"
    encrypt_data := xxtea.Encrypt([]byte(str), []byte(key))
    decrypt_data := string(xxtea.Decrypt(encrypt_data, []byte(key)))
    if str == decrypt_data {
        fmt.Println("success!")
    } else {
        fmt.Println("fail!")
    }
}

func main() {
	// encrypt_img()
    decrypt_img()
}
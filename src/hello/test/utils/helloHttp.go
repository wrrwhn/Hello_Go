package utils

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

var Port *int

func HelloGet() {

	// init
	url := "https://www.baidu.com"

	// request
	res, err := http.Get(url)
	if nil != err {
		log.Printf("fail to read [%v]\n", url)
	}

	log.Printf("Get [%v]\n", res)
}

func HelloPost() {

	// init
	width, height, x, y := 1000, 960, 140, 0
	url := fmt.Sprintf("http://dev-live.yunkai.com/basic-api/resource/cut?width=%d&height=%d&x=%d&y=%d", width, height, x, y)
	filename := "D:\\资料备份\\OneDrive\\Pictures\\生活\\MIAO\\467834901193557943.jpg"

	// buf
	// buf.read
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(filename)
	if nil != err {
		log.Printf("fail to open image[%v]\n", filename)
		return
	}
	defer f.Close()

	// post
	res, err := http.Post(url, w.FormDataContentType(), b)
	if nil != err {
		return
	}
	defer res.Close()
	b, err := ioutil.ReadAll(res.Body)
	if nil != err {
		log.Println("fail to read response")
		return
	}
	log.Printf("Post: [%v]\n", string(b))

}

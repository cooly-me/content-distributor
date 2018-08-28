package main
import (
	"fmt"
	"net/http"
	"log"
	"time"
)

// get 转发请求
func RequestAnother(aid string ,site_id string)  {
	fmt.Println("124677")
	url := ""
	//url := "http://www.ivipcard.cn/parenting/activity/adminActivity/sendAlertNewMsg?aid ="+ aid +"&site_id ="+site_id
	body, err := http.Get(url)
	time.Sleep(time.Millisecond * 200)
	 if err != nil{
		fmt.Println("error:",err)
	 }
	fmt.Println(1,body)
	}

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	site_slice := r.Form["site_id"]
	aid_slice := r.Form["aid"]

	// 转换成字符串
	var aid string
	var site_id string
	for _,v := range site_slice{
		site_id = site_id + v
	}
	for _,v := range aid_slice{
		aid = aid + v
	}

	// 开启多协程
	for i:= 1; i <= 2 ; i++{
		go RequestAnother(aid,site_id)
	}
}

func main() {
	http.HandleFunc("/", hello) //设置访问的路由
	err := http.ListenAndServe(":9092", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

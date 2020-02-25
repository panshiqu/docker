package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 读配置文件
// 监听端口
// 写标准输出
// 写日志文件

var conf = flag.String("conf", "./conf.json", "conf")

type tag struct {
	Name string
}

func handleSignal(server *http.Server, w io.Writer) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	s := <-c
	fmt.Fprintf(w, "Got signal: %v\n", s)
	fmt.Printf("Got signal: %v\n", s)

	server.Close()
}

func main() {
	flag.Parse()

	f, err := os.Open(*conf)
	if err != nil {
		fmt.Println("Open", err)
		return
	}

	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("ReadAll", err)
		return
	}

	js := &tag{}

	if err := json.Unmarshal(body, js); err != nil {
		fmt.Println("Unmarshal", err)
		return
	}

	hn, err := os.Hostname()
	if err != nil {
		fmt.Println("Hostname", err)
		return
	}

	out, err := os.OpenFile(fmt.Sprintf("./log/out_%s.log", hn), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Create", err)
		return
	}
	defer out.Close()

	slot := os.Getenv("TASKSLOT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(out, "[%s][%s] %s, %q\n", hn, slot, js.Name, html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "[%s][%s] %s, %q\n", hn, slot, js.Name, html.EscapeString(r.URL.Path))
		fmt.Printf("[%s][%s] %s, %q\n", hn, slot, js.Name, html.EscapeString(r.URL.Path))
	})

	fmt.Fprintln(out, "http server start.")
	fmt.Println("http server start.")

	s := &http.Server{
		Addr:         ":9090",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go handleSignal(s, out)

	err = s.ListenAndServe()
	fmt.Fprintln(out, err)
	fmt.Println(err)
}

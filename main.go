package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 解析命令行参数
	var (
		url            string
		cookieName     string
		cookieValue    string
		requestInterval int
	)
	flag.StringVar(&url, "url", "http://example.com", "Target URL to send requests to")
	flag.StringVar(&cookieName, "cookie-name", "your_cookie_name", "Name of the cookie")
	flag.StringVar(&cookieValue, "cookie-value", "your_cookie_value", "Value of the cookie")
	flag.IntVar(&requestInterval, "interval", 5, "Interval between requests in seconds")
	flag.Parse()

	// 设置请求间隔
	intervalDuration := time.Duration(requestInterval) * time.Second

	// 准备优雅终止
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	done := make(chan bool, 1)

	// 启动请求循环
	go func() {
		client := &http.Client{}
		for {
			select {
			case <-done:
				return
			default:
				// 创建带有Cookie的请求
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					log.Println("Error creating request:", err)
					continue
				}
				cookie := &http.Cookie{
					Name:   cookieName,
					Value:  cookieValue,
					Path:   "/",
					Domain: url,
				}
				req.AddCookie(cookie)

				// 发送请求
				resp, err := client.Do(req)
				if err != nil {
					log.Println("Error sending request:", err)
					continue
				}
				log.Printf("Received response with status code: %d\n", resp.StatusCode)
				resp.Body.Close()

				// 等待指定的时间间隔
				time.Sleep(intervalDuration)
			}
		}
	}()

	// 等待信号量以优雅地终止程序
	<-signals
	done <- true
	log.Println("Program terminated gracefully")
}

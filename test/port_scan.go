package test

import (
	"fmt"
	"net"
	"time"
)

func main() {
	t := time.Now()
	maxPosrt := 1024
	ports := make(chan int, maxPosrt)
	result := make(chan int)
	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}
	go func() {
		for i := 1; i < maxPosrt; i++ {
			ports <- i
		}
	}()
	for i := 0; i < maxPosrt; i++ {
		port := <-result
		fmt.Println(port)
	}
	since := time.Since(t)/1 ^ 9
	close(ports)
	close(result)
	fmt.Printf("使用时间:%v", since)
}

func worker(ports chan int, result chan int) {
	for value := range ports {
		address := fmt.Sprintf("192.168.110.233:%d", value)
		c, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			continue
		}
		c.Close()
		result <- value
		fmt.Printf("端口打开:%s\n", address)
	}
}

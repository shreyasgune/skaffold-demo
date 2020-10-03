package main

import (
	"fmt"
	"net"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	dt := time.Now()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello from %s, it's %s right now.", getIP(), dt.Format("01-02-2006 15:04:05 Monday")),
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("test endpoint is working"),
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("ping returns 200 response"),
		})
	})
	r.Run(fmt.Sprintf("%s:8085",getIP()))
}

func getIP() net.IP {
	var ip net.IP
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}
	return ip
}

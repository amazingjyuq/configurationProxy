package main

import (
	"fmt"
	"github.com/amazingjyuq/configurationProxy/pkg"
)

func main() {
	pkg.SetProxySettings("127.0.0.1:8001")
	pkg.EnableProxySettings()
	fmt.Println("yes")
}
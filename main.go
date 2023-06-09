package main

import (
	"fmt"
	"github.com/amazingjyuq/configurationProxy/pkg"
)

func main() {
	pkg.WindowsSetProxySettings("127.0.0.1:8001")
	pkg.WindowsEnableProxySettings()
	fmt.Println("yes")
}
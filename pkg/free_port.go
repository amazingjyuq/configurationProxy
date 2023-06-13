package pkg

import (
	lnet "net"
)

func GetFreePort(retry int) (int, error) {
    addr, err := lnet.ResolveTCPAddr("tcp", "127.0.0.1:0")
    if err != nil {
        if retry >= 3 {
            return 0, err
        } else {
            return GetFreePort(retry + 1)
        }
    }
    l, err := lnet.ListenTCP("tcp", addr)
    if err != nil {
        if retry >= 3 {
            return 0, err
        } else {
            return GetFreePort(retry + 1)
        }
    }
    defer l.Close()
    return l.Addr().(*lnet.TCPAddr).Port, nil
}
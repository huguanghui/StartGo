package main

import (
	"fmt"

	"github.com/huguanghui/hghstring"
)

func main() {
	conn, err := hghstring.Hghdial("tcp", "192.168.3.35:6000")
	if err == nil {
		for {
			buf := make([]byte, 1024)
			if length, err := conn.Read(buf); err == nil {
				if length > 0 {
					buf[length] = 0
					fmt.Printf("%s", string(buf[0:length]))
				}
			}
		}
	}
}

package main

import (
	"log"
	"net"
	"os"
)

const (
	gatewayDefault   = "127.0.0.1:4001"
	UnmatchedReplyID = int64(-9223372036854775808)
)

func Main(args []string) error {

	conn, err := net.Dial("tcp", gatewayDefault)
	if err != nil {
		return err
	}

	if _, err := conn.Write([]byte{0x63}); err != nil {
		return err
	}
	return nil
}
func main() {
	if err := Main(os.Args); err != nil {
		log.Fatal(err)
	}
}

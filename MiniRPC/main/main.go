package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"miniRPC"
	"miniRPC/codec"
)

func main() {
	go startServer()

	// client
	conn, _ := net.Dial("tcp", ":9999")
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)

	// option := minirpc.Option{MagicNumber: 0x33333,CodecType: "application/json"}
	// _ = json.NewEncoder(conn).Encode(option)
	_ = json.NewEncoder(conn).Encode(minirpc.DefaultOption)

	cc := codec.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Toma.Ber",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("req %d", h.Seq))

		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)

		log.Println("reply:", reply)
	}

}

func startServer() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", lis.Addr())

	server := minirpc.NewServer()
	server.Accept(lis)
}

package minirpc

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"reflect"
	"sync"

	"miniRPC/codec"
)

const MagicNumber = 0x33333

type Option struct {
	MagicNumber int
}

var DefaultOption = &Option{MagicNumber: 0x33333}

type Server struct{}

type request struct {
	h            *codec.Header
	argv, replyv reflect.Value
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		go server.ServeConn(conn)
	}
}

func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()

	var opt Option
	// use json to decode Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: options error: ", err)
		return
	}
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}

	cc := codec.NewGobCodec(conn)
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	// handle header and body
	for {
		// read header and body(request)
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break //impossible to recover,close the connection
			}
			req.h.Error = err.Error()
			// send response with error
			server.sendResponse(cc, req.h, struct{}{}, sending)
			continue
		}
		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg)

	}
	_ = cc.Close()
}

func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}

	req := &request{h: &h}
	// TODO: body no specific
	req.argv = reflect.New(reflect.TypeOf(""))
	if err := cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}
	return req, nil
}

func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	sending.Lock()
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response error:", err)
	}
}

func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	// process
	body := "not implement yet"
	server.sendResponse(cc, req.h, body, sending)
}

package minirpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"miniRPC/codec"
	"net"
)

// TODO: should use mutex?
type Client struct {
	cc      codec.Codec
	seq     uint64 //distingush different client call request
	pending map[uint64]*Call
}

var _ io.Closer = (*Client)(nil)

func (client *Client) Close() error {
	return client.cc.Close()
}

// func (t *T) MethodName(argType T1, replyType *T2) error
type Call struct {
	Seq           uint64
	ServiceMethod string
	Args          interface{}
	Reply         interface{}
	Error         error
	Done          chan *Call //Asynchronous communication
}

func (call *Call) done() {
	call.Done <- call
}

func (client *Client) registerCall(call *Call) (uint64, error) {
	call.Seq = client.seq
	client.pending[client.seq] = call
	client.seq++
	return call.Seq, nil
}

func (client *Client) removeCall(seq uint64) *Call {
	call:=client.pending[seq]
	delete(client.pending, seq)
	return call
}

func (client *Client) terminateCall(err error) {
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}

// Receive response operation
func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}
		call := client.removeCall(h.Seq)
		switch {
		case call == nil:
			err = client.cc.ReadBody(nil)
		case h.Error != "":
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReadBody(nil)
			call.done()
		default:
			err = client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
	client.terminateCall(err)
}


func NewClient(conn net.Conn,opt *Option)(*Client,error){
	f:=codec.CodecFuncMap[opt.CodecType]
	if f==nil{
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}
	
	// send options to server
	if err:=json.NewEncoder(conn).Encode(opt);err!=nil{
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}
	client:=&Client{
		cc: f(conn),
		seq: 1,
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client,nil
}

func Dial(network,address string,opt *Option)(client *Client,err error){
	conn,err:=net.Dial(network,address)
	if err!=nil {
		return nil,err
	}
	defer func(){
		if client==nil{
			_=conn.Close()
		}
	}()
	return NewClient(conn,opt)
}

func (client *Client)send(call *Call){
	seq,err:=client.registerCall(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}
	header:=codec.Header{
		Seq: seq,
		Error: "",
		ServiceMethod: call.ServiceMethod,
	}

	if err:=client.cc.Write(&header,call.Args);err!=nil{
		// 
	}
}
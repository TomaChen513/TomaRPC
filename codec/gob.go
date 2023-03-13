package codec

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"reflect"
	"sync"
	"unicode"
	"unicode/utf8"
)

// use a new gob determine that the first letter must be capitalized when rpc decoding.
// Use reflection to determine field names.

var checked map[reflect.Type]bool
var mu sync.Mutex

type tomaGobEncoder struct {
	gob *gob.Encoder
}

func NewEncoder(w io.Writer) *tomaGobEncoder  {
	return &tomaGobEncoder{
		gob: gob.NewEncoder(w),
	}
}

func (enc *tomaGobEncoder)Encode(e interface{})error{
	checkValue(e)
	return enc.gob.Encode(e)
}

func (enc *tomaGobEncoder)EncodeValue(value reflect.Value)error{
	checkValue(value.Interface())
	return enc.gob.EncodeValue(value)
}

type tomaGobDecoder struct {
	gob *gob.Decoder
}

func NewDecoder(conn io.ReadWriteCloser)*tomaGobDecoder{
	return &tomaGobDecoder{
		gob: gob.NewDecoder(conn),
	}
}

func (dec *tomaGobDecoder)Decode(e interface{})error{
	checkValue(e)
	return dec.gob.Decode(e)
}

func checkValue(value interface{}){
	checkType(reflect.TypeOf(value))
}

func checkType(t reflect.Type){
	k:=t.Kind()

	mu.Lock()
	if checked==nil {
		checked=map[reflect.Type]bool{}
	}
	if checked[t] {
		mu.Unlock()
		return
	}

	checked[t]=true
	mu.Unlock()

	switch k {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f:=t.Field(i)
			r,_:=utf8.DecodeRuneInString(f.Name)
			if !unicode.IsUpper(r){
				fmt.Printf("labgob error: lower-case field %v of %v in RPC or persist/snapshot will break your Raft\n",
					f.Name, t.Name())
			}
			// recursion check
			checkValue(f.Type)
		}
		return

	case reflect.Slice,reflect.Array,reflect.Ptr:
		checkType(t.Elem())
		return
	
	case reflect.Map:
		checkType(t.Elem())
		checkType(t.Key())
		return

	default:
		return
	}
}


type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *tomaGobDecoder
	enc  *tomaGobEncoder
}

var _ Codec = (*GobCodec)(nil)

func NewTomaGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  NewDecoder(conn),
		enc:  NewEncoder(buf),
	}
}



func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err = c.enc.Encode(h); err != nil {
		log.Println("rpc: gob error encoding header:", err)
		return
	}
	if err = c.enc.Encode(body); err != nil {
		log.Println("rpc: gob error encoding body:", err)
		return
	}
	return
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}

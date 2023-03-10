package mitgob

// This library is used to determine that the first letter must be capitalized when rpc decoding
// Use reflection to determine field names

import (
	"encoding/gob"
	"fmt"
	"io"
	"reflect"
	"sync"
	"unicode"
	"unicode/utf8"
)

var checked map[reflect.Type]bool
var mu sync.Mutex

type mitEncoder struct {
	gob *gob.Encoder
}

func NewEncoder(w io.Writer) *mitEncoder  {
	return &mitEncoder{
		gob: gob.NewEncoder(w),
	}
}

func (enc *mitEncoder)Encode(e interface{})error{
	checkValue(e)
	return enc.gob.Encode(e)
}

func (enc *mitEncoder)EncodeValue(value reflect.Value)error{
	checkValue(value.Interface())
	return enc.gob.EncodeValue(value)
}

type mitDecoder struct {
	gob *gob.Decoder
}

func NewDecoder(r io.Reader)*mitDecoder{
	return &mitDecoder{
		gob: gob.NewDecoder(r),
	}
}

func (dec *mitDecoder)Decode(e interface{})error{
	checkValue(e)
	return dec.gob.Decode(e)
}

func Register(value interface{}){
	checkValue(value)
	gob.Register(value)
}

func RegisterName(name string,value interface{}){
	checkValue(value)
	gob.RegisterName(name,value)
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


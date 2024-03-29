// The codec defines the codec scheme in the RPC framework, and defines the internal form of the basic request header and the form of the body, only Gob is used, and it can be extended later if needed

package codec

import "io"

type Header struct {
	ServiceMethod string //Declare the method name
	Seq           uint64 // unique num chosen by client
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type Type string

// codec type
const (
	GobType Type = "application/gob"
)

type NewCodecFunc func(io.ReadWriteCloser) Codec

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}

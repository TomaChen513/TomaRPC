package codec

import "io"

type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header)error
	ReadBody(interface{})error
	Write(*Header,interface{})error
}


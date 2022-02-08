package codec

import (
	"github.com/ugorji/go/codec"
)

type JsonCodec struct{}

func (*JsonCodec) Encode(v interface{}) ([]byte, error) {
	return encode(new(codec.JsonHandle), v)
}

func (*JsonCodec) Decode(in []byte, v interface{}) error {
	return decode(new(codec.JsonHandle), in, v)
}

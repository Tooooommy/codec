package codec

import "github.com/ugorji/go/codec"

type Codec interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}

func encode(h codec.Handle, v interface{}) (out []byte, err error) {
	err = codec.NewEncoderBytes(&out, h).Encode(v)
	return
}

func decode(h codec.Handle, in []byte, v interface{}) (err error) {
	err = codec.NewDecoderBytes(in, h).Decode(v)
	return
}

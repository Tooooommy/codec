package codec

import "github.com/ugorji/go/codec"

type Msgpack struct {
}

func (*Msgpack) Encode(v interface{}) ([]byte, error) {
	return encode(new(codec.MsgpackHandle), v)
}

func (*Msgpack) Decode(in []byte, v interface{}) error {
	return decode(new(codec.MsgpackHandle), in, v)
}

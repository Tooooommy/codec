package codec

import (
	// "github.com/golang/protobuf/proto"
	"github.com/gogo/protobuf/proto"
)

type Proto struct{}

func (*Proto) Encode(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (*Proto) Decode(in []byte, v interface{}) error {
	return proto.Unmarshal(in, v.(proto.Message))
}

package codec

import (
	"bytes"
	"encoding/gob"
)

type Gob struct{}

func (*Gob) Encode(v interface{}) (out []byte, err error) {
	err = gob.NewEncoder(bytes.NewBuffer(out)).Encode(v)
	return
}

func (*Gob) Decode(in []byte, v interface{}) (err error) {
	err = gob.NewDecoder(bytes.NewReader(in)).Decode(v)
	return err
}

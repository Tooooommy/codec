package codec

type Nop struct{}

func (*Nop) Encode(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}

func (*Nop) Decode(in []byte, v interface{}) error {
	return nil
}

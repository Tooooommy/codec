package codec

import "gopkg.in/yaml.v3"

type Yaml struct {
}

func (*Yaml) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (*Yaml) Decode(in []byte, v interface{}) error {
	return yaml.Unmarshal(in, v)
}

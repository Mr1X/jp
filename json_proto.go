package jp

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
)

// J2P json to protobuf
func J2P(mjson interface{}, mproto proto.Message) error {
	return byte2proto(mjson, mproto)
}

func byte2proto(mjson interface{}, mproto proto.Message) error {
	buf, err := json.Marshal(mjson)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(buf, mproto); err != nil {
		return err
	}
	return nil
}

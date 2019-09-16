package base64

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//Int is a base64 encoded int64
type Int int64

//MarshalJSON handles json marshal
func (b Int) MarshalJSON() ([]byte, error) {
	v := int64(b)
	s := strconv.FormatInt(v, 10)
	buff := []byte(s)
	enc := b64.RawStdEncoding.EncodeToString(buff)
	return json.Marshal(enc)
}

//UnmarshalJSON handles json unmarshal
func (b *Int) UnmarshalJSON(data []byte) error {

	s := string(data)
	s = strings.Trim(s, "\"")
	v, err := decodeFromBase64Int(s)
	if err != nil {
		return fmt.Errorf("Base64Int UnmarshalJSON: %v. Data: %s", err, s)
	}
	*b = v
	return nil
}

func decodeFromBase64Int(s string) (Int, error) {
	dec, err := b64.RawStdEncoding.DecodeString(s)
	if err != nil {
		dec = []byte(s)
	}

	v, err := strconv.ParseInt(string(dec), 10, 64)
	if err != nil {
		return 0, err
	}

	return Int(v), nil
}

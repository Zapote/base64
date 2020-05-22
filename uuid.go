package base64

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

//UUID is a base64 endcoded uuid.UUID
type UUID struct {
	Value uuid.UUID
}

//NewUUID creates a new NewBase64ID
func NewUUID(v uuid.UUID) UUID {
	b := UUID{v}
	return b
}

//NewUUIDFromEncoded creates a new base64.ID from a base64 encoded string
func NewUUIDFromEncoded(s string) UUID {
	v, err := decodeFromBase64ID(s)

	if err != nil {
		return UUID{}
	}

	b := UUID{v}
	return b
}

func (b *UUID) String() string {
	return b.Value.String()
}

//MarshalJSON handles json marshal
func (b UUID) MarshalJSON() ([]byte, error) {
	buff := []byte(b.String())
	s := b64.RawStdEncoding.EncodeToString(buff)
	return json.Marshal(s)
}

//UnmarshalJSON handles json unmarshal
func (b *UUID) UnmarshalJSON(data []byte) error {
	str := string(data)
	v, err := decodeFromBase64ID(str)
	if err != nil {
		return fmt.Errorf("Base64ID UnmarshalJSON: %v. Data: %s", err, str)
	}
	b.Value = v
	return nil
}

//Scan ID
func (b *UUID) Scan(value interface{}) error {
	err := b.Value.Scan(value)

	if err != nil {
		return err
	}

	return nil
}

func decodeFromBase64ID(s string) (uuid.UUID, error) {
	res, err := b64.RawStdEncoding.DecodeString(strings.Trim(s, "\""))
	if err != nil {
		res = []byte(s)
	}
	v, err := uuid.Parse(string(res))
	if err != nil {
		return uuid.UUID{}, err
	}
	return v, nil
}

package base64

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"uuid"
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
func NewUUIDFromEncoded(s string) (UUID, error) {
	v, err := decodeFromBase64ID(s)

	if err != nil {
		return UUID{}, fmt.Errorf("Could not decode: %s", err.Error())
	}

	b := UUID{v}
	return b, nil
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

// Scan implements sql.Scanner. Accepts a UUID string, a 16-byte slice or a
// textual byte slice; nil and empty values leave b untouched.
func (b *UUID) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		if v == "" {
			return nil
		}
		u, err := uuid.Parse(v)
		if err != nil {
			return fmt.Errorf("Scan: %v", err)
		}
		b.Value = u
	case []byte:
		if len(v) == 0 {
			return nil
		}
		if len(v) != 16 {
			return b.Scan(string(v))
		}
		copy(b.Value[:], v)
	default:
		return fmt.Errorf("Scan: unable to scan type %T into UUID", value)
	}
	return nil
}

func decodeFromBase64ID(s string) (uuid.UUID, error) {
	s = strings.Trim(s, "\"")
	res, err := b64.RawStdEncoding.DecodeString(s)
	if err != nil {
		res = []byte(s)
	}
	v, err := uuid.Parse(string(res))
	if err != nil {
		return uuid.UUID{}, err
	}
	return v, nil
}

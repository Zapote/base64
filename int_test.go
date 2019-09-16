package base64

import (
	"testing"
)

func TestBase64IDMarshalJson(t *testing.T) {
	enc := "\"Mjg\""
	var ts Int = 28
	b, err := ts.MarshalJSON()

	if err != nil {
		t.Error(err)
	}

	json := string(b)

	if json != enc {
		t.Errorf("json not correct encoded. Exp (%s), got (%s)", enc, json)
	}
}

func TestBase64IntUnmarshalJSON(t *testing.T) {
	testcases := []struct {
		enc  []byte
		exp  Int
		name string
	}{
		{enc: []byte("\"OQ\""), exp: Int(9), name: "9"},
		{enc: []byte("\"Mjk\""), exp: Int(29), name: "29"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var ts Int

			err := ts.UnmarshalJSON(tc.enc)

			if err != nil {
				t.Error(err)
			}

			if tc.exp != ts {
				t.Errorf("json not correct decoded. Exp (%v), got (%v)", tc.exp, ts)
			}
		})
	}
}

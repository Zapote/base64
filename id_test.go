package base64

import (
	"testing"
	"github.com/google/uuid"
	"github.com/zapote/base64/assert"
)

func TestIDMarshalJson(t *testing.T) {
	enc := "\"MWUyZDNiNzAtODM0Ni00N2E2LTgzNDktYTVlZjAxMTIxZmEy\""
	v, _ := uuid.Parse("1E2D3B70-8346-47A6-8349-A5EF01121FA2")
	ts := ID{
		Value: v,
	}

	b, err := ts.MarshalJSON()

	if err != nil {
		t.Error(err)
	}

	json := string(b)

	if json != enc {
		t.Errorf("json not correct encoded. Exp (%s), got (%s)", enc, json)
	}
}

func TestIDUnmarshalJSON(t *testing.T) {
	v, _ := uuid.Parse("1E2D3B70-8346-47A6-8349-A5EF01121FA2")
	tc := []struct {
		name string
		enc  []byte
		exp  uuid.UUID
	}{
		{"Encoded returns correct UUID", []byte("\"MWUyZDNiNzAtODM0Ni00N2E2LTgzNDktYTVlZjAxMTIxZmEy\""), v},
		{"Not encoded returns same UUID", []byte("\"1E2D3B70-8346-47A6-8349-A5EF01121FA2\""), v},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			act := ID{}
			err := act.UnmarshalJSON(c.enc)
			if err != nil {
				t.Error(err)
			}
			assert.IsEqual(t, act.Value, c.exp)
		})
	}
}

func TestIDScan(t *testing.T) {
	cases := []struct {
		name string
		exp  uuid.UUID
	}{
		{name: "ID 1", exp: uuid.MustParse("1e2D3B70-8346-47A6-8349-A5EF01121FA2")},
		{name: "ID 2", exp: uuid.MustParse("F1959E2E-F2FC-4F4B-9564-A3B101366824")},
	}

	for _, c := range cases {
		t.Run(c.name, func(*testing.T) {
			ts := ID{}
			err := ts.Scan(c.exp.String())

			if err != nil {
				t.Error(err)
			}

			if c.exp != ts.Value {
				t.Errorf("Scan not correct. Exp (%v), got (%v)", c.exp, ts.Value)
			}
		})
	}

}

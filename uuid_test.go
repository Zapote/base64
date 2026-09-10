package base64

import (
	"testing"
	"uuid"

	"github.com/zapote/base64/assert"
)

func TestIDMarshalJson(t *testing.T) {
	enc := "\"MWUyZDNiNzAtODM0Ni00N2E2LTgzNDktYTVlZjAxMTIxZmEy\""
	v, _ := uuid.Parse("1E2D3B70-8346-47A6-8349-A5EF01121FA2")
	ts := UUID{
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
			act := UUID{}
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
			ts := UUID{}
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

func TestIDScanVariants(t *testing.T) {
	exp := uuid.MustParse("1e2d3b70-8346-47a6-8349-a5ef01121fa2")
	raw := exp[:]

	cases := []struct {
		name string
		src  interface{}
		exp  uuid.UUID
	}{
		{"16 raw bytes", raw, exp},
		{"textual bytes", []byte(exp.String()), exp},
		{"empty string leaves nil uuid", "", uuid.UUID{}},
		{"nil leaves nil uuid", nil, uuid.UUID{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts := UUID{}
			if err := ts.Scan(c.src); err != nil {
				t.Fatal(err)
			}
			if ts.Value != c.exp {
				t.Errorf("got %v, want %v", ts.Value, c.exp)
			}
		})
	}

	if err := (&UUID{}).Scan(42); err == nil {
		t.Error("expected error for unsupported type")
	}
}

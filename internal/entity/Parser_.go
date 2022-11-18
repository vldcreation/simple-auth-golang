package entity

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

type Decoder interface{ Decode(v interface{}) error }

type Encoder interface{ Encode(v interface{}) error }

type Parser interface {
	NewDecoder(r io.Reader) Decoder
	NewEncoder(w io.Writer) Encoder
}

type json_ struct{ jsoniter.API }

var (
	JSON Parser = json_{jsoniter.ConfigFastest}
)

func (json json_) NewDecoder(r io.Reader) Decoder { return json.API.NewDecoder(r) }

func (json json_) NewEncoder(w io.Writer) Encoder { return json.API.NewEncoder(w) }

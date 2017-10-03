// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package handlers

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE2a549a6DecodeGithubComKihamoShadowComponentsGrpcInternalHandlers(in *jlexer.Lexer, out *indexHandlerResponseCall) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			out.Result = string(in.String())
		case "error":
			out.Error = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE2a549a6EncodeGithubComKihamoShadowComponentsGrpcInternalHandlers(out *jwriter.Writer, in indexHandlerResponseCall) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Result != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"result\":")
		out.String(string(in.Result))
	}
	if in.Error != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"error\":")
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v indexHandlerResponseCall) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE2a549a6EncodeGithubComKihamoShadowComponentsGrpcInternalHandlers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v indexHandlerResponseCall) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE2a549a6EncodeGithubComKihamoShadowComponentsGrpcInternalHandlers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *indexHandlerResponseCall) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE2a549a6DecodeGithubComKihamoShadowComponentsGrpcInternalHandlers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *indexHandlerResponseCall) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE2a549a6DecodeGithubComKihamoShadowComponentsGrpcInternalHandlers(l, v)
}
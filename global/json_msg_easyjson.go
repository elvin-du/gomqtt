// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package global

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal(in *jlexer.Lexer, out *C2SMsg) {
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
		case "toacc":
			out.ToAcc = string(in.String())
		case "type":
			out.Type = int(in.Int())
		case "msgid":
			out.MsgID = string(in.String())
		case "msg":
			out.Msg = string(in.String())
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
func easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal(out *jwriter.Writer, in C2SMsg) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"toacc\":")
	out.String(string(in.ToAcc))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	out.Int(int(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"msgid\":")
	out.String(string(in.MsgID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"msg\":")
	out.String(string(in.Msg))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v C2SMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v C2SMsg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *C2SMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *C2SMsg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal(l, v)
}
func easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal1(in *jlexer.Lexer, out *Messages) {
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
		case "compress":
			out.Compress = int(in.Int())
		case "data":
			out.Data = string(in.String())
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
func easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal1(out *jwriter.Writer, in Messages) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"compress\":")
	out.Int(int(in.Compress))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"data\":")
	out.String(string(in.Data))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Messages) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Messages) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Messages) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Messages) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal1(l, v)
}

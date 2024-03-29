// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package proto

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

func easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto(in *jlexer.Lexer, out *ModifyAccountSettingsRequestData) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.UserID = string(in.String())
		case "session_id":
			out.SessionID = string(in.String())
		case "new_username":
			out.NewUsername = string(in.String())
		case "new_email":
			out.NewEmail = string(in.String())
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
func easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto(out *jwriter.Writer, in ModifyAccountSettingsRequestData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"session_id\":"
		out.RawString(prefix)
		out.String(string(in.SessionID))
	}
	if in.NewUsername != "" {
		const prefix string = ",\"new_username\":"
		out.RawString(prefix)
		out.String(string(in.NewUsername))
	}
	if in.NewEmail != "" {
		const prefix string = ",\"new_email\":"
		out.RawString(prefix)
		out.String(string(in.NewEmail))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ModifyAccountSettingsRequestData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ModifyAccountSettingsRequestData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ModifyAccountSettingsRequestData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ModifyAccountSettingsRequestData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto(l, v)
}
func easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto1(in *jlexer.Lexer, out *Message) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = Type(in.Int())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				out.Data = in.Bytes()
			}
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
func easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto1(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Type))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Data)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto1(l, v)
}
func easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto2(in *jlexer.Lexer, out *LogoutRequestData) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.UserID = string(in.String())
		case "session_id":
			out.SessionID = string(in.String())
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
func easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto2(out *jwriter.Writer, in LogoutRequestData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"session_id\":"
		out.RawString(prefix)
		out.String(string(in.SessionID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LogoutRequestData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LogoutRequestData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LogoutRequestData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LogoutRequestData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto2(l, v)
}
func easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto3(in *jlexer.Lexer, out *LoginResponseData) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.UserID = string(in.String())
		case "session_id":
			out.SessionID = string(in.String())
		case "expire_at":
			out.ExpireAt = int64(in.Int64())
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
func easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto3(out *jwriter.Writer, in LoginResponseData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"session_id\":"
		out.RawString(prefix)
		out.String(string(in.SessionID))
	}
	{
		const prefix string = ",\"expire_at\":"
		out.RawString(prefix)
		out.Int64(int64(in.ExpireAt))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LoginResponseData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LoginResponseData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LoginResponseData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LoginResponseData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto3(l, v)
}
func easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto4(in *jlexer.Lexer, out *LoginRequestData) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "username":
			out.Username = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto4(out *jwriter.Writer, in LoginRequestData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix[1:])
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LoginRequestData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LoginRequestData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LoginRequestData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LoginRequestData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto4(l, v)
}
func easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto5(in *jlexer.Lexer, out *DeactivateUserRequestData) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.UserID = string(in.String())
		case "session_id":
			out.SessionID = string(in.String())
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
func easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto5(out *jwriter.Writer, in DeactivateUserRequestData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"session_id\":"
		out.RawString(prefix)
		out.String(string(in.SessionID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeactivateUserRequestData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeactivateUserRequestData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeactivateUserRequestData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeactivateUserRequestData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto5(l, v)
}
func easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto6(in *jlexer.Lexer, out *AddFavoritePageRequestData) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.UserID = string(in.String())
		case "session_id":
			out.SessionID = string(in.String())
		case "page_url":
			out.PageURL = string(in.String())
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
func easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto6(out *jwriter.Writer, in AddFavoritePageRequestData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"session_id\":"
		out.RawString(prefix)
		out.String(string(in.SessionID))
	}
	{
		const prefix string = ",\"page_url\":"
		out.RawString(prefix)
		out.String(string(in.PageURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddFavoritePageRequestData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddFavoritePageRequestData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson22b38c74EncodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddFavoritePageRequestData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddFavoritePageRequestData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson22b38c74DecodeGithubComWindscribeAlexandrKarpovichSecureServerChallengeInternalProto6(l, v)
}

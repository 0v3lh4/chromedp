// AUTOGENERATED FILE: easyjson marshaler/unmarshalers.

package serviceworker

import (
	json "encoding/json"
	target "github.com/knq/chromedp/cdp/target"
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

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker(in *jlexer.Lexer, out *Version) {
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
		case "versionId":
			out.VersionID = string(in.String())
		case "registrationId":
			out.RegistrationID = string(in.String())
		case "scriptURL":
			out.ScriptURL = string(in.String())
		case "runningStatus":
			(out.RunningStatus).UnmarshalEasyJSON(in)
		case "status":
			(out.Status).UnmarshalEasyJSON(in)
		case "scriptLastModified":
			out.ScriptLastModified = float64(in.Float64())
		case "scriptResponseTime":
			out.ScriptResponseTime = float64(in.Float64())
		case "controlledClients":
			if in.IsNull() {
				in.Skip()
				out.ControlledClients = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.ControlledClients = make([]target.ID, 0, 4)
				} else {
					out.ControlledClients = []target.ID{}
				}
				for !in.IsDelim(']') {
					var v1 target.ID
					v1 = target.ID(in.String())
					out.ControlledClients = append(out.ControlledClients, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "targetId":
			out.TargetID = target.ID(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker(out *jwriter.Writer, in Version) {
	out.RawByte('{')
	first := true
	_ = first
	if in.VersionID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"versionId\":")
		out.String(string(in.VersionID))
	}
	if in.RegistrationID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"registrationId\":")
		out.String(string(in.RegistrationID))
	}
	if in.ScriptURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"scriptURL\":")
		out.String(string(in.ScriptURL))
	}
	if in.RunningStatus != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"runningStatus\":")
		(in.RunningStatus).MarshalEasyJSON(out)
	}
	if in.Status != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"status\":")
		(in.Status).MarshalEasyJSON(out)
	}
	if in.ScriptLastModified != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"scriptLastModified\":")
		out.Float64(float64(in.ScriptLastModified))
	}
	if in.ScriptResponseTime != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"scriptResponseTime\":")
		out.Float64(float64(in.ScriptResponseTime))
	}
	if len(in.ControlledClients) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"controlledClients\":")
		if in.ControlledClients == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.ControlledClients {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.TargetID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"targetId\":")
		out.String(string(in.TargetID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Version) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Version) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Version) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Version) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker1(in *jlexer.Lexer, out *UpdateRegistrationParams) {
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
		case "scopeURL":
			out.ScopeURL = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker1(out *jwriter.Writer, in UpdateRegistrationParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"scopeURL\":")
	out.String(string(in.ScopeURL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateRegistrationParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateRegistrationParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateRegistrationParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateRegistrationParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker2(in *jlexer.Lexer, out *UnregisterParams) {
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
		case "scopeURL":
			out.ScopeURL = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker2(out *jwriter.Writer, in UnregisterParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"scopeURL\":")
	out.String(string(in.ScopeURL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UnregisterParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UnregisterParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UnregisterParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UnregisterParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker3(in *jlexer.Lexer, out *StopWorkerParams) {
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
		case "versionId":
			out.VersionID = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker3(out *jwriter.Writer, in StopWorkerParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"versionId\":")
	out.String(string(in.VersionID))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StopWorkerParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StopWorkerParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StopWorkerParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StopWorkerParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker4(in *jlexer.Lexer, out *StartWorkerParams) {
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
		case "scopeURL":
			out.ScopeURL = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker4(out *jwriter.Writer, in StartWorkerParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"scopeURL\":")
	out.String(string(in.ScopeURL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StartWorkerParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StartWorkerParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StartWorkerParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StartWorkerParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker5(in *jlexer.Lexer, out *SkipWaitingParams) {
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
		case "scopeURL":
			out.ScopeURL = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker5(out *jwriter.Writer, in SkipWaitingParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"scopeURL\":")
	out.String(string(in.ScopeURL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SkipWaitingParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SkipWaitingParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SkipWaitingParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SkipWaitingParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker5(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker6(in *jlexer.Lexer, out *SetForceUpdateOnPageLoadParams) {
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
		case "forceUpdateOnPageLoad":
			out.ForceUpdateOnPageLoad = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker6(out *jwriter.Writer, in SetForceUpdateOnPageLoadParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"forceUpdateOnPageLoad\":")
	out.Bool(bool(in.ForceUpdateOnPageLoad))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetForceUpdateOnPageLoadParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetForceUpdateOnPageLoadParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetForceUpdateOnPageLoadParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetForceUpdateOnPageLoadParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker6(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker7(in *jlexer.Lexer, out *Registration) {
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
		case "registrationId":
			out.RegistrationID = string(in.String())
		case "scopeURL":
			out.ScopeURL = string(in.String())
		case "isDeleted":
			out.IsDeleted = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker7(out *jwriter.Writer, in Registration) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RegistrationID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"registrationId\":")
		out.String(string(in.RegistrationID))
	}
	if in.ScopeURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"scopeURL\":")
		out.String(string(in.ScopeURL))
	}
	if in.IsDeleted {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"isDeleted\":")
		out.Bool(bool(in.IsDeleted))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Registration) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Registration) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Registration) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Registration) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker7(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker8(in *jlexer.Lexer, out *InspectWorkerParams) {
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
		case "versionId":
			out.VersionID = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker8(out *jwriter.Writer, in InspectWorkerParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"versionId\":")
	out.String(string(in.VersionID))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InspectWorkerParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InspectWorkerParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InspectWorkerParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InspectWorkerParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker8(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker9(in *jlexer.Lexer, out *EventWorkerVersionUpdated) {
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
		case "versions":
			if in.IsNull() {
				in.Skip()
				out.Versions = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Versions = make([]*Version, 0, 8)
				} else {
					out.Versions = []*Version{}
				}
				for !in.IsDelim(']') {
					var v4 *Version
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Version)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Versions = append(out.Versions, v4)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker9(out *jwriter.Writer, in EventWorkerVersionUpdated) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Versions) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"versions\":")
		if in.Versions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Versions {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventWorkerVersionUpdated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventWorkerVersionUpdated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventWorkerVersionUpdated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventWorkerVersionUpdated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker9(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker10(in *jlexer.Lexer, out *EventWorkerRegistrationUpdated) {
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
		case "registrations":
			if in.IsNull() {
				in.Skip()
				out.Registrations = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Registrations = make([]*Registration, 0, 8)
				} else {
					out.Registrations = []*Registration{}
				}
				for !in.IsDelim(']') {
					var v7 *Registration
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(Registration)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.Registrations = append(out.Registrations, v7)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker10(out *jwriter.Writer, in EventWorkerRegistrationUpdated) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Registrations) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"registrations\":")
		if in.Registrations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Registrations {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventWorkerRegistrationUpdated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventWorkerRegistrationUpdated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventWorkerRegistrationUpdated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventWorkerRegistrationUpdated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker10(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker11(in *jlexer.Lexer, out *EventWorkerErrorReported) {
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
		case "errorMessage":
			if in.IsNull() {
				in.Skip()
				out.ErrorMessage = nil
			} else {
				if out.ErrorMessage == nil {
					out.ErrorMessage = new(ErrorMessage)
				}
				(*out.ErrorMessage).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker11(out *jwriter.Writer, in EventWorkerErrorReported) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ErrorMessage != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"errorMessage\":")
		if in.ErrorMessage == nil {
			out.RawString("null")
		} else {
			(*in.ErrorMessage).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventWorkerErrorReported) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventWorkerErrorReported) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventWorkerErrorReported) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventWorkerErrorReported) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker11(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker12(in *jlexer.Lexer, out *ErrorMessage) {
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
		case "errorMessage":
			out.ErrorMessage = string(in.String())
		case "registrationId":
			out.RegistrationID = string(in.String())
		case "versionId":
			out.VersionID = string(in.String())
		case "sourceURL":
			out.SourceURL = string(in.String())
		case "lineNumber":
			out.LineNumber = int64(in.Int64())
		case "columnNumber":
			out.ColumnNumber = int64(in.Int64())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker12(out *jwriter.Writer, in ErrorMessage) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ErrorMessage != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"errorMessage\":")
		out.String(string(in.ErrorMessage))
	}
	if in.RegistrationID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"registrationId\":")
		out.String(string(in.RegistrationID))
	}
	if in.VersionID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"versionId\":")
		out.String(string(in.VersionID))
	}
	if in.SourceURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"sourceURL\":")
		out.String(string(in.SourceURL))
	}
	if in.LineNumber != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"lineNumber\":")
		out.Int64(int64(in.LineNumber))
	}
	if in.ColumnNumber != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"columnNumber\":")
		out.Int64(int64(in.ColumnNumber))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker12(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker13(in *jlexer.Lexer, out *EnableParams) {
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker13(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker13(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker14(in *jlexer.Lexer, out *DispatchSyncEventParams) {
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
		case "origin":
			out.Origin = string(in.String())
		case "registrationId":
			out.RegistrationID = string(in.String())
		case "tag":
			out.Tag = string(in.String())
		case "lastChance":
			out.LastChance = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker14(out *jwriter.Writer, in DispatchSyncEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"origin\":")
	out.String(string(in.Origin))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"registrationId\":")
	out.String(string(in.RegistrationID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tag\":")
	out.String(string(in.Tag))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"lastChance\":")
	out.Bool(bool(in.LastChance))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DispatchSyncEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DispatchSyncEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DispatchSyncEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DispatchSyncEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker14(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker15(in *jlexer.Lexer, out *DisableParams) {
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker15(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker15(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker15(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker15(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker15(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker16(in *jlexer.Lexer, out *DeliverPushMessageParams) {
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
		case "origin":
			out.Origin = string(in.String())
		case "registrationId":
			out.RegistrationID = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker16(out *jwriter.Writer, in DeliverPushMessageParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"origin\":")
	out.String(string(in.Origin))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"registrationId\":")
	out.String(string(in.RegistrationID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"data\":")
	out.String(string(in.Data))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeliverPushMessageParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker16(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeliverPushMessageParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpServiceworker16(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeliverPushMessageParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker16(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeliverPushMessageParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpServiceworker16(l, v)
}

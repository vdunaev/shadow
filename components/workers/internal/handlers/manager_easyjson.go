// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package handlers

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers(in *jlexer.Lexer, out *managerHandlerStats) {
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
		case "tasks_count":
			out.TasksCount = uint64(in.Uint64())
		case "tasks_waiting_count":
			out.TasksWaitingCount = uint64(in.Uint64())
		case "workers_count":
			out.WorkersCount = uint64(in.Uint64())
		case "workers_waiting_count":
			out.WorkersWaitingCount = uint64(in.Uint64())
		case "listeners_count":
			out.ListenersCount = uint64(in.Uint64())
		case "workers":
			if in.IsNull() {
				in.Skip()
				out.Workers = nil
			} else {
				in.Delim('[')
				if out.Workers == nil {
					if !in.IsDelim(']') {
						out.Workers = make([]managerHandlerItemWorker, 0, 1)
					} else {
						out.Workers = []managerHandlerItemWorker{}
					}
				} else {
					out.Workers = (out.Workers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 managerHandlerItemWorker
					(v1).UnmarshalEasyJSON(in)
					out.Workers = append(out.Workers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]managerHandlerItemTask, 0, 1)
					} else {
						out.Tasks = []managerHandlerItemTask{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v2 managerHandlerItemTask
					(v2).UnmarshalEasyJSON(in)
					out.Tasks = append(out.Tasks, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "listeners":
			if in.IsNull() {
				in.Skip()
				out.Listeners = nil
			} else {
				in.Delim('[')
				if out.Listeners == nil {
					if !in.IsDelim(']') {
						out.Listeners = make([]managerHandlerItemListener, 0, 1)
					} else {
						out.Listeners = []managerHandlerItemListener{}
					}
				} else {
					out.Listeners = (out.Listeners)[:0]
				}
				for !in.IsDelim(']') {
					var v3 managerHandlerItemListener
					(v3).UnmarshalEasyJSON(in)
					out.Listeners = append(out.Listeners, v3)
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers(out *jwriter.Writer, in managerHandlerStats) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tasks_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.TasksCount))
	}
	{
		const prefix string = ",\"tasks_waiting_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.TasksWaitingCount))
	}
	{
		const prefix string = ",\"workers_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.WorkersCount))
	}
	{
		const prefix string = ",\"workers_waiting_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.WorkersWaitingCount))
	}
	{
		const prefix string = ",\"listeners_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ListenersCount))
	}
	{
		const prefix string = ",\"workers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Workers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Workers {
				if v4 > 0 {
					out.RawByte(',')
				}
				(v5).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tasks\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Tasks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Tasks {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"listeners\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Listeners == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Listeners {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(in *jlexer.Lexer, out *managerHandlerResponseSuccess) {
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(out *jwriter.Writer, in managerHandlerResponseSuccess) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Result))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerResponseSuccess) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerResponseSuccess) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerResponseSuccess) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerResponseSuccess) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(in *jlexer.Lexer, out *managerHandlerItemWorker) {
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
		case "id":
			out.Id = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "status":
			out.Status = string(in.String())
		case "task":
			if in.IsNull() {
				in.Skip()
				out.Task = nil
			} else {
				if out.Task == nil {
					out.Task = new(managerHandlerItemTask)
				}
				(*out.Task).UnmarshalEasyJSON(in)
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(out *jwriter.Writer, in managerHandlerItemWorker) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Created).MarshalJSON())
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"task\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Task == nil {
			out.RawString("null")
		} else {
			(*in.Task).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerItemWorker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerItemWorker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerItemWorker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerItemWorker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(in *jlexer.Lexer, out *managerHandlerItemTask) {
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
		case "id":
			out.Id = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "priority":
			out.Priority = int64(in.Int64())
		case "repeats":
			out.Repeats = int64(in.Int64())
		case "repeat_interval":
			out.RepeatInterval = time.Duration(in.Int64())
		case "timeout":
			out.Timeout = time.Duration(in.Int64())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "started_at":
			if in.IsNull() {
				in.Skip()
				out.StartedAt = nil
			} else {
				if out.StartedAt == nil {
					out.StartedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.StartedAt).UnmarshalJSON(data))
				}
			}
		case "status":
			out.Status = string(in.String())
		case "attempts":
			out.Attempts = int64(in.Int64())
		case "allow_start_at":
			if in.IsNull() {
				in.Skip()
				out.AllowStartAt = nil
			} else {
				if out.AllowStartAt == nil {
					out.AllowStartAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.AllowStartAt).UnmarshalJSON(data))
				}
			}
		case "first_started_at":
			if in.IsNull() {
				in.Skip()
				out.FirstStartedAt = nil
			} else {
				if out.FirstStartedAt == nil {
					out.FirstStartedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.FirstStartedAt).UnmarshalJSON(data))
				}
			}
		case "last_started_at":
			if in.IsNull() {
				in.Skip()
				out.LastStartedAt = nil
			} else {
				if out.LastStartedAt == nil {
					out.LastStartedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastStartedAt).UnmarshalJSON(data))
				}
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(out *jwriter.Writer, in managerHandlerItemTask) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"priority\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Priority))
	}
	{
		const prefix string = ",\"repeats\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Repeats))
	}
	{
		const prefix string = ",\"repeat_interval\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RepeatInterval))
	}
	{
		const prefix string = ",\"timeout\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Timeout))
	}
	{
		const prefix string = ",\"created_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"started_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StartedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.StartedAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"attempts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Attempts))
	}
	{
		const prefix string = ",\"allow_start_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.AllowStartAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.AllowStartAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"first_started_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.FirstStartedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.FirstStartedAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"last_started_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.LastStartedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.LastStartedAt).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerItemTask) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerItemTask) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerItemTask) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerItemTask) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(in *jlexer.Lexer, out *managerHandlerItemListener) {
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
		case "id":
			out.Id = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "locked":
			out.Locked = bool(in.Bool())
		case "events":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Events = make(map[string]int64)
				} else {
					out.Events = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v10 int64
					v10 = int64(in.Int64())
					(out.Events)[key] = v10
					in.WantComma()
				}
				in.Delim('}')
			}
		case "fires":
			out.Fires = int64(in.Int64())
		case "first_fired_at":
			if in.IsNull() {
				in.Skip()
				out.FirstFiredAt = nil
			} else {
				if out.FirstFiredAt == nil {
					out.FirstFiredAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.FirstFiredAt).UnmarshalJSON(data))
				}
			}
		case "last_fired_at":
			if in.IsNull() {
				in.Skip()
				out.LastFiredAt = nil
			} else {
				if out.LastFiredAt == nil {
					out.LastFiredAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastFiredAt).UnmarshalJSON(data))
				}
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(out *jwriter.Writer, in managerHandlerItemListener) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"locked\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Locked))
	}
	{
		const prefix string = ",\"events\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Events == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v11First := true
			for v11Name, v11Value := range in.Events {
				if v11First {
					v11First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v11Name))
				out.RawByte(':')
				out.Int64(int64(v11Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"fires\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Fires))
	}
	{
		const prefix string = ",\"first_fired_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.FirstFiredAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.FirstFiredAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"last_fired_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.LastFiredAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.LastFiredAt).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerItemListener) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerItemListener) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerItemListener) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerItemListener) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(l, v)
}

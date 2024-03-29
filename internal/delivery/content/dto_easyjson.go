// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package content

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

func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent(in *jlexer.Lexer, out *seriesDTO) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "content":
			(out.Content).UnmarshalEasyJSON(in)
		case "episodes":
			if in.IsNull() {
				in.Skip()
				out.Episodes = nil
			} else {
				in.Delim('[')
				if out.Episodes == nil {
					if !in.IsDelim(']') {
						out.Episodes = make([]episodeDTO, 0, 1)
					} else {
						out.Episodes = []episodeDTO{}
					}
				} else {
					out.Episodes = (out.Episodes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 episodeDTO
					(v1).UnmarshalEasyJSON(in)
					out.Episodes = append(out.Episodes, v1)
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent(out *jwriter.Writer, in seriesDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		(in.Content).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"episodes\":"
		out.RawString(prefix)
		if in.Episodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Episodes {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v seriesDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v seriesDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *seriesDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *seriesDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent1(in *jlexer.Lexer, out *selectionDTO) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent1(out *jwriter.Writer, in selectionDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v selectionDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v selectionDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *selectionDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *selectionDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent1(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent2(in *jlexer.Lexer, out *roleDTO) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent2(out *jwriter.Writer, in roleDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v roleDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v roleDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *roleDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *roleDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent2(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent3(in *jlexer.Lexer, out *personRolesDTO) {
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
		case "person":
			(out.Person).UnmarshalEasyJSON(in)
		case "role":
			(out.Role).UnmarshalEasyJSON(in)
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent3(out *jwriter.Writer, in personRolesDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"person\":"
		out.RawString(prefix[1:])
		(in.Person).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		(in.Role).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v personRolesDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v personRolesDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *personRolesDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *personRolesDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent3(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent4(in *jlexer.Lexer, out *personDTO) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent4(out *jwriter.Writer, in personDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v personDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v personDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *personDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *personDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent4(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent5(in *jlexer.Lexer, out *genreDTO) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent5(out *jwriter.Writer, in genreDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v genreDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v genreDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *genreDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *genreDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent5(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent6(in *jlexer.Lexer, out *filmDTO) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "content_url":
			out.ContentURL = string(in.String())
		case "content":
			(out.Content).UnmarshalEasyJSON(in)
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent6(out *jwriter.Writer, in filmDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"content_url\":"
		out.RawString(prefix)
		out.String(string(in.ContentURL))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		(in.Content).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v filmDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v filmDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *filmDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *filmDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent6(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent7(in *jlexer.Lexer, out *episodeDTO) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "season_num":
			out.SeasonNum = uint32(in.Uint32())
		case "episode_num":
			out.EpisodeNum = uint32(in.Uint32())
		case "content_url":
			out.ContentURL = string(in.String())
		case "preview_url":
			out.PreviewURL = string(in.String())
		case "title":
			if in.IsNull() {
				in.Skip()
				out.Title = nil
			} else {
				if out.Title == nil {
					out.Title = new(string)
				}
				*out.Title = string(in.String())
			}
		case "release_date":
			if in.IsNull() {
				in.Skip()
				out.ReleaseDate = nil
			} else {
				if out.ReleaseDate == nil {
					out.ReleaseDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ReleaseDate).UnmarshalJSON(data))
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent7(out *jwriter.Writer, in episodeDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"season_num\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.SeasonNum))
	}
	{
		const prefix string = ",\"episode_num\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.EpisodeNum))
	}
	{
		const prefix string = ",\"content_url\":"
		out.RawString(prefix)
		out.String(string(in.ContentURL))
	}
	{
		const prefix string = ",\"preview_url\":"
		out.RawString(prefix)
		out.String(string(in.PreviewURL))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		if in.Title == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Title))
		}
	}
	{
		const prefix string = ",\"release_date\":"
		out.RawString(prefix)
		if in.ReleaseDate == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.ReleaseDate).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v episodeDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v episodeDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *episodeDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *episodeDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent7(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent8(in *jlexer.Lexer, out *countryDTO) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent8(out *jwriter.Writer, in countryDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v countryDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v countryDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *countryDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *countryDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent8(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent9(in *jlexer.Lexer, out *contentDTO) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "rating":
			out.Rating = float64(in.Float64())
		case "sum_ratings":
			out.SumRatings = float64(in.Float64())
		case "count_ratings":
			out.CountRatings = uint64(in.Uint64())
		case "year":
			out.Year = int(in.Int())
		case "is_free":
			out.IsFree = bool(in.Bool())
		case "age_limit":
			out.AgeLimit = int(in.Int())
		case "trailer_url":
			out.TrailerURL = string(in.String())
		case "preview_url":
			out.PreviewURL = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "persons_roles":
			if in.IsNull() {
				in.Skip()
				out.PersonsRoles = nil
			} else {
				in.Delim('[')
				if out.PersonsRoles == nil {
					if !in.IsDelim(']') {
						out.PersonsRoles = make([]personRolesDTO, 0, 1)
					} else {
						out.PersonsRoles = []personRolesDTO{}
					}
				} else {
					out.PersonsRoles = (out.PersonsRoles)[:0]
				}
				for !in.IsDelim(']') {
					var v4 personRolesDTO
					(v4).UnmarshalEasyJSON(in)
					out.PersonsRoles = append(out.PersonsRoles, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "genres":
			if in.IsNull() {
				in.Skip()
				out.Genres = nil
			} else {
				in.Delim('[')
				if out.Genres == nil {
					if !in.IsDelim(']') {
						out.Genres = make([]genreDTO, 0, 2)
					} else {
						out.Genres = []genreDTO{}
					}
				} else {
					out.Genres = (out.Genres)[:0]
				}
				for !in.IsDelim(']') {
					var v5 genreDTO
					(v5).UnmarshalEasyJSON(in)
					out.Genres = append(out.Genres, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "selections":
			if in.IsNull() {
				in.Skip()
				out.Selections = nil
			} else {
				in.Delim('[')
				if out.Selections == nil {
					if !in.IsDelim(']') {
						out.Selections = make([]selectionDTO, 0, 2)
					} else {
						out.Selections = []selectionDTO{}
					}
				} else {
					out.Selections = (out.Selections)[:0]
				}
				for !in.IsDelim(']') {
					var v6 selectionDTO
					(v6).UnmarshalEasyJSON(in)
					out.Selections = append(out.Selections, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "countries":
			if in.IsNull() {
				in.Skip()
				out.Countries = nil
			} else {
				in.Delim('[')
				if out.Countries == nil {
					if !in.IsDelim(']') {
						out.Countries = make([]countryDTO, 0, 2)
					} else {
						out.Countries = []countryDTO{}
					}
				} else {
					out.Countries = (out.Countries)[:0]
				}
				for !in.IsDelim(']') {
					var v7 countryDTO
					(v7).UnmarshalEasyJSON(in)
					out.Countries = append(out.Countries, v7)
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent9(out *jwriter.Writer, in contentDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float64(float64(in.Rating))
	}
	{
		const prefix string = ",\"sum_ratings\":"
		out.RawString(prefix)
		out.Float64(float64(in.SumRatings))
	}
	{
		const prefix string = ",\"count_ratings\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.CountRatings))
	}
	{
		const prefix string = ",\"year\":"
		out.RawString(prefix)
		out.Int(int(in.Year))
	}
	{
		const prefix string = ",\"is_free\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsFree))
	}
	{
		const prefix string = ",\"age_limit\":"
		out.RawString(prefix)
		out.Int(int(in.AgeLimit))
	}
	{
		const prefix string = ",\"trailer_url\":"
		out.RawString(prefix)
		out.String(string(in.TrailerURL))
	}
	{
		const prefix string = ",\"preview_url\":"
		out.RawString(prefix)
		out.String(string(in.PreviewURL))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"persons_roles\":"
		out.RawString(prefix)
		if in.PersonsRoles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.PersonsRoles {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"genres\":"
		out.RawString(prefix)
		if in.Genres == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Genres {
				if v10 > 0 {
					out.RawByte(',')
				}
				(v11).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"selections\":"
		out.RawString(prefix)
		if in.Selections == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Selections {
				if v12 > 0 {
					out.RawByte(',')
				}
				(v13).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"countries\":"
		out.RawString(prefix)
		if in.Countries == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Countries {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v contentDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v contentDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *contentDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *contentDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20231ContentDealersInternalDeliveryContent9(l, v)
}

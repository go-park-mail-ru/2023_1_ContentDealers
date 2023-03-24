package session

type sessionDTO struct {
	UserID          uint64 `json:"user_id"`
	ExpiresAtString string `json:"expires_at"`
}

// `json:"expires_at"`
// // Implement Marshaler and Unmarshaler interface
// func (j *JsonTime) UnmarshalJSON(b []byte) error {
// 	s := strings.Trim(string(b), "\"")
// 	t, err := time.Parse("2021-02-18T21:54:42.123Z", s)
// 	if err != nil {
// 		return err
// 	}
// 	*j = JsonTime(t)
// 	return nil
// }

// func (j JsonTime) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(time.Time(j))
// }

// // Maybe a Format function for printing your date
// func (j JsonTime) Format(s string) string {
// 	t := time.Time(j)
// 	return t.Format(s)
// }

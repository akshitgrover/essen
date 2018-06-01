package essen

import (
	"encoding/json"
)

func (r Response) Json(status int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	r.Res.Header().Set("Content-Type", "application/json")
	r.Res.WriteHeader(status)
	r.Res.Write(b)
	return nil
}

package model

type PreviousNext struct {
	ID   *int64 `json:"id,omitempty"`
	Prev *int64 `json:"prev,omitempty"`
	Next *int64 `json:"next,omitempty"`
}

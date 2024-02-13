package idam

import "time"

type IdamUserType uint8

const (
	StandardUserType IdamUserType = iota
)

type User struct {
	Id           string       `json:"id"`
	Username     string       `json:"username"`
	Email        string       `json:"email"`
	Verified     bool         `json:"verified"`
	Type         IdamUserType `json:"type"`
	Provider     string       `json:"provider"`
	CreatedAtUTC time.Time    `json:"created_at_utc"`
	Features     []string     `json:"features"`
}

package models

type Type string

const (
	TypeFriend   Type = "FRIEND"
	TypeBlocked  Type = "BLOCKING"
	TypeFollowed Type = "FOLLOWING"
)

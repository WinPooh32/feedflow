package previlegies

//Role enum
type Role int8

const (
	Guest Role = iota
	User
	Admin
)

package models

// User defines the data model for individual user records
type User struct {
	Id        uint32 `json:"id"`
	Firstname string `json:"name_first"`
	Lastname  string `json:"name_last"`
}

// UserQueryOptions and its associated setting functions
// provide configurable query filters
// when interacting with the data access layer
type UserQueryOptions struct {
	Id        *uint32
	Firstname *string
	Lastname  *string
}

type UserOption func(*UserQueryOptions)

func WithUserId(id uint32) UserOption {
	return func(uo *UserQueryOptions) {
		uo.Id = &id
	}
}

func WithUserName(first, last string) UserOption {
	return func(uo *UserQueryOptions) {
		uo.Firstname = &first
		uo.Lastname = &last
	}
}

// NewUserQueryOptions receives a given option-setting function
// This could be made variadic, looping over WithOption functions,
// should the data access layer support it
func NewUserQueryOptions(opt UserOption) UserQueryOptions {
	uo := &UserQueryOptions{}
	opt(uo)
	return *uo
}

package main

import "time"

// User defines user detail struct
type User struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	FullName   string    `json:"fullname"`
	Email      string    `json:"email"`
	ProfilePic string    `json:"profilePic"`
	CreateAt   time.Time `json:"createAt" time_format:"2017-08-30T13:35:00Z"`
	IsActive   bool      `json:"isActive"`
}

// Repository defines repository interface
type Repository interface {
	AddNewUser(*User) error
	GetAllUsers() ([]*User, error)
	GetUserByID(int) (*User, error)
	UpdateUser(*User) error
	ReplaceUser(*User) error
	DeleteUserByID(int) error
}

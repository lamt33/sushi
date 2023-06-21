package dto

import uuid "github.com/satori/go.uuid"

type User struct {
	UserID    uuid.UUID
	UserName  string
	FirstName string
	LastName  string
	Email     string
}

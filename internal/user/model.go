package user

import "time"

type User struct {
	ID             string    `json:"id"`
	FullName       string    `json:"fullName"`
	Email          string    `json:"email"`
	Role           string    `json:"role"`
	Status         string    `json:"status"`
	ProfilePicture string    `json:"profilePicture"`
	CreatedAt      time.Time `json:"createdAt"`
}

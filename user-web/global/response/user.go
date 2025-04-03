package response

import "time"

type UserResponse struct {
	Id       int32     `json:"id"`
	Password string    `json:"password"`
	Mobile   string    `json:"mobile"`
	Nickname string    `json:"nickname"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	Role     int32     `json:"role"`
}

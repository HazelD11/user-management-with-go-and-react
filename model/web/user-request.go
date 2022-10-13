package web

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	ID       uint64
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

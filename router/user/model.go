package user

type LoginRequest struct {
	UserId   int64  `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token" binding:"required"`
}

type RegisterRequest struct {
	UserId   int64  `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	UserId int64 `json:"user_id" binding:"required"`
}

type InfoRequest struct {
}

type InfoResponse struct {
	UserId   int64  `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

package vo

type LoginRequestVo struct {
	Username string `json:"username" validate:"required,max=32"`
	Password string `json:"password" validate:"required,max=32"`
}

type LoginResponseVo struct {
}

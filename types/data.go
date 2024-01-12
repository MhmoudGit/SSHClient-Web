package types

type LoginConfig struct {
	User     string `form:"username"`
	Password string `form:"password"`
	Addr     string `form:"address"`
}

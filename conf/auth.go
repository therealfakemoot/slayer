package conf

type Full struct {
	Auth AuthOptions
}

type AuthOptions struct {
	User     string
	Token    string
	Password string
}

package API

type User struct {
    Login string `json:"login"`
    Email string `json:"email"`
    Pass  string `json:"password"`
}

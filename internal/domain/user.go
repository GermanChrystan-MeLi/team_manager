package domain

type User struct {
	ID             string  `json:"id"`
	LdapUser       string  `json:"ldap_user"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Email          string  `json:"email"`
	MeliSite       string  `json:"meli_site"`
	IsAdmin        bool    `json:"is_admin"`
	HashedPassword string  `json:"hashed_password"`
	Account        float32 `json:"account"`
}

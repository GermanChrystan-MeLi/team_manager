package dto

type UserData struct {
	LdapUser  string  `json:"ldap_user"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	MeliSite  string  `json:"meli_site"`
	Account   float32 `json:"account"`
}

type UserRegister struct {
	LdapUser  string `json:"ldap_user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	MeliSite  string `json:"meli_site"`
	Password  string `json:"password"`
}

type UserEdit struct {
	ID        string `json:"id"`
	LdapUser  string `json:"ldap_user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	MeliSite  string `json:"meli_site"`
}

type UserLogin struct {
	EmailOrUsername string `json:"email_or_username"`
	Password        string `json:"password"`
}

type FullProfile struct {
	UserData UserData `json:"user_data"`
	Team     Team     `json:"team"`
	// SessionID string   `json:"session_id"`
}

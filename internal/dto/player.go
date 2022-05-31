package dto

type BasicPlayerDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Country   int    `json:"country"`
	Injured   bool   `json:"injured"`
	Suspended bool   `json:"suspended"`
}

type PlayerStatsData struct {
	PlayerID    int `json:"player_id"`
	Appearances int `json:"appearances"`
	Goals       int `json:"goals"`
	YellowCards int `json:"yellow_cards"`
	RedCards    int `json:"red_cards"`
	Assists     int `json:"assists"`
	Passes      int `json:"passes"`
	Suspensions int `json:"suspensions"`
	Injuries    int `json:"injuries"`
}

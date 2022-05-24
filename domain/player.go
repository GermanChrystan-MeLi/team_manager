package domain

type Player struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Country   string `json:"country"`

	Height float32 `json:"height"`
	Weight float32 `json:"weight"`

	Appearances int `json:"appearances"`
	Goals       int `json:"goals"`
	YellowCards int `json:"yellow_cards"`
	RedCards    int `json:"red_cards"`
	Assists     int `json:"assists"`
	Passes      int `json:"passes"`

	Speed        int `json:"speed"`
	Agility      int `json:"agility"`
	Endurance    int `json:"endurance"`
	Strength     int `json:"strength"`
	GameReading  int `json:"game_reading"`
	Leadership   int `json:"leadership"`
	ShotAccuracy int `json:"shot_accuracy"`
	PassAccuracy int `json:"pass_accuracy"`
	Creativity   int `json:"creativity"`
	GoalKeeping  int `json:"goal_keeping"`
	BallHandling int `json:"ball_handling"`
}

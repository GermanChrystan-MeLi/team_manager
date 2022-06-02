package domain

type PlayerBaseSkillsData struct {
	ID                  string `json:"id"`
	PlayerID            string `json:"player_id"`
	Charisma            int    `json:"charisma"`
	PlauyerIntelligence int    `json:"intelligence"`
	Endurance           int    `json:"endurance"`
	Accuracy            int    `json:"accuracy"`
	Strength            int    `json:"strength"`
	Agility             int    `json:"agility"`
	BallHandling        int    `json:"ball_handling"`
	Blocking            int    `json:"blocking"`
}

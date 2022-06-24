package stats

var DerivedStatsMap = map[string][]string{
	"Goal/PassProne":         {"strength", "accuracy"},
	"SuccesfulEvasionChance": {"ball_handling", "agility"},
	"CriticalMove":           {"ball_handling", "intelligence"}, // Suma chances al próximo jugador del próximo equipo que tome el balón
	"GoalKeeping":            {"ball_handling", "blocking"},
	"MinPlayed":              {"endurance", "agility"},
	"SuccesfulBlockChance":   {"intelligence", "blocking"},
	"Popularity":             {"charisma", "accuracy"},
	"RecoveryTime":           {"endurance", "strength"},
}

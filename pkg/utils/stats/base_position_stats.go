package stats

var GoalKeeper = map[string]int{
	"Goal/PassProne":         6,
	"SuccesfulEvasionChance": 0,
	"CriticalMove":           3,
	"GoalKeeping":            7,
	"MinPlayed":              5,
	"SuccesfulBlockChance":   4,
	"Popularity":             2,
	"RecoveryTime":           1,
}

var Defender = map[string]int{
	"SuccesfulBlockChance":   7,
	"MinPlayed":              6,
	"GoalKeeping":            5,
	"RecoveryTime":           4,
	"CriticalMove":           3,
	"Goal/PassProne":         2,
	"SuccesfulEvasionChance": 1,
	"Popularity":             0,
}

var MidFielder = map[string]int{
	"SuccesfulEvasionChance": 4,
	"CriticalMove":           4,
	"SuccesfulBlockChance":   4,
	"MinPlayed":              4,
	"RecoveryTime":           4,
	"Goal/PassProne":         4,
	"Popularity":             4,
	"GoalKeeping":            0,
}

var Forward = map[string]int{
	"Goal/PassProne":         7,
	"Popularity":             6,
	"SuccesfulEvasionChance": 5,
	"CriticalMove":           4,
	"MinPlayed":              3,
	"RecoveryTime":           2,
	"SuccesfulBlockChance":   1,
	"GoalKeeping":            0,
}

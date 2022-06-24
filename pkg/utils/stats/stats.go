package stats

import (
	"errors"
	"math/rand"
	"time"

	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/constants"
)

func DerivedToRandomBaseStat(derivedStat string, points int) map[string]int {
	rand.Seed(time.Now().UTC().UnixNano())
	// Partitioning the points between base stats
	partition := rand.Intn((points-1)-1+1) + 1

	result := make(map[string]int)
	result[DerivedStatsMap[derivedStat][0]] = partition
	result[DerivedStatsMap[derivedStat][1]] = points - partition
	return result
}

func CreateBaseStats(basePosition map[string]int) map[string]int {
	result := map[string]int{
		"endurance":     0,
		"strength":      0,
		"charisma":      0,
		"accuracy":      0,
		"intelligence":  0,
		"blocking":      0,
		"agility":       0,
		"ball_handling": 0,
	}

	for key1, value1 := range basePosition {
		receivedIncrement := DerivedToRandomBaseStat(key1, value1)
		for key2, value2 := range receivedIncrement {
			result[key2] += value2
		}
	}
	return result
}

func CreateStatsFromBasePosition(position constants.BasePosition) (map[string]int, error) {
	switch position {
	case constants.GoalKeeper:
		return CreateBaseStats(GoalKeeper), nil
	case constants.Defenders:
		return CreateBaseStats(Defender), nil
	case constants.Midfielders:
		return CreateBaseStats(MidFielder), nil
	case constants.Forwards:
		return CreateBaseStats(Forward), nil
	default:
		return map[string]int{}, errors.New("position not valid")
	}
}

/*
GoalKeeper
	GoalKeeping 7
	Goal/PassProne 6
	Min.Played 5
	SuccesfulBlockChance 4
	CriticalMove 3
	Popularity 2
	RecoveryTime 1
	SuccesfulEvasionChance 0

Midfielder
	SuccesfulEvasionChance 4
	CriticalMove 4
	SuccesfulBlockChance 4
	Min.Played 4
	RecoveryTime 4
	Goal/PassProne 4
	Popularity 4
	GoalKeeping 0

Defender
	SuccesfulBlockChance 7
	Min.Played 6
	GoalKeeping 5
	RecoveryTime 4
	CriticalMove 3
	Goal/PassProne 2
	SuccesfulEvasionChance 1
	Popularity 0

Fordward
	Goal/PassProne 7
	Popularity 6
	SuccesfulEvasionChance 5
	CriticalMove 4
	Min.Played 3
	RecoveryTime 2
	SuccesfulBlockChance 1
	GoalKeeping 0
*/

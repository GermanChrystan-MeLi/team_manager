package constants

//===============================//
type BasePosition int

const (
	GoalKeeper BasePosition = iota
	Defenders
	Midfielders
	Forwards
)

//===============================//
type PhysicalState int

const (
	BadlyInjured PhysicalState = iota
	ModeratelyInjured
	MildlyInjured
	GoodShape
	PerfectShape
)

//===============================//
type ImpairmentType int

const (
	Injury ImpairmentType = iota
	Suspension
)

//===============================//
type ImpairmentGravity int

const (
	VeryBad ImpairmentType = iota
	Bad
	Moderate
)

//===============================//
type Footedness int

const (
	Left Footedness = iota
	Right
)

//===============================//
type Talent int

const (
	Low Talent = iota
	Common
	Good
	Great
	Excellent
	Messi
)

//================================//
type Country int

const (
	Argentina Country = iota
	Brazil
	Chile
	Colombia
	Ecuador
	Mexico
	Peru
	Venezuela
)

//================================//
var MinHeight float32 = 1.5
var MaxHeight float32 = 2

//================================//
var MinAge int = 16
var MaxAge int = 42

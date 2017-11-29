package sleepstate

//go:generate stringer -type=SleepState
type SleepState int

const (
	Awake      SleepState = 0
	LightSleep            = 1
	DeepSleep             = 2
	REM                   = 3
)

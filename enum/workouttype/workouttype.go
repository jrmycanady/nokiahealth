package workouttype

//go:generate stringer -type=WorkoutType
type WorkoutType int

// WorkoutType constants for the nokia health api.
const (
	Walk         WorkoutType = 1
	Run          WorkoutType = 2
	Hiking       WorkoutType = 3
	Staking      WorkoutType = 4
	BMX          WorkoutType = 5
	Bicycling    WorkoutType = 6
	Swim         WorkoutType = 7
	Surfing      WorkoutType = 8
	KiteSurfing  WorkoutType = 9
	WindSurfing  WorkoutType = 10
	Bodyboard    WorkoutType = 11
	Tennis       WorkoutType = 12
	TableTennis  WorkoutType = 13
	Squash       WorkoutType = 14
	Badminton    WorkoutType = 15
	LiftWeights  WorkoutType = 16
	Calisthenics WorkoutType = 17
	Elliptical   WorkoutType = 18
	Pilate       WorkoutType = 19
	Basketball   WorkoutType = 20
	Soccer       WorkoutType = 21
	Football     WorkoutType = 22
	Rugby        WorkoutType = 23
	Vollyball    WorkoutType = 24
	WaterPolo    WorkoutType = 25
	HorseRiding  WorkoutType = 26
	Golf         WorkoutType = 27
	Yoga         WorkoutType = 28
	Dancing      WorkoutType = 29
	Boxing       WorkoutType = 30
	Fencing      WorkoutType = 31
	Wrestling    WorkoutType = 32
	MartialArts  WorkoutType = 33
	Skiing       WorkoutType = 34
	SnowBoarding WorkoutType = 35
	Base         WorkoutType = 186
	Rowing       WorkoutType = 187
	Zumba        WorkoutType = 188
	Baseball     WorkoutType = 191
	Handball     WorkoutType = 192
	Hockey       WorkoutType = 194
	Climbing     WorkoutType = 195
	IceSkating   WorkoutType = 196
)

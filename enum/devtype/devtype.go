package devtype

//go:generate stringer -type=DevType
type DevType int

// DevType constants for the nokia health api.
const (
	UserRelated          DevType = 0
	BodyScale                    = 1
	BloodPressureMonitor         = 4
	ActivityTracker              = 16
	SleepMonitor                 = 32
)

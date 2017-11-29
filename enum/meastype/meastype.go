package meastype

//go:generate stringer -type=MeasType
type MeasType int

// MeasType constants for the nokia health api.
const (
	Weight                     MeasType = 1
	Height                              = 4
	FatFreeMassKg                       = 5
	FatRatio                            = 6
	FatMassWeightKg                     = 8
	DiastolicBloodPressureMMHG          = 9
	SystolicBloodPressureMMHG           = 10
	HeartPulseBPM                       = 11
	Temperature                         = 12
	SP02Percent                         = 54
	BodyTemperature                     = 71
	SkinTemperature                     = 73
	MuscleMass                          = 76
	Hydration                           = 77
	BoneMass                            = 88
	PulseWaveVelocity                   = 91
)

package nokiahealth

import (
	"math"
	"net/url"
	"reflect"
	"time"

	"github.com/jrmycanady/nokiahealth/enum/meastype"
	"github.com/jrmycanady/nokiahealth/enum/sleepstate"

	"github.com/jrmycanady/nokiahealth/enum/devtype"
	"github.com/jrmycanady/nokiahealth/enum/status"

	"github.com/jrmycanady/nokiahealth/enum/workouttype"
)

// GetFieldName returns the json filed name for the field if one is found. If
// one is not found it will return an empty string.
func GetFieldName(s interface{}, name string) string {
	t := reflect.TypeOf(s)
	f, ok := t.FieldByName(name)

	var tag string
	if ok {
		tag = f.Tag.Get("json")
	}
	return tag
}

// RevokeNotificationParam provides the query parameters nessasry to revoke a notification.
type RevokeNotificationParam struct {
	CallbackURL url.URL `json:"callbackurl"`
	Appli       *int    `json:"appli"`
}

// RevokeNotificationResp is the response from trying to revoke a notification.
type RevokeNotificationResp struct {
	Status      status.Status `json:"status"`
	RawResponse []byte
	Path        string
}

// NotificationInfoParam provides the query parameters nessasary to retrieve
// information about a specific notification.
type NotificationInfoParam struct {
	CallbackURL url.URL `json:"callbackurl"`
	Appli       *int    `json:"appli"`
}

// NotificationInfoResp represents the unmarshelled api reponse for viewing
// a single notification.
type NotificationInfoResp struct {
	Status      status.Status             `json:"status"`
	Body        *NotificationInfoRespBody `json:"body"`
	RawResponse []byte
	Path        string
}

// NotificationInfoRespBody represents the body of the notification response.
type NotificationInfoRespBody struct {
	Expires       int64      `json:"expires"`
	Comment       string     `json:"comment"`
	ExpiresParsed *time.Time `json:"expiresparsed"`
}

// ListNotificationsParam provides the query parameters nessasary to list
// all the notifications configured for the user.
type ListNotificationsParam struct {
	Appli *int `json:"appli"`
}

// ListNotificationsResp represents the unmarshelled api response for listing notifications.
type ListNotificationsResp struct {
	Status      status.Status              `json:"status"`
	Body        *ListNotificationsRespBody `json:"body"`
	RawResponse []byte
	Path        string
}

// ListNotificationsRespBody represents the notification list body.
type ListNotificationsRespBody struct {
	Profiles []NotificationProfile `json:"profiles"`
}

// NotificationProfile is a notification profile for the user.
type NotificationProfile struct {
	Expires       int64      `json:"expires"`
	Comment       string     `json:"comment"`
	ExpiresParsed *time.Time `json:"expiresparsed"`
}

// CreateNotificationParam provides the query parameters nessasary to create a notication
// via the Nokia Health API.
type CreateNotificationParam struct {
	CallbackURL url.URL `json:"callbackurl"`
	Comment     string  `json:"comment"`
	Appli       int     `json:"appli"`
}

// CreateNotificationResp provides the response of the create request.
type CreateNotificationResp struct {
	Status      status.Status `json:"status"`
	Error       string        `json:"error"`
	RawResponse []byte
	Path        string
}

// SleepSummaryQueryParam provides the query parameters for requests of sleep
// summary data. A date must be specified either with the StartDateYMD/EndDateYMD pair or
// setting the LastUpdate.
// The LastUpdate allow setting to zero for the first call so a time.Time struct
// is not accepted but rather the raw UNIX time.
type SleepSummaryQueryParam struct {
	StartDateYMD *time.Time `json:"startdateymd"`
	EndDateYMD   *time.Time `json:"enddateymd"`
	LastUpdate   *int64     `json:"lastupdate"`
	Offset       *int       `json:"offset"`
}

// SleepMeasuresQueryParam acts as the config parameter for sleep measures requests.
type SleepMeasuresQueryParam struct {
	UserID    int       `json:"userid"`
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
}

// SleepSummaryResp represents the unmarshelled api response for sleep summary.
type SleepSummaryResp struct {
	Status      status.Status     `json:"status"`
	Body        *SleepSummaryBody `json:"body"`
	RawResponse []byte
	Path        string
}

// SleepSummaryBody represents the unmarshelled api response for the sleep summary body.
type SleepSummaryBody struct {
	Series []SleepSummary `json:"series"`
	More   bool           `json:"more"`
}

// SleepSummary is a summary of one sleep entry.
type SleepSummary struct {
	ID              int64            `json:"id"`
	StartDate       int64            `json:"startdate"`
	EndDate         int64            `json:"enddate"`
	Date            string           `json:"date"`
	TimeZone        string           `json:"timezone"`
	Model           int              `json:"model"`
	Data            SleepSummaryData `json:"data"`
	Modified        int64            `json:"modified"`
	StartDateParsed *time.Time       `json:"startdateparsed"`
	EndDateParsed   *time.Time       `json:"enddateparsed"`
	DateParsed      *time.Time       `json:"dateparsed"`
}

// SleepSummaryData contains the summary data for the sleep summary. Not all fields are required
// so some are pointers and can be nil.
type SleepSummaryData struct {
	WakeUpDuration     int  `json:"wakeupduration"`
	LightSleepDuration int  `json:"lightsleepduration"`
	DeepSleepDuration  int  `json:"deepsleepduration"`
	REMSleepDuration   *int `json:"remsleepduration"`
	WakeUpCount        int  `json:"wakeupcount"`
	DurationToSleep    int  `json:"durationtosleep"`
	DurationToWakeUp   *int `json:"durationtowakeup"`
}

// SleepMeasuresResp represents the unmarshelled api response for sleep measures.
type SleepMeasuresResp struct {
	Status      status.Status          `json:"status"`
	Body        *SleepMeasuresRespBody `json:"body"`
	RawResponse []byte
	Path        string
}

// SleepMeasuresRespBody actrepresents the unmarshelled api response for sleep measures body.
type SleepMeasuresRespBody struct {
	Series []SleepMeasure `json:"series"`
	Model  int            `json:"model"`
}

// SleepMeasure is a specific instance of sleep returned by the API.
type SleepMeasure struct {
	StartDate       int64                 `json:"startdate"`
	EndDate         int64                 `json:"enddate"`
	State           sleepstate.SleepState `json:"state"`
	StartDateParsed *time.Time            `json:"startdateparsed"`
	EndDateParsed   *time.Time            `'json:"enddateparsed"`
}

// IntradayActivityQueryParam acts as the config parameter for intraday activity retrieval requests.
type IntradayActivityQueryParam struct {
	UserID    int        `json:"userid"`
	StartDate *time.Time `json:"startdate"`
	EndDate   *time.Time `json:"enddate"`
}

// IntradayActivityResp represents the unmarshelled api response for intraday activities.
type IntradayActivityResp struct {
	Status      status.Status             `json:"status"`
	Error       string                    `json:"error"`
	Body        *IntradayActivityRespBody `json:"body"`
	RawResponse []byte
	Path        string
}

// IntradayActivityRespBody represents the unmarshelled api response body for intraday activities.
type IntradayActivityRespBody struct {
	Series map[int64]IntraDayActivity `json:"series"`
}

// IntraDayActivity represents an intra day activity as returned by the API.
// Their is likey work to be done here as the documentation does not provide
// musch information reegarding what paramters it should contain.
type IntraDayActivity struct {
	Calories  *float64 `json:"calories"`
	Distance  *float64 `json:"distance"`
	Duration  *int     `json:"duration"`
	Elevation *float64 `json:"elevation"`
	Steps     *int     `json:"steps"`
	PoolLap   *int     `json:"pool_lap"`
}

// WorkoutsQueryParam acts as the config parameter for workout retrieval requests.
type WorkoutsQueryParam struct {
	UserID       int        `json:"userid"`
	StartDateYMD *time.Time `json:"startdateymd"`
	EndDateYMD   *time.Time `json:"enddateymd"`
}

// WorkoutResponse represents the unmarshelled api response for workouts.
type WorkoutResponse struct {
	Status      status.Status    `json:"status"`
	Body        *WorkoutRespBody `json:"body"`
	RawResponse []byte
	Path        string
}

// WorkoutRespBody represents the unmarshelled body of the workout api resposne.
type WorkoutRespBody struct {
	Series []Workout `json:"series"`
}

// Workout contains each workout entry as returned by the API. The raw dates are provided
// but fully parsed timeTime structs can be accessed via the same name as the field
// but with Parsed added. i.e. StartDate => StartDateParsed
type Workout struct {
	ID              int                      `json:"id"`
	UserID          int                      `json:"userid"`
	Category        *workouttype.WorkoutType `json:"category"`
	StartDate       int64                    `json:"startdate"`
	EndDate         int64                    `json:"enddate"`
	Model           int                      `json:"model"`
	Attrib          int                      `json:"attrib"`
	Date            string                   `json:"date"`
	TimeZone        string                   `json:"timezone"`
	Modified        int                      `json:"modified"`
	Data            map[string]float64       `json:"data"`
	StartDateParsed *time.Time               `json:"startdateparsed"`
	EndDateParsed   *time.Time               `json:"enddateparsed"`
	DateParsed      *time.Time               `json:"dateparsed"`
}

// ActivityMeasuresQueryParam acts as the config parameter for activity measurement queries.
// All options feilds can be set to null but at least one of the date fields need to be
// specified or the API will fail. Additionally there is no ParseResponse option as
// there is no need to because the activities response doesn't need further parsing.
type ActivityMeasuresQueryParam struct {
	UserID int `json:"userid"`
	// Date             *time.Time `json:"date"`
	StartDateYMD     *time.Time `json:"startdateymd"`
	EndDateYMD       *time.Time `json:"enddateymd"`
	LasteUpdate      *time.Time `json:"lastupdate"`
	Offset           *int       `json:"offset"`
	DisableDateParse bool       `json:"diabledateparse"`
}

// ActivitiesMeasuresResp contains the unmarshalled response from the api.
// If the client has been set to include raw respeonse the RawResponse byte slice
// will be populated with raw bytes returned by the API.
type ActivitiesMeasuresResp struct {
	Status      status.Status               `json:"status"`
	Error       string                      `json:"error"`
	Body        *ActivitiesMeasuresRespBody `json:"body"`
	RawResponse []byte
	Path        string
}

// ActivitiesMeasuresRespBody contains the response body as provided by the
// api. The Nokia Health API includes single values responses directly in the
// body. As such they are all pointers. You may check SingleValue to determine
// if a single value was provided.
type ActivitiesMeasuresRespBody struct {
	ParsedDate  *time.Time `json:"parseddate"`
	Date        *string    `json:"date"`
	Steps       *float64   `json:"steps"`
	Distance    *float64   `json:"distance"`
	Calories    *float64   `json:"calories"`
	Elevation   *float64   `json:"elevation"`
	Soft        *int       `json:"soft"`
	Moderate    *int       `json:"moderate"`
	Intense     *int       `json:"intense"`
	TimeZone    *string    `json:"timezone"`
	Activities  []Activity `json:"activity"`
	More        bool       `json:"more"`
	Offset      int        `json:"offset"`
	SingleValue bool       `json:"singleValue"`
}

// Activity represents an activity as recorded by Nokia Health.
type Activity struct {
	ParsedDate *time.Time `json:"parseddate"`
	Date       string     `json:"date"`
	Steps      float64    `json:"steps"`
	Distance   float64    `json:"distance"`
	Calories   float64    `json:"calories"`
	Elevation  float64    `json:"elevation"`
	Soft       int        `json:"soft"`
	Moderate   int        `json:"moderate"`
	Intense    int        `json:"intense"`
	TimeZone   string     `json:"timezone"`
}

// BodyMeasuresQueryParams acts as the config parameter for body measurement queries.
// All optional field can be set to null.
// The ParsedResponse can be set to true and the request will automatically parse
// the response into easy to use structs. Otherwise this can be done manually when
// needed via the Parse method.
type BodyMeasuresQueryParams struct {
	UserID        int                `json:"userid"`
	StartDate     *time.Time         `json:"startdate"`
	EndDate       *time.Time         `json:"enddate"`
	LastUpdate    *time.Time         `json:"lastupdate"`
	DevType       *devtype.DevType   `json:"devtype"`
	MeasType      *meastype.MeasType `json:"meastype"`
	Category      *int               `json:"category"`
	Limit         *int               `json:"limit"`
	Offset        *int               `json:"offset"`
	ParseResponse bool
}

// BodyMeasuresResp contains the unmarshalled response from the api.
// If the client has been set to include raw respeonse the RawResponse byte slice
// will be populated with raw bytes returned by the API.
type BodyMeasuresResp struct {
	Status         status.Status        `json:"status"`
	Body           *BodyMeasureRespBody `json:"body"`
	RawResponse    []byte
	Path           string
	ParsedResponse *BodyMeasures
}

// BodyMeasureRespBody represents the body portion of the body measure response.
// The body portion is not required and thus this may not be found in the response
// object.
type BodyMeasureRespBody struct {
	Updatetime  int64                  `json:"updatetime"`
	More        int                    `json:"more"`
	Timezone    string                 `json:"timezone"`
	MeasureGrps []BodyMeasureGroupResp `json:"measuregrps"`
}

// BodyMeasureGroupResp is a single body measurment group as found in the resposne.
// Each group has a set of measures that can then be parsed manually or via the
// Parse method on BodyMeasuresQueryParams.
type BodyMeasureGroupResp struct {
	GrpID    int                   `json:"grpid"`
	Attrib   int                   `json:"attrib"`
	Date     int64                 `json:"date"`
	Category int                   `json:"category"`
	Measures []BodyMeasuresMeasure `json:"measures"`
}

// BodyMeasuresMeasure is a single body measure found in the response.
type BodyMeasuresMeasure struct {
	Value int               `json:"value"`
	Type  meastype.MeasType `json:"type"`
	Unit  int               `json:"unit"`
}

type Weight struct {
	Date     time.Time
	Kgs      float64
	Attrib   int
	Category int
}

type Height struct {
	Date     time.Time
	Meters   float64
	Attrib   int
	Category int
}

type FatFreeMass struct {
	Date     time.Time
	Kgs      float64
	Attrib   int
	Category int
}

type FatMassWeight struct {
	Date     time.Time
	Kgs      float64
	Attrib   int
	Category int
}

type FatRatio struct {
	Date     time.Time
	Ratio    float64
	Attrib   int
	Category int
}

type DiastolicBloodPressure struct {
	Date     time.Time
	MmHg     float64
	Attrib   int
	Category int
}

type SystolicBloodPressure struct {
	Date     time.Time
	MmHg     float64
	Attrib   int
	Category int
}

type HeartPulse struct {
	Date     time.Time
	BPM      float64
	Attrib   int
	Category int
}

type Temperature struct {
	Date     time.Time
	Celcius  float64
	Attrib   int
	Category int
}

type SP02Percent struct {
	Date       time.Time
	Percentage float64
	Attrib     int
	Category   int
}

type BodyTemperature struct {
	Date     time.Time
	Celcius  float64
	Attrib   int
	Category int
}

type SkinTemperature struct {
	Date     time.Time
	Celcius  float64
	Attrib   int
	Category int
}

type MuscleMass struct {
	Date     time.Time
	Mass     float64
	Attrib   int
	Category int
}

type Hydration struct {
	Date      time.Time
	Hydration float64
	Attrib    int
	Category  int
}

type BoneMass struct {
	Date     time.Time
	Mass     float64
	Attrib   int
	Category int
}

type PulseWaveVelocity struct {
	Date     time.Time
	Velocity float64
	Attrib   int
	Category int
}

type BodyMeasures struct {
	Weights                 []Weight
	Heights                 []Height
	FatFreeMass             []FatFreeMass
	FatRatios               []FatRatio
	FatMassWeights          []FatMassWeight
	DiastolicBloodPressures []DiastolicBloodPressure
	SystolicBloodPressures  []SystolicBloodPressure
	HeartPulses             []HeartPulse
	Temperatures            []Temperature
	SP02Percents            []SP02Percent
	BodyTemperatures        []BodyTemperature
	SkinTemperatures        []SkinTemperature
	MuscleMasses            []MuscleMass
	Hydration               []Hydration
	BoneMasses              []BoneMass
	PulseWaveVelocity       []PulseWaveVelocity
}

// ParseData parses all the data provided into buckets of each type of
// measurement. It also performs the nessasary date and unit conversion.
func (rm BodyMeasuresResp) ParseData() *BodyMeasures {
	bm := BodyMeasures{}

	if rm.Body != nil {
		// process all measurements
		for mgID, _ := range rm.Body.MeasureGrps {
			// build the time
			d := time.Unix(int64(rm.Body.MeasureGrps[mgID].Date), 0)

			for mID, _ := range rm.Body.MeasureGrps[mgID].Measures {
				switch {
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.Weight:
					w := Weight{
						Date:     d,
						Kgs:      convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.Weights = append(bm.Weights, w)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.Height:
					h := Height{
						Date:     d,
						Meters:   convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.Heights = append(bm.Heights, h)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.FatFreeMassKg:
					ffm := FatFreeMass{
						Date:     d,
						Kgs:      convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.FatFreeMass = append(bm.FatFreeMass, ffm)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.FatRatio:
					fr := FatRatio{
						Date:     d,
						Ratio:    convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.FatRatios = append(bm.FatRatios, fr)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.FatMassWeightKg:
					fmw := FatMassWeight{
						Date:     d,
						Kgs:      convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.FatMassWeights = append(bm.FatMassWeights, fmw)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.DiastolicBloodPressureMMHG:
					dbp := DiastolicBloodPressure{
						Date:     d,
						MmHg:     convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.DiastolicBloodPressures = append(bm.DiastolicBloodPressures, dbp)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.SystolicBloodPressureMMHG:
					sbp := SystolicBloodPressure{
						Date:     d,
						MmHg:     convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.SystolicBloodPressures = append(bm.SystolicBloodPressures, sbp)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.HeartPulseBPM:
					hp := HeartPulse{
						Date:     d,
						BPM:      convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.HeartPulses = append(bm.HeartPulses, hp)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.Temperature:
					t := Temperature{
						Date:     d,
						Celcius:  convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.Temperatures = append(bm.Temperatures, t)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.SP02Percent:
					p := SP02Percent{
						Date:       d,
						Percentage: convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:     rm.Body.MeasureGrps[mgID].Attrib,
						Category:   rm.Body.MeasureGrps[mgID].Category,
					}
					bm.SP02Percents = append(bm.SP02Percents, p)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.BodyTemperature:
					t := BodyTemperature{
						Date:     d,
						Celcius:  convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.BodyTemperatures = append(bm.BodyTemperatures, t)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.SkinTemperature:
					t := SkinTemperature{
						Date:     d,
						Celcius:  convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.SkinTemperatures = append(bm.SkinTemperatures, t)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.MuscleMass:
					m := MuscleMass{
						Date:     d,
						Mass:     convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.MuscleMasses = append(bm.MuscleMasses, m)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.Hydration:
					h := Hydration{
						Date:      d,
						Hydration: convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:    rm.Body.MeasureGrps[mgID].Attrib,
						Category:  rm.Body.MeasureGrps[mgID].Category,
					}
					bm.Hydration = append(bm.Hydration, h)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.BoneMass:
					m := BoneMass{
						Date:     d,
						Mass:     convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.BoneMasses = append(bm.BoneMasses, m)
				case rm.Body.MeasureGrps[mgID].Measures[mID].Type == meastype.PulseWaveVelocity:
					v := PulseWaveVelocity{
						Date:     d,
						Velocity: convertUnits(rm.Body.MeasureGrps[mgID].Measures[mID].Value, rm.Body.MeasureGrps[mgID].Measures[mID].Unit),
						Attrib:   rm.Body.MeasureGrps[mgID].Attrib,
						Category: rm.Body.MeasureGrps[mgID].Category,
					}
					bm.PulseWaveVelocity = append(bm.PulseWaveVelocity, v)
				}
			}
		}
	}

	return &bm
}

// convertUnits converts the value to the units specified.
func convertUnits(value int, unit int) float64 {
	return float64(value) * math.Pow(10, float64(unit))
}

package nokiahealth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/jrmycanady/oauth1"
)

const (
	getBodyMeasuresURL = "https://api.health.nokia.com/measure?action=getmeas"
)

// NokiaHealthEndpoint provides the oauth endpoint URLs for the Nokia Health API.
// These have been exposed so they may be overridden if the access URLs have
// been changed.
var NokiaHealthEndpoint = oauth1.Endpoint{
	RequestTokenURL: "https://developer.health.nokia.com/account/request_token",
	AuthorizeURL:    "https://developer.health.nokia.com/account/authorize",
	AccessTokenURL:  "https://developer.health.nokia.com/account/access_token",
}

// Client contains all the required information to interact with the nokia API.
type Client struct {
	OAuthConfig     oauth1.Config
	SaveRawResponse bool
}

// NewClient creates a new client using the consumer information provided. The
// required parameters can be obtained when developers register with Nokia
// to use the API.
func NewClient(consumerKey string, consumerSecret string, callbackURL string) Client {

	return Client{
		OAuthConfig: oauth1.Config{
			ConsumerKey:            consumerKey,
			ConsumerSecret:         consumerSecret,
			CallbackURL:            callbackURL,
			Endpoint:               NokiaHealthEndpoint,
			DisableCallbackConfirm: true,
			IncludeQueryParams:     true,
		},
	}
}

// AccessRequest represents a request for access to a user. Once created it
// can generate the authorization URL and then generate a user based on
// the verifier and user ID provided.
type AccessRequest struct {
	RequestToken     string
	RequestSecret    string
	AuthorizationURL *url.URL
	Client           Client
}

// CreateAccessRequest creates a new access request based on the clients credentials.
// The returned AccessRequest will contain the authorization URL
// needed for users to authorize access.
func (c Client) CreateAccessRequest() (AccessRequest, error) {

	ar := AccessRequest{
		Client: c,
	}

	var err error
	ar.RequestToken, ar.RequestSecret, err = c.OAuthConfig.RequestToken()
	if err != nil {
		return ar, err
	}

	ar.AuthorizationURL, err = c.OAuthConfig.AuthorizationURL(ar.RequestToken)
	if err != nil {
		return ar, err
	}

	return ar, nil
}

// User is a Nokia Health user account that can be interacted with via the
// api.
type User struct {
	UserID          int
	AccessTokenStr  string
	AccessSecretStr string
	Client          Client
	AccessToken     *oauth1.Token
}

// GenerateUser creates a validated user object via the access request return
// verifier string. Note, the user ID is needed as it is required for all
// further queries to the api.
func (ar AccessRequest) GenerateUser(verifier string, userID int) (User, error) {
	u := User{
		Client: ar.Client,
		UserID: userID,
	}

	var err error
	u.AccessTokenStr, u.AccessSecretStr, err = ar.Client.OAuthConfig.AccessToken(ar.RequestToken, ar.RequestSecret, verifier)
	if err != nil {
		return u, err
	}

	u.AccessToken = oauth1.NewToken(u.AccessTokenStr, u.AccessSecretStr)

	return u, nil
}

// GetActivityMeasures retrieves the activity measurements as specified by the config
// provided.
func (u User) GetActivityMeasures(params *ActivityMeasureQueryParam) (RawActivitiesMeasuresResponse, error) {
	activityMeasureResponse := RawActivitiesMeasuresResponse{}

	httpClient := u.Client.OAuthConfig.Client(oauth1.NoContext, u.AccessToken)

	// Build query params
	v := url.Values{}
	v.Add("userid", strconv.Itoa(u.UserID))
	v.Add("action", "getactivity")

	if params != nil {
		if params.Date != nil {
			v.Add(GetFieldName(*params, "Date"), params.Date.Format("2006-01-02"))
		}
		if params.StartDateYMD != nil {
			v.Add(GetFieldName(*params, "StartDateYMD"), params.Date.Format("2006-01-02"))
		}
		if params.EndDateYMD != nil {
			v.Add(GetFieldName(*params, "EndDateYMD"), params.Date.Format("2006-01-02"))
		}
		if params.LasteUpdate != nil {
			v.Add(GetFieldName(*params, "LasteUpdate"), strconv.FormatInt(params.LasteUpdate.Unix(), 10))
		}
	}

	path := fmt.Sprintf("http://api.health.nokia.com/v2/measure?%s", v.Encode())

	log.Println(path)
	resp, err := httpClient.Get(path)
	if err != nil {
		return activityMeasureResponse, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return activityMeasureResponse, err
	}
	if u.Client.SaveRawResponse {
		activityMeasureResponse.RawResponse = body
	}

	err = json.Unmarshal(body, &activityMeasureResponse)
	if err != nil {
		return activityMeasureResponse, err
	}

	// Parse date time if possible.
	if activityMeasureResponse.Body.Date != nil && activityMeasureResponse.Body.TimeZone != nil {
		location, err := time.LoadLocation(*activityMeasureResponse.Body.TimeZone)
		if err != nil {
			return activityMeasureResponse, err
		}

		t, err := time.Parse("2006-01-02", *activityMeasureResponse.Body.Date)
		if err != nil {
			return activityMeasureResponse, err
		}

		t = t.In(location)
		activityMeasureResponse.Body.ParsedDate = &t

		activityMeasureResponse.Body.SingleValue = true
	}

	for aID, _ := range activityMeasureResponse.Body.Activities {
		location, err := time.LoadLocation(activityMeasureResponse.Body.Activities[aID].TimeZone)
		if err != nil {
			return activityMeasureResponse, err
		}

		t, err := time.Parse("2006-01-02", activityMeasureResponse.Body.Activities[aID].Date)
		if err != nil {
			return activityMeasureResponse, err
		}

		t = t.In(location)
		activityMeasureResponse.Body.Activities[aID].ParsedDate = &t
	}

	return activityMeasureResponse, nil
}

func (u User) GetWorkouts() {
	httpClient := u.Client.OAuthConfig.Client(oauth1.NoContext, u.AccessToken)

	v := url.Values{}
	v.Add("userid", strconv.Itoa(u.UserID))
	v.Add("action", "getworkouts")

	path := fmt.Sprintf("http://api.health.nokia.com/v2/measure?%s", v.Encode())

	resp, err := httpClient.Get(path)
	if err != nil {
		// return bodyMeasureResponse, err
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		// return bodyMeasureResponse, err
	}

	log.Println(string(body))

}

// GetBodyMeasure retrieves the measurements as specified by the config
// provided.
func (u User) GetBodyMeasure(params *BodyMeasuresQueryParams) (RawBodyMeasuresResponse, error) {
	bodyMeasureResponse := RawBodyMeasuresResponse{}

	httpClient := u.Client.OAuthConfig.Client(oauth1.NoContext, u.AccessToken)

	// Build query params
	v := url.Values{}
	v.Add("userid", strconv.Itoa(u.UserID))
	v.Add("action", "getmeas")

	if params != nil {
		if params.StartDate != nil {
			v.Add(GetFieldName(*params, "StartDate"), strconv.FormatInt(params.StartDate.Unix(), 10))
		}
		if params.EndDate != nil {
			v.Add(GetFieldName(*params, "EndDate"), strconv.FormatInt(params.EndDate.Unix(), 10))
		}
		if params.LastUpdate != nil {
			v.Add(GetFieldName(*params, "LastUpdate"), strconv.FormatInt(params.EndDate.Unix(), 10))
		}
		if params.DevType != nil {
			v.Add(GetFieldName(*params, "DevType"), strconv.Itoa(int(*params.DevType)))
		}
		if params.MeasType != nil {
			v.Add(GetFieldName(*params, "MeasType"), strconv.Itoa(int(*params.MeasType)))
		}
		if params.Category != nil {
			v.Add(GetFieldName(*params, "Category"), strconv.Itoa(*params.Category))
		}
		if params.Limit != nil {
			v.Add(GetFieldName(*params, "Limit"), strconv.Itoa(*params.Limit))
		}
		if params.Offset != nil {
			v.Add(GetFieldName(*params, "Offset"), strconv.Itoa(*params.Offset))
		}
	}

	path := fmt.Sprintf("http://api.health.nokia.com/measure?%s", v.Encode())

	resp, err := httpClient.Get(path)
	if err != nil {
		return bodyMeasureResponse, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bodyMeasureResponse, err
	}
	if u.Client.SaveRawResponse {
		bodyMeasureResponse.RawResponse = body
	}

	err = json.Unmarshal(body, &bodyMeasureResponse)
	if err != nil {
		return bodyMeasureResponse, err
	}

	if params.ParseResponse {
		bodyMeasureResponse.ParsedResponse = bodyMeasureResponse.ParseData()
	}

	return bodyMeasureResponse, nil

}

// GenerateUser creates a new user object based on the values provided for the
// user.
//
// All values must be provided for requests to process successfully. Use the
// AccessRequest GenerateUser to obtain all of the values for later user with
// this method.
func (c Client) GenerateUser(accessToken string, accessSecret string, userID int) User {
	u := User{
		Client:          c,
		UserID:          userID,
		AccessTokenStr:  accessToken,
		AccessSecretStr: accessSecret,
	}

	u.AccessToken = oauth1.NewToken(u.AccessTokenStr, u.AccessSecretStr)

	return u
}

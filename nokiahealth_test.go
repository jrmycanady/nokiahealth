package nokiahealth

import (
	"io/ioutil"
	"net/url"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
)

type TestConfig struct {
	DevToken   string
	DevSecret  string
	UserToken  string
	UserSecret string
	UserID     int
}

func LoadConfig() error {
	rawConfigData, err := ioutil.ReadFile("./test.toml")
	if err != nil {
		return err
	}

	_, err = toml.Decode(string(rawConfigData), &tc)
	if err != nil {
		return err
	}

	return nil
}

var tc TestConfig

func TestGetBodyMeasures(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Fatal("Failed to load config file.")
	}

	// Build the client.
	c := NewClient(tc.DevToken, tc.DevSecret, "")
	c.SaveRawResponse = true

	// Build the user
	u := c.GenerateUser(tc.UserToken, tc.UserSecret, tc.UserID)

	startDate := time.Now().AddDate(0, 0, -2)
	endDate := time.Now()

	p := BodyMeasuresQueryParams{
		StartDate: &startDate,
		EndDate:   &endDate,
	}

	m, err := u.GetBodyMeasures(&p)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if m.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", m.Status, m.Status.String())
	}
}

func TestGetActivityMeasures(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Fatal("Failed to load config file.")
	}

	// Build the client.
	c := NewClient(tc.DevToken, tc.DevSecret, "")
	c.SaveRawResponse = true

	// Build the user
	u := c.GenerateUser(tc.UserToken, tc.UserSecret, tc.UserID)

	m, err := u.GetActivityMeasures(nil)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if m.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", m.Status, m.Status.String())
	}
}

func TestGetIntradayActivity(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Fatal("Failed to load config file.")
	}

	// Build the client.
	c := NewClient(tc.DevToken, tc.DevSecret, "")
	c.SaveRawResponse = true

	// Build the user
	u := c.GenerateUser(tc.UserToken, tc.UserSecret, tc.UserID)

	m, err := u.GetIntradayActivity(nil)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if m.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", m.Status, m.Status.String())
	}
}

func TestGetWorkouts(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Fatal("Failed to load config file.")
	}

	// Build the client.
	c := NewClient(tc.DevToken, tc.DevSecret, "")
	c.SaveRawResponse = true

	// Build the user
	u := c.GenerateUser(tc.UserToken, tc.UserSecret, tc.UserID)

	m, err := u.GetWorkouts(nil)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if m.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", m.Status, m.Status.String())
	}
}

func TestGetSleepMeasures(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Fatal("Failed to load config file.")
	}

	// Build the client.
	c := NewClient(tc.DevToken, tc.DevSecret, "")
	c.SaveRawResponse = true

	// Build the user
	u := c.GenerateUser(tc.UserToken, tc.UserSecret, tc.UserID)

	m, err := u.GetSleepMeasures(nil)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if m.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", m.Status, m.Status.String())
	}
}

func TestGetSleepSummary(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Fatal("Failed to load config file.")
	}

	// Build the client.
	c := NewClient(tc.DevToken, tc.DevSecret, "")
	c.SaveRawResponse = true

	// Build the user
	u := c.GenerateUser(tc.UserToken, tc.UserSecret, tc.UserID)

	m, err := u.GetSleepSummary(nil)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if m.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", m.Status, m.Status.String())
	}
}

func TestNotificationFunctions(t *testing.T) {

	err := LoadConfig()
	if err != nil {
		t.Fatal("Failed to load config file.")
	}

	// Build the client.
	c := NewClient(tc.DevToken, tc.DevSecret, "")
	c.SaveRawResponse = true

	// Build the user
	u := c.GenerateUser(tc.UserToken, tc.UserSecret, tc.UserID)

	// Test creating a notification
	var ul *url.URL
	ul, err = url.Parse("http://example.com")
	p := CreateNotificationParam{
		CallbackURL: *ul,
		Comment:     "this is a test",
		Appli:       1,
	}
	n, err := u.CreateNotification(&p)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if n.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", n.Status, n.Status.String())
	}

	// Test finding the notification.
	appli := 1
	p2 := NotificationInfoParam{
		CallbackURL: *ul,
		Appli:       &appli,
	}

	gn, err := u.GetNotificationInformation(&p2)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if gn.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", gn.Status, gn.Status.String())
	}

	// Test revoking the notification
	p3 := RevokeNotificationParam{
		CallbackURL: *ul,
		Appli:       &appli,
	}
	rn, err := u.RevokeNotification(&p3)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if rn.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", rn.Status, rn.Status.String())
	}
}

func TestListNotifications(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Fatal("Failed to load config file.")
	}

	// Build the client.
	c := NewClient(tc.DevToken, tc.DevSecret, "")
	c.SaveRawResponse = true

	// Build the user
	u := c.GenerateUser(tc.UserToken, tc.UserSecret, tc.UserID)

	m, err := u.ListNotifications(nil)
	if err != nil {
		t.Fatalf("failed to get body measurements: %v", err)
	}

	if m.Status != 0 {
		t.Fatalf("failed to get body measurements with api error %d => %v", m.Status, m.Status.String())
	}
}

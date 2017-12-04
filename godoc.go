// The package nokiahealth is a client library for for working with the Nokia
// Health (Withings) API. It includes support for all resources listed on the public
// api documentation at https://developer.health.nokia.com/api/doc. This
// includes everything from user access token generation to retreiving
// body measurments and setting up notifications.
//
// Due to limitEd access of Nokia Health devices, testing of each
// resource has been limited. If you have access to more data or run into
// any bugs/missing features, please place an issue on github at
// https://github.com/jrmycanady/nokiahealth. The API documentation has been
// found to be severly lacking or compeltely incorrect so oddities are
// expected to be found. Most are easily fixed as long as the raw request data
// can be provided.
//
// Oauth Notes
//
// The Nokia Health Oauth implementation required modifications to the
// Oauth package. As such this package relies upon a fork of dghubble/oauth1
// located at jrmycanady/oauth1.
//
// Supported Resources
//
//  Generate User Access Request and URL
//
//  Get Activity Measures
//
// * Get Body Measures
// * Get Intraday Activity (Requires additional approval by Nokia Health)
// * Get Sleep Measures
// * Get Sleep Summary
// * Get Workouts
// * Create Notification
// * Get Notification Information
// * List Notifications
// * Revoke Notification
//
// Basic Usage
//
// The basic use pattern for this client is to instantiate the client using
// the developer Oath consumer key, consumer secret, and callback URL. The
// callback URL can be an empty string if there is no need for generating
// user authorization URLs.
//   client := nokiahealth.NewClient("consumer_key","consumer_secret", "callback_url")
// The client can be used to handle user authorization as well as generate
// user structs from stores user authorization tokens. Details of that can be
// found in the authorization section. The following is an examle of creating
// a user from known tokens and secrets. All three parameters are required as
// the API doesn't rely on just the user_tokent to identify the user.
//   user := client.GenerateUser("user_token", "user_secret", "user_id")
// With the user defined API calls can be made by utilizing on of the methods
// for the User struct. i.e.
//   m, err := u.GetBodyMeasures(nil)
// In the above example m will contain the results of the request. All methods
// support an optional param that contains any specifics on the request such
// as start date or end date. Refer to the Nokia Health documentation on when
// and how to use these parameters. https://developer.health.nokia.com/api/doc
//
//
package nokiahealth

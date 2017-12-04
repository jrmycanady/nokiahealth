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
//  Get Activity Measures
//  Get Body Measures
//  Get Intraday Activity (Requires additional approval by Nokia Health)
//  Get Sleep Measures
//  Get Sleep Summary
//  Get Workouts
//  Create Notification
//  Get Notification Information
//  List Notifications
//  Revoke Notification
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
// In the above example, m will contain the results of the request. All methods
// support an optional param that contains any specifics on the request such
// as start date or end date. Refer to the Nokia Health documentation on when
// and how to use these parameters. https://developer.health.nokia.com/api/doc.
// Each param type is specific to the request type.
//
// Each method varies somewhat in the foramt it returns due to the inconsistencies
// in the Nokia Health API. This package does what it can to work around
// these inconsistencies but you should refer to the documentation for each
// method for more information. The one commonality is that the package does
// parse all dates into time.Time structs. The API returns a random assortment
// of time formats. The raw data is always included but for each raw field there
// should be a parsed counterpart.
//
// Authorization Overview
//
// Authorization is like most other Oauth1 based flows. You must first register
// as a developer to obtain your consumer key and secret at
//  https://developer.health.nokia.com/partner/dashboard. This will allow you
// to make requests for authorization to access a users data.
//
// Next you must obtain authorization to access a users data via the API. This
// is done by generating a specially crafted URL that the user must navigate to,
// then authenticate and approve your access. Once approved, the user will either
// be presented with a token, verifier, and user id or be redirected to the callback
// URL specified. In either case, you need to obtain the user token, the verifier,
// and the user's ID. NOTE: You will not yet have all the information needed
// to access the user account. You will need to save the access request
// token an secret until you get the information from the user. The verifier will
// be used to generate the user secret that is needed for requests.
//  // Generate an access request. (Assuming client is already created.)
//  ar, err != client.CreateAccessRequest()
//  if err != nil {
//    log.Fatal(err)
//  }
//  // Somehow present the AuthorizationURL to the user to navigate to.
//  // There is a 2 minute time window before this URL is no longer valid.
//  fmt.Println(ar.AuthorizationURL)
// The callback or user will have access to the verifier, userid, user token,
// and user secret. From this you can generate a user which essentially
// obtains the user secret and builds a normal user struct for the user.
//  u, err != ar.GenerateUser("<verifier", <userid>)
//  if err != nil {
//	  log.Fatal(err)
//  }
// If all has gone well, you will now have a user struct to perform actions
// against.
//
// In some cases, due to the application design, the callback may hit
// a different process or the original access request struct is not available.
// In such a case the generate user cannot be used which means you cannot
// obtain the secret. In those cases you can rebuild an access request token
// using the RebuildAccessReuqest() method of the client. This still requires
// storing the access request token and secret somewhere. Both of which are
// public fields on the AccessRequest struct.
//  ar := client.RebuildAccessReuqest("token", "secret")

package nokiahealth

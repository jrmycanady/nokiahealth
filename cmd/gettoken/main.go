// This provides a simple cli based tool to test the interface.
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jrmycanady/nokiahealth"
)

func main() {

	// Getting client information from user.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Client ID: ")
	clientID, _ := reader.ReadString('\n')
	clientID = strings.TrimSuffix(clientID, "\n")
	clientID = strings.TrimSuffix(clientID, "\r")

	// Get client secret
	fmt.Print("Client Secret: ")
	clientSecret, _ := reader.ReadString('\n')
	clientSecret = strings.TrimSuffix(clientSecret, "\n")
	clientSecret = strings.TrimSuffix(clientSecret, "\r")

	// Get the redirect URL
	fmt.Print("Redirect URL: ")
	clientRedirectURL, _ := reader.ReadString('\n')
	clientRedirectURL = strings.TrimSuffix(clientRedirectURL, "\n")
	clientRedirectURL = strings.TrimSuffix(clientRedirectURL, "\r")

	// Building new nokiahealth client.
	client := nokiahealth.NewClient(clientID, clientSecret, clientRedirectURL)
	client.IncludePath = true
	fmt.Println("--------------- Client ---------------")
	fmt.Printf("Client ID: %s\nClient Secret: %s\nClient Redirect URL: %s\n", clientID, clientSecret, clientRedirectURL)
	fmt.Println("--------------- ------ ---------------")

	// Provide user with authorization URL.
	authURL, _, err := client.AuthCodeURL() // Ignoring state for this simple test.
	if err != nil {
		fmt.Println("failed to generate url: %s\n", err)
		return
	}
	fmt.Println("Navigate to the following URL and copy out the code from the params and provide below.")
	fmt.Printf("URL: %s\n", authURL)

	// Get code from user.
	fmt.Print("Code: ")
	code, _ := reader.ReadString('\n')
	code = strings.TrimSuffix(code, "\n")
	code = strings.TrimSuffix(code, "\r")

	// Get user.
	u, err := client.NewUserFromAuthCode(context.Background(), code)
	if err != nil {
		fmt.Printf("failed to get user: %s\n", err)
		return
	}

	t, err := u.Token.Token()
	if err != nil {
		fmt.Printf("failed to get token: %s\n", err)
	}
	fmt.Printf("Token:%s\nRefresh Token:%s", t.AccessToken, t.RefreshToken)
}

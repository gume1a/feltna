package middleware

import (
	"context"
	"fmt"
	githubapi "github.com/google/go-github/v50/github"
	oauthclient "github.com/gume1a/oauthproxy/pkg/client"
	"github.com/gume1a/oauthproxy/pkg/identity"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"log"
	"net/url"
)

func GithubAuthentication(scopes []string, next func(token *oauth2.Token, client *githubapi.Client)) {
	redirectURL, _ := url.Parse("http://127.0.0.1:1420/oauth2/callback")
	proxyURL, _ := url.Parse("http://localhost:8081")
	authURL, _ := url.Parse(github.Endpoint.AuthURL)

	oauthClient := oauthclient.
		NewFactory(identity.GITHUB, "242f79440a257b6370b8").
		WithScopes(scopes).
		WithRedirectURL(redirectURL).
		WithProxyURL(proxyURL).
		WithAuthURL(authURL).
		Build()

	authUrl, tokenResponseChan := oauthClient.AuthCodeURL(
		"5ca75bd30",
		oauth2.AccessTypeOffline,
	)

	fmt.Printf("Please authenticate on the following url:\n%s\n", authUrl)
	_ = open.Start(authUrl)

	tokenResponse := <-tokenResponseChan
	if err := tokenResponse.Err; err != nil {
		log.Fatal(err)
	}
	token := tokenResponse.Token

	log.Printf("Successfully authenticated")
	log.Printf("Got the access token: %v", token.AccessToken)

	// Create a new GitHub client using the access token http.Client.
	client := githubapi.NewClient(oauthClient.GetClient(context.Background(), token))

	next(token, client)
}

package internal

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	Accept                    = "Accept"
	ContentType               = "Content-Type"
	Authorization             = "Authorization"
	JSONContentType           = "application/json"
	FormUrlEncodedContentType = "application/x-www-form-urlencoded"
)

func CreateAccessTokenRequest(baseURL, username, password string) (*http.Request, error) {
	tokenURL := fmt.Sprintf("%s/CxRestAPI/auth/identity/connect/token", baseURL)
	data := url.Values{}
	data.Add("username", username)
	data.Add("password", password)
	data.Add("grant_type", "password")
	data.Add("scope", "access_control_api sast_api")
	data.Add("client_id", "resource_owner_sast_client")
	data.Add("client_secret", "014DF517-39D1-4453-B7B3-9930C563627C")
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add(ContentType, FormUrlEncodedContentType)
	req.Header.Add(Accept, FormUrlEncodedContentType)
	return req, nil
}

func CreateRequest(httpMethod, url string, requestBody io.Reader, token *AccessToken) (*http.Request, error) {
	resp, err := http.NewRequest(httpMethod, url, requestBody)
	if err != nil {
		return nil, err
	}

	resp.Header.Add(ContentType, JSONContentType)
	if token != nil {
		resp.Header.Add(Authorization, fmt.Sprintf("%s %s", token.TokenType, token.AccessToken))
	}
	return resp, nil
}

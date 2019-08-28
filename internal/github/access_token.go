// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    accessToken, err := UnmarshalAccessToken(bytes)
//    bytes, err = accessToken.Marshal()

package github

import "encoding/json"

// UnmarshalAccessToken handles the conversion from JSON to struct.
func UnmarshalAccessToken(data []byte) (AccessToken, error) {
	var r AccessToken
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal handles the conversion from struct to JSON.
func (r *AccessToken) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// AccessToken is a GitHub access token.
type AccessToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

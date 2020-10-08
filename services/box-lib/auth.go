package box_lib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)



type BoxToken struct {
	AccessToken  string   `json:"access_token"` // The token, the most important part!
	ExpiresIn    int      `json:"expires_in"` // expiry time, usually around 4000 (?)
	RestrictedTo []string `json:"restricted_to"`
	TokenType    string   `json:"token_type"`
}

func AuthOfBox(clientID string, clientSecret string, privateKey string) string {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		panic(err)
	}

	jti_str, err_jti := GenerateRandomString(64)
	if err_jti != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"iss":          "dsxx0wcdv75dp2a9k3cqnyk0i44he1sa", // boxAppSettings.clientID
		"sub":          "2250583", // enterpriseID
		"box_sub_type": "enterprise",
		"aud":          "https://api.box.com/oauth2/token",
		"jti":          jti_str,
		"exp":          time.Now().Unix() + 10, // in seconds
	})
	token.Header["kid"] = "5l00qwuy" // appAuth.publicKeyID

	tokenStr, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}
	// 4. Request access Token
	values := url.Values{}
	values.Add("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	//values.Add("grant_type", "authorization_code")
	values.Add("client_id", clientID) // boxAppSettings.clientID
	values.Add("client_secret", clientSecret) // boxAppSettings.clientSecret
	values.Add("assertion", tokenStr)
	values.Add("scope", "root_readwrite")

	req, err := http.NewRequest(http.MethodPost, "https://api.box.com/oauth2/token", strings.NewReader(values.Encode()))
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("https://api.box.com/oauth2/token failed:", resp)
		rb, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(rb))
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var boxToken BoxToken
	if err := json.Unmarshal(responseBody, &boxToken); err != nil {
		panic(err)
	}
	fmt.Printf("BOX TOKEN: %v expires in %v\n", boxToken.AccessToken, boxToken.ExpiresIn)
	return boxToken.AccessToken

}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
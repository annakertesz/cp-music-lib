package box_lib

import (
	"crypto/rand"
	"fmt"
	"log"
)

const (
	clientID     = "dsxx0wcdv75dp2a9k3cqnyk0i44he1sa"
	clientSecret = "zLYOmqvvXQSWbNaNA6bpho7IRXi0teTP"
	publicKeyID  = "gu01ngiz"
	privateKey   = "-----BEGIN ENCRYPTED PRIVATE KEY-----\nMIIFDjBABgkqhkiG9w0BBQ0wMzAbBgkqhkiG9w0BBQwwDgQIqODzFze8zUMCAggA\nMBQGCCqGSIb3DQMHBAj6SLCWU7xwmQSCBMgB0K/E4hZ1FvbMigwkdZbUL80zUy7U\nPA7xkX9wd++4VuLRWB/y+gV+MHeuI3EhpJvywZexQr6pS9peRQ6GJ0HdFEpMM/iA\ntP81jtZa/orG/g3z6kekisQ4toX6QV05d8Ljo+X6rq4ksWhhxrh+fGyKX9mhtc/i\nZASzUExjl8Pk+sO62nfXRCP3Oa/vN81E8vAgDP57mO3zb7SGc0/eePETnvVmu6W+\n37T+hXDuuEDnz2gVSqkMYLZuxCPQNxQA0jYv4OZIQapOC3gAgJlYinoYwZjcazSw\nwHwgKlnD7Byd96s1l6ANhgcosH2OJm3qOgb0T5ukEGhEh9O9XOaXNaZazO2HkNck\n49LfEz903GaZZRqCv702Vx12ljf9MCECX6ICr1pDf1WbASentrqLbCtB4bKfO3Yz\nWEvzwF0MIKOYYMMZijN8fGDabRMH/I1K4g1D187Y18sBYTpasydX6jv5Tuc9bexD\nmdaVoAyzO9xbqQy3m6ZkEySExRv9YgAYgcyBt1u55tafgZhJTt16KDtTNUa4/K+z\nHu4X3qaXRaMSsSl1FoMd7tNn8sRCGZStuIjik+UFynpFjIPB4c+7bQgm9KuJ4s3Q\nmXzKrZhmR57OFGSh9vO/s0KAhMNn7aFvxWyS1+S2lJ2Q3UicvqIEw+uu+9D7617v\npzexl6AUsFIMUjzaKvW1Gke6Df8JR9QTMtm7YTTwvYjXyI6pOvQ31aC3upSIzwr6\nvQIGNnthA8FqWV6o1AIc2xeLh/rOfCiJLE7uHQ0aPNfMJvnQZfP3qv8Hxxhc2ypG\nDGkq1bH45FJrHpA758VGxVI8paDOKvnizoX+08oD3X7B11gBAtSnDvHTXvfeqHT4\nZk5FYbPUpx/TKmPmRUxBnEjvMzIpFUHN37r8YhlwLUw3yVmu4RR9oqLhXMatAGQF\ndfTga29RQDX3EraKmRwDsCZ1VRHgWyeN7iGdiogNyRm1TnWvitJaXplGYZa/swK7\np4cbvIuM9j7DFZ54CNKBfz6+irxlxw3bF/wMowZrIuQufuFGUXwJ78dcpzLHyDGf\nMxdKWi/2fahyi83gBrN01yAzpZvvtMOCvJH08Bif/nYeb9i+Ks9EPS34GnoIR7Ad\n+ZiRBmlEI74P/Mnq+os7U3hd7fXwvC7nC3nWaOi61+0DAuZp7qszfbY3mPNUssgQ\nNRzMYgX6QNjjmg33aOSyeE7p5lKjnT0K9BQwvKAIrWKTq6kDjf4cTQ3kqLtrzjbu\nfXe0YWFUOrcKbsMP4CQtEXWWoVbdYFfEjXggfGGyQ+3I2kpmQRjTXqhsIoLSrnIH\nE08BwNQCaALhmSj+H1XkXufIHDLNRhlwBE82L5KbciykiOcEsFAc7rc4qA5xtfoa\nJK9h1dXmlYED+BHuc9HErJqtru9b3xmD6RbpvYLk3kA8zYlSmIAGig4NBFiDEFKQ\n7zek3mcFrCtczo/x+jhID+en5giLgbKg7bL0D10MPR0wmHnicWQT6A46naHIpxpt\n6PIfrNQAsftwySlZ8RKchj97enWamMddJwYmKUyGYIzoOLvPNl9uHi/sgIa9nsgu\n10lWy+QbSZLLEfPsI0Dra1sk8WplE2O6jX1K+P0IQOjN2Kr/y2mm5cjXvNyWjfY4\n8Js=\n-----END ENCRYPTED PRIVATE KEY-----\n"
	passphrase   = "f327bd602484192246ef695f4f9289cf"
	enterpriseID = "2250583"
)

//func encodePrivateKey() {
//	block, rest := pem.Decode([]byte(privateKey))
//	fmt.Printf()
//	claim := jwt.NewClaim()
//	claim.Set("iss", clientID)
//	claim.Set("sub", enterpriseID)
//	claim.Set("box_sub_type", "enterprise")
//	claim.Set("aud", "https://api.box.com/oauth2/token")
//	claim.Set("jti", generateUUID())
//	claim.Set("exp", 0.75)
//
//}

func generateUUID() string{
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

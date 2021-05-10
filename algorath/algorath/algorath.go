package algorath

type Credentials struct {
	APIKey	    string	`db:"APIKEY" json:"APIKey"`
	APISecret	string	`db:"APISECRET" json:"APISecret"`
}

var Running bool
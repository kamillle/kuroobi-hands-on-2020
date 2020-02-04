package config

import "fmt"

const (
	Port = 8080
)

// 1-3. Client ID、Client Secretを定義
const (
	ClientID     = "dj00aiZpPWJjUURCN0RpZ3dvZSZzPWNvbnN1bWVyc2VjcmV0Jng9OWM-"
	ClientSecret = "mQaUQeSHXdQ2UvPVivSW2Hyxkp1yZTKVUO2Uh2Hl"
)

// 1-4. リダイレクトURIを定義
var RedirectURI = fmt.Sprintf("http://localhost:%d/callback", Port)

// 1-5. OpenID ConnectのURLを定義
const (
	OIDCURL = "https://auth.login.yahoo.co.jp"
)

package config

import (
	"fmt"
	"os"
)

const (
	Port = 8080
)

// 1-3. Client ID、Client Secretを定義
const (
	ClientID     = os.Getenv("YAHOO_CLIENT_ID")
	ClientSecret = os.Getenv("YAHOO_CLIENT_SECRET")
)

// 1-4. リダイレクトURIを定義
var RedirectURI = fmt.Sprintf("http://localhost:%d/callback", Port)

// 1-5. OpenID ConnectのURLを定義
const (
	OIDCURL = "https://auth.login.yahoo.co.jp"
)

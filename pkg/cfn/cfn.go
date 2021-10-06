package cfn

import (
	"fmt"
	"net/http"
	"os"

	"github.com/antchfx/htmlquery"
)

const (
	cookieFile = "cookies.json"
	cfnURL     = "https://game.capcom.com/cfn/sfv"
)

type CFN struct {
	token string
}

func NewCFN(token string) *CFN {
	return &CFN{token: token}
}

func (c *CFN) GetProfile(fighterId string) Profile {

	// Create a cookie file with the token (usually 1st request)
	if _, err := os.Stat(cookieFile); os.IsNotExist(err) {
		var cookies []*http.Cookie
		cookies = append(cookies, &http.Cookie{Name: "scirid", Value: c.token})
		ExportCookies(cookies)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://game.capcom.com/cfn/sfv/profile/%s", fighterId), nil)
	if err != nil {
		fmt.Println("MAIN GET", err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	req.AddCookie(&http.Cookie{Name: "language", Value: "en"})

	for _, cookie := range ImportCookies() {
		req.AddCookie(cookie)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("MAIN DO", err)
	}
	defer resp.Body.Close()

	ExportCookies(resp.Cookies())

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		fmt.Println("PARSER", err)
	}

	return Profile{doc: doc}
}

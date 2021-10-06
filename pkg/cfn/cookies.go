package cfn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ExportCookies(cookies []*http.Cookie) {
	data, _ := json.MarshalIndent(cookies, "", " ")
	_ = ioutil.WriteFile(cookieFile, data, 0644)
}

func ImportCookies() (cookies []*http.Cookie) {
	data, _ := ioutil.ReadFile(cookieFile)
	json.Unmarshal(data, &cookies)
	return
}

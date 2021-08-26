package repo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/danisbagus/edagang-pkg/logger"
	"github.com/danisbagus/edagang-transaction/internal/core/port"
)

type AuthRepo struct{}

func NewAuthRepo() port.IAuthRepo {
	return &AuthRepo{}
}

func (r AuthRepo) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := generateVerifyURL(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

func generateVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:9010", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

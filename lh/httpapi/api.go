package httpapi

import (
	"io/ioutil"
	"net/http"
)

func GetHttpApi(_url string) (string, error) {
	resp, err := http.Get(_url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

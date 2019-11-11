package dnspod

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func reqPost(url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "go-dnspod (shallwedance@126.com)")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// DoRequest start to request
func DoRequest(url string, params url.Values) ([]byte, error) {
	// fmt.Println("Params: ", params.Encode())
	return reqPost(url, strings.NewReader(params.Encode()))
}

// DoRequestForm launch a post request with map
func DoRequestForm(url string, data map[string]string) ([]byte, error) {
	str := Encode(data)
	fmt.Println("Params: ", str)
	return reqPost(url, strings.NewReader(str))
}

// Encode return string from map[string]string
func Encode(data map[string]string) string {
	var strSlice []string
	for k, v := range data {
		strSlice = append(strSlice, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(strSlice, "&")
}

package dnspod

import (
	"fmt"
	"log"
	"regexp"
)

const (
	// DnspodURL dnspod api url
	DnspodURL = "https://dnsapi.cn/"
	// DnspodURL = "https://api.dnspod.com/"
)

// Client Config for dnspod
type Client struct {
	// LoginToken API token, https://support.dnspod.cn/Kb/showarticle/tsid/227/
	LoginToken    string
	LoginEmail    string
	LoginPassword string
	Format        string
	Lang          string
	ErrorOnEmpty  string
	Params        map[string]string
}

// InitClient return a dnspod client
func (cli *Client) InitClient() {
	if cli.Format == "" {
		cli.Format = "json"
	}

	if cli.Params == nil {
		cli.Params = map[string]string{
			"format": cli.Format,
		}
	}

	if cli.LoginToken != "" {
		cli.Params["login_token"] = cli.LoginToken
	} else {
		cli.Params["login_email"] = cli.LoginEmail
		cli.Params["login_password"] = cli.LoginPassword
	}

}

// Do will start a dnspod request
func (cli *Client) Do(action string, data map[string]string) {
	if data == nil {
		data = map[string]string{}
	}

	for k, v := range cli.Params {
		data[k] = v
	}

	url := DnspodURL + TransferAPI(action)

	resp, err := DoRequestForm(url, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp))
}

// TransferAPI 转换为 dnspod 可用的 Action 格式
// ex, DomainList -> Domain.List
func TransferAPI(origin string) string {
	re := regexp.MustCompile(`([A-Z])`)
	result := re.ReplaceAllString(origin, ".$1")
	return result[1:]
}

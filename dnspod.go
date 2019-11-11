package dnspod

import (
	"log"
	"regexp"
)

const (
	// DnspodAPIcn 国内版
	DnspodAPIcn = "https://dnsapi.cn/"
	// DnspodAPIcom 国际版
	DnspodAPIcom = "https://api.dnspod.com/"
)

// Client Config for dnspod
type Client struct {
	// LoginToken API token, https://support.dnspod.cn/Kb/showarticle/tsid/227/
	LoginToken string

	// LoginEmail 登录邮箱
	LoginEmail string
	// LoginPassword 登录密码
	LoginPassword string

	// Format 返回值格式
	Format string

	// Params 请求参数格式
	Params map[string]string

	// Region dnspod api 供应商
	// cn 国内版 ; com 国际版
	Region string
	// DnspodAPI
	DnspodAPI string
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

	if cli.Region == "com" {
		cli.DnspodAPI = DnspodAPIcom
	} else {
		cli.DnspodAPI = DnspodAPIcn
	}

	if cli.LoginToken != "" {
		cli.Params["login_token"] = cli.LoginToken
	} else {
		cli.Params["login_email"] = cli.LoginEmail
		cli.Params["login_password"] = cli.LoginPassword
	}

}

// Do will start a dnspod request
func (cli *Client) Do(action string, data map[string]string) []byte {
	if data == nil {
		data = map[string]string{}
	}

	for k, v := range cli.Params {
		data[k] = v
	}

	url := cli.DnspodAPI + TransferAPI(action)

	resp, err := DoRequestForm(url, data)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(resp))
	return resp
}

// TransferAPI 转换为 dnspod 可用的 Action 格式
// ex, DomainList -> Domain.List
func TransferAPI(origin string) string {
	re := regexp.MustCompile(`([A-Z])`)
	result := re.ReplaceAllString(origin, ".$1")
	return result[1:]
}

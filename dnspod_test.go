package dnspod

import (
	"fmt"
	"testing"
)

func Test_DoReq(t *testing.T) {

	token := userID + "," + userToken

	cli := Client{
		LoginToken: token,
		Region:     "cn",
	}
	cli.InitClient()

	resp := cli.Do("DomainList", nil)
	fmt.Println(string(resp))
}

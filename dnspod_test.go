package dnspod

import (
	"testing"
)

func Test_DoReq(t *testing.T) {

	token := userID + "," + userToken

	cli := Client{
		LoginToken: token,
	}
	cli.InitClient()

	cli.Do("DomainList", nil)
}

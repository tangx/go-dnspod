# go-dnspod

根据 `https://www.dnspod.cn/docs/index.html#` API 文档，自行构建 body(`map[string]string`)

## example

```go
token := userID + "," + userToken

cli := Client{
    LoginToken: token,
    Region:     "cn", // com 国际版
}

cli.InitClient()

resp := cli.Do("DomainList", nil) // 请求 https://apihost/Domain.List, 返回 json
fmt.Println(string(resp))
    
```
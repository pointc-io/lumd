package cmd

import (
	"log"
	"net/http"
	"github.com/elazarl/goproxy"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/aws/client"
)

func main() {
	cloudwatchlogs.New(client.ConfigProvider(}, aws.Config{}))
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		return resp
	})
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		return req, nil
	})
	log.Fatal(http.ListenAndServe(":8088", proxy))
}

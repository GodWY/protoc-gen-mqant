// Code generated by protoc-gen-go-hi. DO NOT EDIT.
// versions:

package greeter1

import (
	greeter "proto/examples/greeter"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
import (
	"errors"
	basemodule "github.com/liangdas/mqant/module/base"
	client "github.com/liangdas/mqant/module"
	mqrpc "github.com/liangdas/mqant/rpc"
	"golang.org/x/net/context"
)

// generated mqant method
type Greeter interface {
	//  @GET@gin.Logger()
	Hello(in *greeter.Request) (out *Response, err error)
	//  @POST
	Stream(in *greeter.Request) (out *Response, err error)
}

func RegisterGreeterTcpHandler(m *basemodule.BaseModule, ser Greeter) {
	m.GetServer().RegisterGO("hello", ser.Hello)
	m.GetServer().RegisterGO("stream", ser.Stream)
}

// generated proxxy handle
type ClientProxyService struct {
	cli  client.App
	name string
}

var ClientProxyIsNil = errors.New("proxy is nil")

func NewGreeterClient(cli client.App, name string) *ClientProxyService {
	return &ClientProxyService{
		cli:  cli,
		name: name,
	}
}
func (proxy *ClientProxyService) Hello(req *greeter.Request) (rsp *Response, err error) {
	if proxy == nil {
		return nil, ClientProxyIsNil
	}
	rsp = &Response{}
	err = mqrpc.Proto(rsp, func() (reply interface{}, err interface{}) {
		return proxy.cli.Call(context.TODO(), proxy.name, "hello", mqrpc.Param(req))
	})
	return rsp, err
}
func (proxy *ClientProxyService) Stream(req *greeter.Request) (rsp *Response, err error) {
	if proxy == nil {
		return nil, ClientProxyIsNil
	}
	rsp = &Response{}
	err = mqrpc.Proto(rsp, func() (reply interface{}, err interface{}) {
		return proxy.cli.Call(context.TODO(), proxy.name, "stream", mqrpc.Param(req))
	})
	return rsp, err
}

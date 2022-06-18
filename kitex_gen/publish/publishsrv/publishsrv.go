// Code generated by Kitex v0.3.2. DO NOT EDIT.

package publishsrv

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
	"kitexdousheng/kitex_gen/publish"
)

func serviceInfo() *kitex.ServiceInfo {
	return publishSrvServiceInfo
}

var publishSrvServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PublishSrv"
	handlerType := (*publish.PublishSrv)(nil)
	methods := map[string]kitex.MethodInfo{
		"PublishAction": kitex.NewMethodInfo(publishActionHandler, newPublishActionArgs, newPublishActionResult, false),
		"PublishList":   kitex.NewMethodInfo(publishListHandler, newPublishListArgs, newPublishListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "publish",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(publish.DouyinPublishActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(publish.PublishSrv).PublishAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishActionArgs:
		success, err := handler.(publish.PublishSrv).PublishAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishActionResult)
		realResult.Success = success
	}
	return nil
}
func newPublishActionArgs() interface{} {
	return &PublishActionArgs{}
}

func newPublishActionResult() interface{} {
	return &PublishActionResult{}
}

type PublishActionArgs struct {
	Req *publish.DouyinPublishActionRequest
}

func (p *PublishActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishActionArgs) Unmarshal(in []byte) error {
	msg := new(publish.DouyinPublishActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishActionArgs_Req_DEFAULT *publish.DouyinPublishActionRequest

func (p *PublishActionArgs) GetReq() *publish.DouyinPublishActionRequest {
	if !p.IsSetReq() {
		return PublishActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishActionResult struct {
	Success *publish.DouyinPublishActionResponse
}

var PublishActionResult_Success_DEFAULT *publish.DouyinPublishActionResponse

func (p *PublishActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishActionResult) Unmarshal(in []byte) error {
	msg := new(publish.DouyinPublishActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishActionResult) GetSuccess() *publish.DouyinPublishActionResponse {
	if !p.IsSetSuccess() {
		return PublishActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*publish.DouyinPublishActionResponse)
}

func (p *PublishActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(publish.DouyinPublishListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(publish.PublishSrv).PublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishListArgs:
		success, err := handler.(publish.PublishSrv).PublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishListResult)
		realResult.Success = success
	}
	return nil
}
func newPublishListArgs() interface{} {
	return &PublishListArgs{}
}

func newPublishListResult() interface{} {
	return &PublishListResult{}
}

type PublishListArgs struct {
	Req *publish.DouyinPublishListRequest
}

func (p *PublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishListArgs) Unmarshal(in []byte) error {
	msg := new(publish.DouyinPublishListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishListArgs_Req_DEFAULT *publish.DouyinPublishListRequest

func (p *PublishListArgs) GetReq() *publish.DouyinPublishListRequest {
	if !p.IsSetReq() {
		return PublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishListResult struct {
	Success *publish.DouyinPublishListResponse
}

var PublishListResult_Success_DEFAULT *publish.DouyinPublishListResponse

func (p *PublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishListResult) Unmarshal(in []byte) error {
	msg := new(publish.DouyinPublishListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishListResult) GetSuccess() *publish.DouyinPublishListResponse {
	if !p.IsSetSuccess() {
		return PublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*publish.DouyinPublishListResponse)
}

func (p *PublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PublishAction(ctx context.Context, Req *publish.DouyinPublishActionRequest) (r *publish.DouyinPublishActionResponse, err error) {
	var _args PublishActionArgs
	_args.Req = Req
	var _result PublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, Req *publish.DouyinPublishListRequest) (r *publish.DouyinPublishListResponse, err error) {
	var _args PublishListArgs
	_args.Req = Req
	var _result PublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

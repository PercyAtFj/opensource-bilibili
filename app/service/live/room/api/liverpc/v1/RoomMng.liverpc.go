// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/RoomMng.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// =================
// RoomMng Interface
// =================

type RoomMngRPCClient interface {
	// * 开通直播间
	//
	CreateRoom(ctx context.Context, req *RoomMngCreateRoomReq, opts ...liverpc.CallOption) (resp *RoomMngCreateRoomResp, err error)

	// * 获取监控列表
	//
	GetSecondVerifyList(ctx context.Context, req *RoomMngGetSecondVerifyListReq, opts ...liverpc.CallOption) (resp *RoomMngGetSecondVerifyListResp, err error)

	// * 查询是否是黑名单，没有roomid参数时返回全部黑名单map
	//
	IsBlack(ctx context.Context, req *RoomMngIsBlackReq, opts ...liverpc.CallOption) (resp *RoomMngIsBlackResp, err error)

	// * 主播填写公告命中审核词生成审核记录写入db
	//
	SaveHistory(ctx context.Context, req *RoomMngSaveHistoryReq, opts ...liverpc.CallOption) (resp *RoomMngSaveHistoryResp, err error)

	// * 是否全网封禁
	//
	IsAllNetBanned(ctx context.Context, req *RoomMngIsAllNetBannedReq, opts ...liverpc.CallOption) (resp *RoomMngIsAllNetBannedResp, err error)
}

// =======================
// RoomMng Live Rpc Client
// =======================

type roomMngRPCClient struct {
	client *liverpc.Client
}

// NewRoomMngRPCClient creates a client that implements the RoomMngRPCClient interface.
func NewRoomMngRPCClient(client *liverpc.Client) RoomMngRPCClient {
	return &roomMngRPCClient{
		client: client,
	}
}

func (c *roomMngRPCClient) CreateRoom(ctx context.Context, in *RoomMngCreateRoomReq, opts ...liverpc.CallOption) (*RoomMngCreateRoomResp, error) {
	out := new(RoomMngCreateRoomResp)
	err := doRPCRequest(ctx, c.client, 1, "RoomMng.createRoom", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomMngRPCClient) GetSecondVerifyList(ctx context.Context, in *RoomMngGetSecondVerifyListReq, opts ...liverpc.CallOption) (*RoomMngGetSecondVerifyListResp, error) {
	out := new(RoomMngGetSecondVerifyListResp)
	err := doRPCRequest(ctx, c.client, 1, "RoomMng.getSecondVerifyList", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomMngRPCClient) IsBlack(ctx context.Context, in *RoomMngIsBlackReq, opts ...liverpc.CallOption) (*RoomMngIsBlackResp, error) {
	out := new(RoomMngIsBlackResp)
	err := doRPCRequest(ctx, c.client, 1, "RoomMng.isBlack", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomMngRPCClient) SaveHistory(ctx context.Context, in *RoomMngSaveHistoryReq, opts ...liverpc.CallOption) (*RoomMngSaveHistoryResp, error) {
	out := new(RoomMngSaveHistoryResp)
	err := doRPCRequest(ctx, c.client, 1, "RoomMng.saveHistory", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomMngRPCClient) IsAllNetBanned(ctx context.Context, in *RoomMngIsAllNetBannedReq, opts ...liverpc.CallOption) (*RoomMngIsAllNetBannedResp, error) {
	out := new(RoomMngIsAllNetBannedResp)
	err := doRPCRequest(ctx, c.client, 1, "RoomMng.isAllNetBanned", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

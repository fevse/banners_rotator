package grpcserver

import (
	"context"

	rotatorpb "github.com/fevse/banners_rotator/internal/server/grpc/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (r *RotatorServer) Add(_ context.Context, m *rotatorpb.SlotBanner) (*emptypb.Empty, error) {
	sid := m.SlotID
	bid := m.BannerID
	err := r.App.Storage.Add(int(sid), int(bid))
	return nil, err
}

func (r *RotatorServer) Delete(_ context.Context, m *rotatorpb.SlotBanner) (*emptypb.Empty, error) {
	sid := m.SlotID
	bid := m.BannerID
	err := r.App.Storage.Delete(int(sid), int(bid))
	return nil, err
}

func (r *RotatorServer) Click(_ context.Context, m *rotatorpb.BannerSlotGroup) (*emptypb.Empty, error) {
	sid := m.SlotID
	bid := m.BannerID
	gid := m.GroupID
	err := r.App.Storage.Click(int(bid), int(sid), int(gid))
	return nil, err
}

func (r *RotatorServer) Choose(_ context.Context, m *rotatorpb.SlotGroup) (*rotatorpb.Banner, error) {
	sid := m.SlotID
	gid := m.GroupID
	bid, err := r.App.Storage.Choose(int(sid), int(gid))
	pbbid := &rotatorpb.Banner{BannerID: int64(bid)}
	return pbbid, err
}

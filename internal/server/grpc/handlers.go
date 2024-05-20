package grpcserver

import (
	"context"
	"fmt"

	rotatorpb "github.com/fevse/banners_rotator/internal/server/grpc/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (r *RotatorServer) Add(_ context.Context, m *rotatorpb.SlotBanner) (*emptypb.Empty, error) {
	sid := m.SlotID
	bid := m.BannerID
	err := r.App.Storage.AddBannerToSlot(int(sid), int(bid))
	info := fmt.Sprintf("banner %v added to slot %v", bid, sid)
	r.Logger.Info(info)
	return nil, err
}

func (r *RotatorServer) Delete(_ context.Context, m *rotatorpb.SlotBanner) (*emptypb.Empty, error) {
	sid := m.SlotID
	bid := m.BannerID
	err := r.App.Storage.DeleteBannerFromSlot(int(sid), int(bid))
	info := fmt.Sprintf("banner %v deleted from slot %v", bid, sid)
	r.Logger.Info(info)
	return nil, err
}

func (r *RotatorServer) Click(_ context.Context, m *rotatorpb.BannerSlotGroup) (*emptypb.Empty, error) {
	sid := m.SlotID
	bid := m.BannerID
	gid := m.GroupID
	err := r.App.Storage.ClickBanner(int(bid), int(sid), int(gid))
	info := fmt.Sprintf("banner %v was clicked from slot %v and group %v", bid, sid, gid)
	r.Logger.Info(info)
	return nil, err
}

func (r *RotatorServer) Choose(_ context.Context, m *rotatorpb.SlotGroup) (*rotatorpb.Banner, error) {
	sid := m.SlotID
	gid := m.GroupID
	bid, err := r.App.Storage.ChooseBannerToShow(int(sid), int(gid))
	pbbid := &rotatorpb.Banner{BannerID: int64(bid)}
	info := fmt.Sprintf("banner %v was choosen for slot %v and group %v", bid, sid, gid)
	r.Logger.Info(info)
	return pbbid, err
}

package master

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/yasshi2525/rushhourm/mock_pb"
	"github.com/yasshi2525/rushhourm/pb"
)

type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestCreateResidence(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_pb.NewMockModelServiceClient(ctrl)

	request := &pb.Residence{
		Id:       1,
		OwnerID:  1,
		Pos:      &pb.Point{X: 1, Y: 1},
		Capacity: 1,
	}

	client.
		EXPECT().
		CreateResidence(
			gomock.Any(), &rpcMsg{msg: request},
		).
		Return(request, nil)

	testCreateResidence(t, client)
}

func testCreateResidence(t *testing.T, client pb.ModelServiceClient) {
	res, err := client.CreateResidence(context.Background(), &pb.Residence{
		Id:       1,
		OwnerID:  1,
		Pos:      &pb.Point{X: 1, Y: 1},
		Capacity: 1,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log("response: ", res)
}

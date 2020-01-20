package slave

import "github.com/yasshi2525/rushhourm/pb"

import "net"

import "google.golang.org/grpc"

import "context"

type model struct {
	pb.UnimplementedModelServiceServer
	residences map[int64]*pb.Residence
	companies  map[int64]*pb.Company
}

func (m *model) CreateResidence(ctx context.Context, in *pb.Residence) (*pb.Residence, error) {
	m.residences[in.GetId()] = in
	return in, nil
}

func listenModel(m *model, address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterModelServiceServer(s, m)

	if err := s.Serve(l); err != nil {
		return err
	}

	return nil
}

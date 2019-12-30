package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/yasshi2525/rushhourm/pb"
	"google.golang.org/grpc"
)

type childPos int

const (
	nw childPos = iota
	ne
	sw
	se
)

type nodeInfo struct {
	address string
}

type serverInfo struct {
	pb.UnimplementedRegisterServiceServer
	parent   *nodeInfo
	children []*nodeInfo
	level    int64
}

func (s *serverInfo) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if len(s.children) > 4 {
		return nil, fmt.Errorf("no space to add child")
	}
	ch := &nodeInfo{address: req.GetAddress()}
	s.children = append(s.children, ch)
	res := &pb.RegisterResponse{Level: s.level + 1}
	log.Printf("[server] accept child %s, set level = %d\n", req.GetAddress(), res.GetLevel())
	return res, nil
}

func main() {
	myInfo := &serverInfo{}
	myHost := flag.String("host", "localhost", "this server's host name")
	myPort := flag.Int("port", 8080, "this server's port number")
	parentAddress := flag.String("parent", "", "parent address if exists")
	flag.Parse()

	myAddress := fmt.Sprintf("%s:%d", *myHost, *myPort)

	if *parentAddress != "" {
		myInfo.parent = &nodeInfo{address: *parentAddress}
		conn, err := grpc.Dial(*parentAddress, grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			panic(fmt.Errorf("failed to connect %s: %v", *parentAddress, err))
		}
		client := pb.NewRegisterServiceClient(conn)
		res, err := client.Register(context.Background(), &pb.RegisterRequest{Address: myAddress})
		if err != nil {
			panic(fmt.Errorf("failed to register %v", err))
		}
		myInfo.level = res.GetLevel()
		log.Printf("[client] register %s to %s, level = %d\n", myAddress, *parentAddress, res.GetLevel())
	}

	lis, err := net.Listen("tcp", myAddress)
	if err != nil {
		panic(fmt.Errorf("failed to listen port %d: %v", myPort, err))
	}

	srv := grpc.NewServer()
	pb.RegisterRegisterServiceServer(srv, myInfo)
	log.Printf("[server] listen %s\n", myAddress)
	if err := srv.Serve(lis); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}

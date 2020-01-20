package master

import "google.golang.org/grpc"

import "github.com/yasshi2525/rushhourm/pb"

import "context"

import "log"

const modelAddress = ":8081"

func createResidence() {
	conn, err := grpc.Dial(modelAddress, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatal("failed to connect model service: ", err.Error())
		return
	}

	client := pb.NewModelServiceClient(conn)
	_, err = client.CreateResidence(context.Background(), &pb.Residence{})

	if err != nil {
		log.Fatal("failed to create residence: ", err.Error())
	}
}

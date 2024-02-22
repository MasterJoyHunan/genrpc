package rpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// UserClient user.UserClient
)

func Setup() {
	// UserClient = user.NewUserClient(setup("127.0.0.1:8989"))
}


func setup(target string) *grpc.ClientConn {
	timeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	dial, err := grpc.DialContext(timeout,
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	return dial
}
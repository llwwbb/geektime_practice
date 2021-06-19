package v1

import (
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestUserClient_GetUser(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	c := NewUserClient(conn)
	r, err := c.GetUser(context.Background(), &GetUserReq{
		Id: "60cdfee5152aba1df2617bfd",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

package user_service_client

import (
	"fmt"
	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
	"github.com/google/uuid"
	"testing"
)

func TestCreateUser(t *testing.T) {
	client, _ := NewUserServiceClient("localhost", 1, 8081)

	response, err := client.CreateUser(&pb.User{
		Id:        1,
		Uuid:      uuid.New().String(),
		FirstName: "Hossein",
		LastName:  "Araghi",
	})

	if err != nil {
		return
	}

	fmt.Println(response)
}

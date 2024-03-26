package userapi

import (
	"context"
	"log"
	"testing"
)

func TestUserFromJwt(t *testing.T) {
	api := &UserApiClient{}
	err := api.Init("https://user.env0.app0.luojm.com:9443/query")
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	usersJson, err := api.GetUsersJson("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkyMTk5MDUsInVzZXIiOiJ7XCJpZFwiOlwiODQ2MjI3OTM3XCIsXCJyb2xlc1wiOltcIm5vcm1hbFwiXX0ifQ.J_8ZlZ5EKdc1WrrOQzlmN4cq4M34S9ZA-zj05WjhCXXV0rYpLyIDLKRNaTE13jlxkY68cOYc6S_IUa8TsprmhFD-0b_fbM5f0eurTdMBvbhmCxqa71-WSSh1nqNJnOP3OzJ_TSRQlh4ghVdv9ShUYoJBKO2KE60M2FbqKBtJBps0ZnQHI1NBHmNgnOHlRoG860QNtrJTfGsz9P_Px1yiSZPtMYRnQkWKtw6RsNvD4eKin_zhMl8e0qVnnfGjkm6q8ekIOcQ_Bp7NHmipTD8a4e8EB_gGEuE9TJG2qHfeO0obSqfCjvWoK0LvRaAixxnmqekoxWssVpNeeysM67-0cA", ctx, []string{"846227937"})
	if err != nil {
		t.Error(err)
	}
	log.Printf("%#v\n", usersJson)
	// if user.Id != "pdKBcAc7lKzSC9nyglCwQ"{
	// 	t.Errorf("userFromJwt want test, but: %s", user.Id)
	// }
}

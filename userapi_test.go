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
	usersJson, err := api.GetUsersInfo("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkyMTk5MDUsInVzZXIiOiJ7XCJpZFwiOlwiODQ2MjI3OTM3XCIsXCJyb2xlc1wiOltcIm5vcm1hbFwiXX0ifQ.J_8ZlZ5EKdc1WrrOQzlmN4cq4M34S9ZA-zj05WjhCXXV0rYpLyIDLKRNaTE13jlxkY68cOYc6S_IUa8TsprmhFD-0b_fbM5f0eurTdMBvbhmCxqa71-WSSh1nqNJnOP3OzJ_TSRQlh4ghVdv9ShUYoJBKO2KE60M2FbqKBtJBps0ZnQHI1NBHmNgnOHlRoG860QNtrJTfGsz9P_Px1yiSZPtMYRnQkWKtw6RsNvD4eKin_zhMl8e0qVnnfGjkm6q8ekIOcQ_Bp7NHmipTD8a4e8EB_gGEuE9TJG2qHfeO0obSqfCjvWoK0LvRaAixxnmqekoxWssVpNeeysM67-0cA",
		ctx,
		[]string{"846227937", "758350760", "300335865"})
	if err != nil {
		t.Error(err)
	}
	log.Printf("%#v\n", usersJson)
	// if user.Id != "pdKBcAc7lKzSC9nyglCwQ"{
	// 	t.Errorf("userFromJwt want test, but: %s", user.Id)
	// }
}


func TestGetUserName(t *testing.T) {
	api := &UserApiClient{}
	err := api.Init("http://127.0.0.1:5556/query")
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	name, err := api.GetUserName("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjgxMTY4OTUsInVzZXIiOiJ7XCJpZFwiOlwiODQ2MjI3OTM3XCIsXCJyb2xlc1wiOltcIm5vcm1hbFwiXX0ifQ.STEWtvoms7y4qmNf4AFRgyrjpK_Nyc3CMzkEFZ3zQkllXPgt6U2tL3fQCdoaWxAvXpidBctTKi46rNr_eIhQF_kR35Pn8moQldnh1w8rJpDogmX2EB0RdNLFWwHnqHlmfweYyE2Ir7CcXy5tZuyqC09u2dve2j5_FNPYxnvOkkhyMMGpKqdO49PECeWtWq9qbwEXa8EpTJcO825kQ10VqmHyUMMCrytW3Ifq4Q6GkokGxIjcXiHo0KJiAMqh-_Xrt3JrZoki5OOATZNfuLO_pUflSGOdACcGEJx8cZFqbHnq0SZkYkMJULdZuYRin1vHKtDzdBL1SPTwWcpoDTzjEQ",
		ctx,
		"846227937")
	if err != nil {
		t.Error(err)
	}
	log.Printf("%#v\n", name)
	// if user.Id != "pdKBcAc7lKzSC9nyglCwQ"{
	// 	t.Errorf("userFromJwt want test, but: %s", user.Id)
	// }
}

package userapi

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

type UserApiClient struct {
	client *graphql.Client
}

func (client *UserApiClient) Init(endpoint string) error {
	if endpoint == "" {
		return fmt.Errorf("endpoint can't empty")
	}
	client.client = graphql.NewClient(endpoint)
	return nil
}

func (client *UserApiClient) GetUsersInfo(token string, ctx context.Context, ids []string) (map[string]string, error) {

	req := graphql.NewRequest(`
	query getUsers($ids: [ID!]!){
		getUsers(ids:$ids){
			id
			info
		}
	  }
`)
	req.Var("ids", ids)
	req.Header.Set("authorization", "Bearer "+token)
	var response map[string]any
	if err := client.client.Run(ctx, req, &response); err != nil {
		return nil, err
	}
	usersJson, ok := response["getUsers"].([]any)
	if !ok {
		return nil, fmt.Errorf("response[getUsers] parse error")
	}
	usersInfo := map[string]string{}
	for _, userJson := range usersJson {
		userInfo, ok := userJson.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("userJson.(map[string]string) parse error")
		}
		userId, ok := userInfo["id"].(string)
		if !ok {
			return nil, fmt.Errorf("userInfo[id] parse error")
		}
		userInfoJson, ok := userInfo["info"].(string)
		if !ok {
			return nil, fmt.Errorf("userInfo[info] parse error")
		}
		usersInfo[userId] = userInfoJson
	}
	return usersInfo, nil
}

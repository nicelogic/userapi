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
		return fmt.Errorf("user api client endpoint is empty")
	}
	client.client = graphql.NewClient(endpoint)
	return nil
}

func (client *UserApiClient) GetUsersJson(token string, ctx context.Context, ids []string, ) (string, error) {

	req := graphql.NewRequest(`
	query getUsers($ids: [ID!]!){
		getUsers(ids:$ids)
	  }
`)
	req.Var("ids", ids)
	req.Header.Set("authorization", "Bearer "+token)
	var response map[string]interface{}
	if err := client.client.Run(ctx, req, &response); err != nil {
		return "", err
	}
	usersJson, ok := response["getUsers"].(string)
	if !ok {
		return "", fmt.Errorf("response parse error")
	}
	return usersJson, nil
}


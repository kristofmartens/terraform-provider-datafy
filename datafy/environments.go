package datafy

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Environment struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	DeletionProtection bool   `json:"deletionProtection"`
	Id                 string `json:"id"`
	TenantId           string `json:"tenantId"`
	State              string `json:"state"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
}

type Environments struct {
	Environments []Environment `json:"environments"`
}

func (c *Client) GetEnvironments(state string) (*Environments, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/environments?state=%s", c.HostURL, state), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var envs Environments
	err = json.Unmarshal(body, &envs)
	if err != nil {
		return nil, err
	}

	return &envs, nil
}

func (c *Client) GetEnvironment(id string) (*Environment, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/environments/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	env := Environment{}
	err = json.Unmarshal(body, &env)
	if err != nil {
		return nil, err
	}

	return &env, nil
}

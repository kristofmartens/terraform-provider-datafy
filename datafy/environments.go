package datafy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

type EnvironmentInput struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	DeletionProtection bool   `json:"deletionProtection"`
}

type EnvironmentUpdate struct {
	DeletionProtection bool `json:"deletionProtection"`
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

func (c *Client) CreateEnvironment(input *EnvironmentInput) (*Environment, error) {
	in, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/environments", c.HostURL), strings.NewReader(string(in)))
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

func (c *Client) UpdateEnvironment(id string, update *EnvironmentUpdate) (*Environment, error) {
	// TODO: something is still wrong with the update FIXME
	up, err := json.Marshal(update)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/environments/%s", c.HostURL, id), strings.NewReader(string(up)))
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

func (c *Client) DeleteEnvironment(id string) (*Environment, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/environments/%s", c.HostURL, id), nil)
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

package datafy

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Project struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	GitRepo      string `json:"gitRepo"`
	State        string `json:"state"`
	TenantId     string `json:"tenantId"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	LastActivity string `json:"lastActivity"`
}

func (c *Client) GetProject(id string) (*Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/projects/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var project Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

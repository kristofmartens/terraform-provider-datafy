package datafy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

type ProjectInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	GitRepo     string `json:"gitRepo"`
}

type ProjectUpdate struct {
	Description string `json:"description"`
	GitRepo     string `json:"gitRepo"`
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

func (c *Client) CreateProject(input *ProjectInput) (*Project, error) {
	in, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/projects", c.HostURL), strings.NewReader(string(in)))
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

func (c *Client) UpdateProject(id string, input *ProjectUpdate) (*Project, error) {
	in, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/projects/%s/info", c.HostURL, id), strings.NewReader(string(in)))
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

func (c *Client) DeleteProject(id string) (*Project, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/projects/%s", c.HostURL, id), nil)
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

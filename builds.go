package main

import (
	"fmt"
	"net/http"
)

// BuildsService handles communication with the builds
// related methods of the Travis CI API.
type BuildsService struct {
	client *Client
}

// Build represents a Travis CI build
type Build struct {
	Id                uint   `json:"id,omitempty"`
	RepositoryId      uint   `json:"repository_id,omitempty"`
	Slug              string `json:"slug,omitempty"`
	CommitId          uint   `json:"commit_id,omitempty"`
	Number            string `json:"number,omitempty"`
	PullRequest       bool   `json:"pull_request,omitempty"`
	PullRequestTitle  string `json:"pull_request_title,omitempty"`
	PullRequestNumber uint   `json:"pull_request_number,omitempty"`
	State             string `json:"state,omitempty"`
	StartedAt         string `json:"started_at,omitempty"`
	FinishedAt        string `json:"finished_at,omitempty"`
	Duration          uint   `json:"duration,omitempty"`
	JobIds            []uint `json:"job_ids,omitempty"`
	AfterNumber       uint   `json:"after_number,omitempty"`
	EventType         string `json:"event_type,omitempty"`
}

// ListBuildsResponse represents the response of a call
// to the Travis CI list builds endpoint.
type ListBuildsResponse struct {
	Builds []Build `json:"builds"`
}

// GetBuildResponse represents the response of a call
// to the Travis CI get build endpoint.
type GetBuildResponse struct {
	Build Build `json:"build"`
	// Jobs  []Job `json:"jobs"`
}

// BuildListOptions specifies the optional parameters to the
// BuildsService.List method.
type BuildListOptions struct {
	Ids          []uint `url:"ids,omitempty"`
	RepositoryId uint   `url:"repository_id,omitempty"`
	Slug         string `url:"slug,omitempty"`
	Number       string `url:"number,omitempty"`
	AfterNumber  uint   `url:"after_number,omitempty"`
	EventType    string `url:"event_type,omitempty"`
}

// List the builds for the authenticated user.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BuildsService) List(opt *BuildListOptions) ([]Build, *http.Response, error) {
	u, err := urlWithOptions("/builds", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var buildsResp ListBuildsResponse
	resp, err := bs.client.Do(req, &buildsResp)
	if err != nil {
		return nil, resp, err
	}

	return buildsResp.Builds, resp, err
}

// List a repository builds based on it's provided slug.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BuildsService) ListByRepository(slug string, opt *BuildListOptions) ([]Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repos/%v/builds", slug), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var buildsResp ListBuildsResponse
	resp, err := bs.client.Do(req, &buildsResp)
	if err != nil {
		return nil, resp, err
	}

	return buildsResp.Builds, resp, err
}

// Get fetches a build based on the provided id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BuildsService) Get(id uint) (*Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/builds/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var buildResp GetBuildResponse
	resp, err := bs.client.Do(req, &buildResp)
	if err != nil {
		return nil, resp, err
	}

	return &buildResp.Build, resp, err
}

// Cancel build with the provided id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BuildsService) Cancel(id uint) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/builds/%d/cancel", id), nil)
	if err != nil {
		return nil, err
	}

	req, err := bs.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := bs.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Restart build with the provided id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BuildsService) Restart(id uint) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/builds/%d/restart", id), nil)
	if err != nil {
		return nil, err
	}

	req, err := bs.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := bs.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
package leetcode

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    "https://leetcode.com/graphql",
	}
}

type GraphQLRequest struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

type Question struct {
	Content    string `json:"content"`
	Title      string `json:"title"`
	TitleSlug  string `json:"titleSlug"`
	Difficulty string `json:"difficulty"`
}

type GraphQLResponse struct {
	Data struct {
		Question     Question `json:"question"`
		QuestionList struct {
			Data []Question `json:"data"`
		} `json:"questionList"`
	} `json:"data"`
}

func (c *Client) executeQuery(query string, variables interface{}) (*GraphQLResponse, error) {
	reqBody := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response GraphQLResponse
	err = json.Unmarshal(body, &response)
	return &response, err
}
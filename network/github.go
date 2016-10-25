package network

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GitHubSearchResult struct {
	Name        string
	Stars       int
	Description string
	PackagePath string
}

type GitHubSearchResults struct {
	Results []GitHubSearchResult
}

type Item struct {
	Name        string `json:"name"`
	Stars       int    `json:"stargazers_count"`
	Description string `json:"description"`
	FullName    string `json:"full_name"`
}

type GitHubApiResponse struct {
	TotalCount int    `json:"total_count"`
	Items      []Item `json:"items"`
}

func PerformGitHubSearch(query string) *GitHubSearchResults {
	url := "https://api.github.com/search/repositories?q=" + query + "+language:go&sort=stars&order=desc"
	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var gitHubApiResponse GitHubApiResponse
	err = json.Unmarshal(body, &gitHubApiResponse)

	results := make([]GitHubSearchResult, 5)

	for i, item := range gitHubApiResponse.Items[:5] {
		results[i] = GitHubSearchResult{item.Name, item.Stars, item.Description, item.FullName}
	}

	return &GitHubSearchResults{results}
}

package main

import (
	"fmt"
	"net/url"
	"path"

	"github.com/imroc/req"
)

type giteaUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type giteaRepoInfo struct {
	RepoName    string `json:"repo_name"`
	CloneAddr   string `json:"clone_addr"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Mirror      bool   `json:"mirror"`
	UID         int    `json:"uid"`
}

func giteaGet(httpUrl string, cfg configGitea) (*req.Resp, error) {
	u, err := url.Parse(cfg.Host)
	if err != nil {
		return nil, err
	}
	fmt.Println(u)
	u.Path = path.Join("api/v1", httpUrl)
	fmt.Println(u)

	r, err := req.Get(u.String(),
		req.Header{
			"Authorization": "token " + cfg.Token,
			"Accept":        "application/json",
			"Content-Type":  "application/json",
		})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return r, nil
}

func giteaGetUID(cfg configGitea) int {
	r, err := giteaGet(path.Join("orgs", cfg.User), cfg)
	if err != nil {
		panic(err)
	}

	var user giteaUser
	if err := r.ToJSON(&user); err != nil {
		panic(err)
	}

	return user.ID
}

func giteaPost(httpUrl string, repo giteaRepoInfo, cfg configGitea) (*req.Resp, error) {
	u, err := url.Parse(cfg.Host)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join("api/v1", httpUrl)

	r, err := req.Post(u.String(),
		req.Header{
			"Authorization": "token " + cfg.Token,
			"Accept":        "application/json",
			"Content-Type":  "application/json",
		}, req.BodyJSON((&repo)))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return r, nil
}

func giteaMigrate(repo giteaRepoInfo, cfg configGitea) {
	r, err := giteaPost("repos/migrate", repo, cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.ToString())
}

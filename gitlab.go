package main

import (
	"fmt"
	"net/url"

	"github.com/imroc/req"
)

type gitlabRepo struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	URL         string `json:"http_url_to_repo"`
}

type gitlab []gitlabRepo

func gitlabGetProjectList(cfg configGitlab) gitlab {
	u, err := url.Parse(cfg.Host)
	if err != nil {
		return nil
	}
	u.Path = "api/v4/projects"

	r, err := req.Get(u.String(),
		req.Header{
			"PRIVATE-TOKEN": cfg.Token,
		},
		req.Param{
			// "page":     1,
			// "per_page": 100,
			"search": cfg.User,
		})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var list gitlab
	if err := r.ToJSON(&list); err != nil {
		fmt.Println(err)
	}

	return list
}

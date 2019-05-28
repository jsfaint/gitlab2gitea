package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Println("Migrate repository from GitLab to Gitea")
	cfg := getConfigFromFile("./config.ini")
	spew.Dump(cfg)

	spew.Println("Get Gitea UID")
	uid := giteaGetUID(cfg.Gitea)
	spew.Println("UID:", uid)

	spew.Println("Get GitLab repository list")
	gitlabRepoList := gitlabGetProjectList(cfg.Gitlab)
	spew.Dump(gitlabRepoList)

	panic("bye")

	spew.Println("Start repository migrating")
	for _, v := range gitlabRepoList {
		spew.Dump(v)
		giteaMigrate(giteaRepoInfo{
			RepoName:    v.Name,
			CloneAddr:   v.URL,
			Description: v.Description,
			UID:         uid,
		}, cfg.Gitea)
	}
}

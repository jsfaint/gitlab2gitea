package main

import (
	"gopkg.in/ini.v1"
)

type configGitea struct {
	Host  string `ini:"host"`
	User  string `ini:"user"`
	Token string `ini:"token"`
}

type configGitlab struct {
	Host  string `ini:"host"`
	User  string `ini:"username"`
	Token string `ini:"token"`
}

type config struct {
	Gitea  configGitea  `ini:"gitea"`
	Gitlab configGitlab `ini:"gitlab"`
}

func getConfigFromFile(name string) config {
	cfg, err := ini.Load(name)
	if err != nil {
		panic(err)
	}

	var c config
	if err := cfg.MapTo(&c); err != nil {
		panic(err)
	}

	return c
}

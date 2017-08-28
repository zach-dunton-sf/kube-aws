package model

type Image struct {
	Repo          string `yaml:",omitempty"`
	Tag           string `yaml:",omitempty"`
	RktPullDocker bool   `yaml:",omitempty"`
}

/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package datasource

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"go.uber.org/zap"
)

type Docker struct {
	Datasource
}

func (docker *Docker) Latest(repositoryName string) (tags []string, err error) {
	repository, err := name.NewRepository(repositoryName)
	if err != nil {
		zap.L().Error("Problem while parsing repository", zap.String("repository", repositoryName))
		return nil, fmt.Errorf("problem while parsing repository %s", repositoryName)
	}
	tags, err = listTags(repository)
	if err != nil {
		zap.L().Error("Problem while retrieving tags", zap.String("repository", repositoryName))
		return nil, fmt.Errorf("problem while retrieving tags for %s", repositoryName)
	}
	if len(tags) == 0 {
		zap.L().Error("Couldn't find any tags", zap.String("repository", repositoryName))
		return nil, fmt.Errorf("couldn't find any tags")
	}
	return tags, nil
}

func listTags(repository name.Repository) (tags []string, err error) {
	platform := remote.WithPlatform(v1.Platform{
		Architecture: "amd64",
		OS:           "linux",
	})

	tags, err = remote.List(repository, platform)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package datasource

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/shurcooL/githubv4"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type GithubTags struct {
	Datasource
}

func (githubTags *GithubTags) Latest(repositoryName string) (tags []string, err error) {
	ctx := context.Background()
	owner, repository, found := strings.Cut(repositoryName, "/")
	if !found {
		return nil, fmt.Errorf("repository name must be in format 'owner/repo'")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable is required for github-tags datasource")
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(ctx, src)
	client := githubv4.NewClient(httpClient)

	var query struct {
		Repository struct {
			Refs struct {
				Nodes []struct {
					Name string
				}
			} `graphql:"refs(refPrefix: \"refs/tags/\", first: 100, orderBy: {field: TAG_COMMIT_DATE, direction: DESC})"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(repository),
	}

	err = client.Query(ctx, &query, variables)
	if err != nil {
		zap.L().Error("Problem while getting GitHub tags", zap.String("repository", repositoryName), zap.Error(err))
		return nil, fmt.Errorf("problem while getting tags for %s: %w", repositoryName, err)
	}

	if len(query.Repository.Refs.Nodes) == 0 {
		return nil, fmt.Errorf("couldn't find any tags")
	}

	for _, node := range query.Repository.Refs.Nodes {
		tags = append(tags, node.Name)
	}

	return tags, nil
}

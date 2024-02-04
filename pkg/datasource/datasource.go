/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package datasource

import (
	"errors"
	"go.uber.org/zap"
)

type IDatasource interface {
	Latest(dependencyName string) (tags []string, err error)
}

type Datasource struct {
}

func Get(datasource string) (IDatasource, error) {
	datasources := map[string]func() IDatasource{
		"docker": func() IDatasource { return &Docker{} },
	}
	datasourceFactory, exists := datasources[datasource]
	if !exists {
		return nil, errors.New("datasource not found")
	}
	zap.L().Debug("Found datasource", zap.String("datasource", datasource))
	return datasourceFactory(), nil
}

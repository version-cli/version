/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package datasource

import (
	"errors"
)

type IDatasource interface {
	Latest(dependencyName string) (tags []string, err error)
}

type Datasource struct {
}

func Get(datasource string) (IDatasource, error) {
	switch datasource {
	case "docker":
		return &Docker{}, nil
	}
	return nil, errors.New("datasource not found")
}

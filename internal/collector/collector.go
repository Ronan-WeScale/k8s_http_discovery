package collector

import "context"

type Target struct {
	URL    string
	Labels map[string]string
}

type Collector interface {
	Name() string
	Collect(ctx context.Context) ([]Target, error)
}

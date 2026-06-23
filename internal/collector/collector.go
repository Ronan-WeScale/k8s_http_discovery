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

// AnnotationProbePath is the annotation key used to override the monitored
// path on any supported route resource.
const AnnotationProbePath = "k8s-http-discovery.io/probe-path"

// probePath returns the probe path annotation value if set, else empty string.
// A non-empty result means the collector should emit one target per host
// pointing to this path instead of the route's declared paths.
func probePath(annotations map[string]string) string {
	return annotations[AnnotationProbePath]
}

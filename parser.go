package dockerparser

import (
	"strings"

	"github.com/crowley-io/docker-parser/docker"
)

// Name returns the image's name. (ie: debian[:8.2])
func Name(remote string) (string, error) {
	return parse(remote, func(e docker.Named, t string) string {
		return e.RemoteName() + t
	})
}

// ShortName returns the image's name (ie: debian)
func ShortName(remote string) (string, error) {
	return parse(remote, func(e docker.Named, t string) string {
		return e.RemoteName()
	})
}

// Tag returns the image's tag.
func Tag(remote string) (string, error) {
	return parse(remote, func(e docker.Named, t string) string {
		if len(t) > 1 {
			return t[1:]
		}
		return ""
	})
}

// Registry returns the image's registry. (ie: host[:port])
func Registry(remote string) (string, error) {
	return parse(remote, func(e docker.Named, t string) string {
		return e.Hostname()
	})
}

// Repository returns the image's repository. (ie: registry/name)
func Repository(remote string) (string, error) {
	return parse(remote, func(e docker.Named, t string) string {
		return e.FullName()
	})
}

// Remote returns the image's remote identifier. (ie: registry/name[:tag])
func Remote(remote string) (string, error) {
	return parse(remote, func(e docker.Named, t string) string {
		return e.FullName() + t
	})
}

func clean(url string) string {

	s := url

	if strings.HasPrefix(url, "http://") {
		s = strings.Replace(url, "http://", "", 1)
	} else if strings.HasPrefix(url, "https://") {
		s = strings.Replace(url, "https://", "", 1)
	}

	return s
}

type handler func(n docker.Named, t string) string

func parse(remote string, handle handler) (string, error) {

	n, err := docker.ParseNamed(clean(remote))

	if err != nil {
		return "", err
	}

	n = docker.WithDefaultTag(n)

	var t string
	switch x := n.(type) {
	case docker.Canonical:
		t = "@" + x.Digest().String()
	case docker.NamedTagged:
		t = ":" + x.Tag()
	}

	return handle(n, t), nil
}

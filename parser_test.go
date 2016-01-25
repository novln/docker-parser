package dockerparser

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortParse(t *testing.T) {

	remote := "foo/bar"

	check(t, Name, "foo/bar", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "", remote)
	check(t, Registry, "docker.io", remote)
	check(t, Repository, "docker.io/foo/bar", remote)
	check(t, Remote, "docker.io/foo/bar", remote)

}

func TestShortParseWithTag(t *testing.T) {

	remote := "foo/bar:1.1"

	check(t, Name, "foo/bar:1.1", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "1.1", remote)
	check(t, Registry, "docker.io", remote)
	check(t, Repository, "docker.io/foo/bar", remote)
	check(t, Remote, "docker.io/foo/bar:1.1", remote)

}

func TestShortParseWithDigest(t *testing.T) {

	remote := "foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb"

	check(t, Name, "foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", remote)
	check(t, Registry, "docker.io", remote)
	check(t, Repository, "docker.io/foo/bar", remote)
	check(t, Remote, "docker.io/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", remote)

}

func TestRegistry(t *testing.T) {

	remote := "localhost.localdomain/foo/bar"

	check(t, Name, "foo/bar", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "", remote)
	check(t, Registry, "localhost.localdomain", remote)
	check(t, Repository, "localhost.localdomain/foo/bar", remote)
	check(t, Remote, "localhost.localdomain/foo/bar", remote)

}

func TestRegistryWithTag(t *testing.T) {

	remote := "localhost.localdomain/foo/bar:1.1"

	check(t, Name, "foo/bar:1.1", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "1.1", remote)
	check(t, Registry, "localhost.localdomain", remote)
	check(t, Repository, "localhost.localdomain/foo/bar", remote)
	check(t, Remote, "localhost.localdomain/foo/bar:1.1", remote)

}

func TestRegistryWithDigest(t *testing.T) {

	remote := "localhost.localdomain/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb"

	check(t, Name, "foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", remote)
	check(t, Registry, "localhost.localdomain", remote)
	check(t, Repository, "localhost.localdomain/foo/bar", remote)
	check(t, Remote, "localhost.localdomain/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", remote)

}

func TestRegistryWithPort(t *testing.T) {

	remote := "localhost.localdomain:5000/foo/bar"

	check(t, Name, "foo/bar", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "", remote)
	check(t, Registry, "localhost.localdomain:5000", remote)
	check(t, Repository, "localhost.localdomain:5000/foo/bar", remote)
	check(t, Remote, "localhost.localdomain:5000/foo/bar", remote)

}

func TestRegistryWithPortAndTag(t *testing.T) {

	remote := "localhost.localdomain:5000/foo/bar:1.1"

	check(t, Name, "foo/bar:1.1", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "1.1", remote)
	check(t, Registry, "localhost.localdomain:5000", remote)
	check(t, Repository, "localhost.localdomain:5000/foo/bar", remote)
	check(t, Remote, "localhost.localdomain:5000/foo/bar:1.1", remote)

}

func TestRegistryWithPortAndDigest(t *testing.T) {

	remote := "localhost.localdomain:5000/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb"

	check(t, Name, "foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", remote)
	check(t, Registry, "localhost.localdomain:5000", remote)
	check(t, Repository, "localhost.localdomain:5000/foo/bar", remote)
	check(t, Remote, "localhost.localdomain:5000/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", remote)

}

func TestHttpRegistryClean(t *testing.T) {

	remote := "http://localhost.localdomain:5000/foo/bar:latest"

	check(t, Name, "foo/bar:latest", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "latest", remote)
	check(t, Registry, "localhost.localdomain:5000", remote)
	check(t, Repository, "localhost.localdomain:5000/foo/bar", remote)
	check(t, Remote, "localhost.localdomain:5000/foo/bar:latest", remote)

}

func TestHttpsRegistryClean(t *testing.T) {

	remote := "https://localhost.localdomain:5000/foo/bar:latest"

	check(t, Name, "foo/bar:latest", remote)
	check(t, ShortName, "foo/bar", remote)
	check(t, Tag, "latest", remote)
	check(t, Registry, "localhost.localdomain:5000", remote)
	check(t, Repository, "localhost.localdomain:5000/foo/bar", remote)
	check(t, Remote, "localhost.localdomain:5000/foo/bar:latest", remote)

}

func TestParseError(t *testing.T) {

	remote := "sftp://user:passwd@example.com/foo/bar:latest"
	s, err := Name(remote)
	assert.NotEmpty(t, err)
	assert.Empty(t, s)

}

func check(t *testing.T, parse func(string) (string, error), expected, remote string) {

	s, err := parse(remote)

	if err != nil {
		assert.FailNow(t, "An error has occurred with the given parser")
	}

	assert.Equal(t, expected, s)

}

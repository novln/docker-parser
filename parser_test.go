//
// Copyright (C) 2015  Thomas LE ROUX <thomas@november-eleven.fr>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package dockerparser

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortParse(t *testing.T) {

	remote := "foo/bar"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar:latest", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "latest", reference.Tag())
		assert.Equal(t, "docker.io", reference.Registry())
		assert.Equal(t, "docker.io/foo/bar", reference.Repository())
		assert.Equal(t, "docker.io/foo/bar:latest", reference.Remote())
	}
}

func TestShortParseWithTag(t *testing.T) {

	remote := "foo/bar:1.1"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar:1.1", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "1.1", reference.Tag())
		assert.Equal(t, "docker.io", reference.Registry())
		assert.Equal(t, "docker.io/foo/bar", reference.Repository())
		assert.Equal(t, "docker.io/foo/bar:1.1", reference.Remote())
	}
}

func TestShortParseWithDigest(t *testing.T) {

	remote := "foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", reference.Tag())
		assert.Equal(t, "docker.io", reference.Registry())
		assert.Equal(t, "docker.io/foo/bar", reference.Repository())
		assert.Equal(t, "docker.io/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", reference.Remote())
	}
}

func TestRegistry(t *testing.T) {

	remote := "localhost.localdomain/foo/bar"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar:latest", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "latest", reference.Tag())
		assert.Equal(t, "localhost.localdomain", reference.Registry())
		assert.Equal(t, "localhost.localdomain/foo/bar", reference.Repository())
		assert.Equal(t, "localhost.localdomain/foo/bar:latest", reference.Remote())
	}
}

func TestRegistryWithTag(t *testing.T) {

	remote := "localhost.localdomain/foo/bar:1.1"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar:1.1", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "1.1", reference.Tag())
		assert.Equal(t, "localhost.localdomain", reference.Registry())
		assert.Equal(t, "localhost.localdomain/foo/bar", reference.Repository())
		assert.Equal(t, "localhost.localdomain/foo/bar:1.1", reference.Remote())
	}
}

func TestRegistryWithDigest(t *testing.T) {

	remote := "localhost.localdomain/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", reference.Tag())
		assert.Equal(t, "localhost.localdomain", reference.Registry())
		assert.Equal(t, "localhost.localdomain/foo/bar", reference.Repository())
		assert.Equal(t, "localhost.localdomain/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", reference.Remote())
	}
}

func TestRegistryWithPort(t *testing.T) {

	remote := "localhost.localdomain:5000/foo/bar"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar:latest", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "latest", reference.Tag())
		assert.Equal(t, "localhost.localdomain:5000", reference.Registry())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar", reference.Repository())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar:latest", reference.Remote())
	}
}

func TestRegistryWithPortAndTag(t *testing.T) {

	remote := "localhost.localdomain:5000/foo/bar:1.1"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar:1.1", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "1.1", reference.Tag())
		assert.Equal(t, "localhost.localdomain:5000", reference.Registry())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar", reference.Repository())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar:1.1", reference.Remote())
	}
}

func TestRegistryWithPortAndDigest(t *testing.T) {

	remote := "localhost.localdomain:5000/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", reference.Tag())
		assert.Equal(t, "localhost.localdomain:5000", reference.Registry())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar", reference.Repository())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar@sha256:bc8813ea7b3603864987522f02a76101c17ad122e1c46d790efc0fca78ca7bfb", reference.Remote())
	}
}

func TestHttpRegistryClean(t *testing.T) {

	remote := "http://localhost.localdomain:5000/foo/bar:latest"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar:latest", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "latest", reference.Tag())
		assert.Equal(t, "localhost.localdomain:5000", reference.Registry())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar", reference.Repository())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar:latest", reference.Remote())
	}
}

func TestHttpsRegistryClean(t *testing.T) {

	remote := "https://localhost.localdomain:5000/foo/bar:latest"

	if reference := parse(t, remote); reference != nil {
		assert.Equal(t, "foo/bar:latest", reference.Name())
		assert.Equal(t, "foo/bar", reference.ShortName())
		assert.Equal(t, "latest", reference.Tag())
		assert.Equal(t, "localhost.localdomain:5000", reference.Registry())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar", reference.Repository())
		assert.Equal(t, "localhost.localdomain:5000/foo/bar:latest", reference.Remote())
	}
}

func TestParseError(t *testing.T) {

	remote := "sftp://user:passwd@example.com/foo/bar:latest"
	x, err := Parse(remote)
	assert.NotEmpty(t, err)
	assert.Nil(t, x)

}

func parse(t *testing.T, remote string) *Reference {

	reference, err := Parse(remote)

	if err != nil {
		assert.FailNow(t, "An error has occurred with the parser")
		return nil
	}
	assert.NotNil(t, reference)
	assert.NotEmpty(t, reference)
	return reference

}

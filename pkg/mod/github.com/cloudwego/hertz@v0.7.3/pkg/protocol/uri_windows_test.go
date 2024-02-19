// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protocol

import (
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestURIPathNormalizeIssue86(t *testing.T) {
	t.Parallel()

	var u URI

	testURIPathNormalize(t, &u, `a`, `/a`)
	testURIPathNormalize(t, &u, "/../../../../../foo", "/foo")
	testURIPathNormalize(t, &u, "/..\\..\\..\\..\\..\\", "/")
	testURIPathNormalize(t, &u, "/..%5c..%5cfoo", "/foo")
}

func TestGetScheme(t *testing.T) {
	scheme, path := getScheme([]byte("E:\\file.go"))
	assert.DeepEqual(t, "", string(scheme))
	assert.DeepEqual(t, "E:\\file.go", string(path))

	scheme, path = getScheme([]byte("E:\\"))
	assert.DeepEqual(t, "", string(scheme))
	assert.DeepEqual(t, "E:\\", string(path))

	scheme, path = getScheme([]byte("https://foo.com"))
	assert.DeepEqual(t, "https", string(scheme))
	assert.DeepEqual(t, "//foo.com", string(path))

	scheme, path = getScheme([]byte("://"))
	assert.DeepEqual(t, "", string(scheme))
	assert.DeepEqual(t, "", string(path))

	scheme, path = getScheme([]byte("ws://127.0.0.1"))
	assert.DeepEqual(t, "ws", string(scheme))
	assert.DeepEqual(t, "//127.0.0.1", string(path))
}

// Copyright © 2019 Titan Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inmem

import (
	"testing"

	"github.com/atlaskerr/titan/core"

	"github.com/tidwall/buntdb"
)

func TestNewNamespace(t *testing.T) {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		t.Fatal("could not open database")
	}

	backend := Backend{db: db}

	params := core.NewNamespaceP{
		Name:         "test-ns",
		StorageLimit: 1234567,
		RepoLimit:    50,
		Labels: map[string]string{
			"label-one":   "key-one",
			"label-two":   "key-two",
			"label-three": "key-three",
		},
	}

	backend.NewNamespace(params)
}

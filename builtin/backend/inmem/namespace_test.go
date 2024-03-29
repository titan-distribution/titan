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

	resp := backend.NewNamespace(params)
	if resp.Error != nil {
		t.Fatal("failed to create namespace")
	}

	resp = backend.NewNamespace(params)
	if resp.Error == nil {
		t.Fatal("duplicate namespace created")
	}
}

func TestBatchNewNamespace(t *testing.T) {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		t.Fatal("could not open database")
	}

	backend := Backend{db: db}

	ns1 := core.NewNamespaceP{
		Name:         "test-ns-one",
		StorageLimit: 1234567,
		RepoLimit:    50,
		Labels: map[string]string{
			"label-one":   "key-one",
			"label-two":   "key-two",
			"label-three": "key-three",
		},
	}

	ns2 := core.NewNamespaceP{
		Name:         "test-ns-two",
		StorageLimit: 1234567,
		RepoLimit:    50,
		Labels: map[string]string{
			"label-one":   "key-one",
			"label-two":   "key-two",
			"label-three": "key-three",
		},
	}

	ns3 := core.NewNamespaceP{
		Name:         "test-ns-three",
		StorageLimit: 1234567,
		RepoLimit:    50,
		Labels: map[string]string{
			"label-one":   "key-one",
			"label-two":   "key-two",
			"label-three": "key-three",
		},
	}

	namespaces := core.BatchNewNamespaceP{
		Namespaces: []core.NewNamespaceP{ns1, ns2, ns3},
	}

	resp := backend.BatchNewNamespace(namespaces)
	if resp.Errors != nil {
		t.Fatal("failed to create batch namespaces")
	}

	resp = backend.BatchNewNamespace(namespaces)
	if resp.Errors == nil {
		t.Fatal("duplicate batch namespaces created")
	}

}

func TestDeleteNamespace(t *testing.T) {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		t.Fatal("could not open database")
	}

	backend := Backend{db: db}

	newParams := core.NewNamespaceP{
		Name:         "test-ns",
		StorageLimit: 1234567,
		RepoLimit:    50,
		Labels: map[string]string{
			"label-one":   "key-one",
			"label-two":   "key-two",
			"label-three": "key-three",
		},
	}

	newResp := backend.NewNamespace(newParams)
	if newResp.Error != nil {
		t.Fatal("failed to create namespace")
	}

	delParams := core.DelNamespaceP{
		Name: "test-ns",
	}

	delResp := backend.DelNamespace(delParams)
	if delResp.Error != nil {
		t.Fatal("failed to delete namespace")
	}

	duplicateDelResp := backend.DelNamespace(delParams)
	if duplicateDelResp.Error == nil {
		t.Fatal("namespace terminated more than one time")
	}

	notExistDelParams := core.DelNamespaceP{
		Name: "unknown",
	}
	notExistResp := backend.DelNamespace(notExistDelParams)
	if notExistResp.Error == nil {
		t.Fatal("nonexistent namespace deleted")
	}

}

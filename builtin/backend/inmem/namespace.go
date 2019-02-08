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
	"github.com/atlaskerr/titan/core"

	"github.com/tidwall/buntdb"
)

// NewNamespace creates a registry namespace.
func (b *Backend) NewNamespace(p core.NewNamespaceP) core.NewNamespaceR {
	var resp core.NewNamespaceR
	fields := make(map[string]string)
	fields["storage-limit"] = string(p.StorageLimit)
	fields["repo-limit"] = string(p.RepoLimit)
	fields["status"] = "active"

	for k, v := range p.Labels {
		key := join("label", k)
		fields[key] = v
	}

	fn := func(tx *buntdb.Tx) error {
		baseKey := join("namespace", p.Name)
		_, err := tx.Get(baseKey + ":status")
		if err != buntdb.ErrNotFound {
			return core.ErrNamespaceExists{Namespace: p.Name}
		}

		for k, v := range fields {
			fullKey := join(baseKey, k)
			tx.Set(fullKey, v, nil)
		}
		return nil
	}

	err := b.db.Update(fn)
	if err != nil {
		resp.Error = err
	}
	return resp
}

// BatchNewNamespace creates multiple namespaces in the registry.
//func (b *Backend) BatchNewNamespace(p core.BatchNewNamespaceP) core.BatchNewNamespaceR {
//	for _, ns := range p.Namespaces {
//		resp := b.NewNamespace(ns)
//		if resp.Error != nil {
//		}
//	}
//}

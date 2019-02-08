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
	fields := make(map[string]string)
	fields["storage-limit"] = string(p.StorageLimit)
	fields["repo-limit"] = string(p.RepoLimit)

	for k, v := range p.Labels {
		key := join("label", k)
		fields[key] = v
	}

	fn := func(tx *buntdb.Tx) error {
		for k, v := range fields {
			fullKey := join("namespace", p.Name, k)
			tx.Set(fullKey, v, nil)
		}

		return nil
	}

	b.db.Update(fn)
	return core.NewNamespaceR{Error: nil}
}

//func batchNewNamespace(db *buntdb.DB, p core.BatchNewNamespaceP) core.BatchNewNamespaceR {
//}

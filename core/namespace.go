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

package core

// Namespacer defines methods for working with registry namespaces.
type Namespacer interface {
	NamespaceCreater
	NamespaceDeleter
	NamespaceGetter
}

// NamespaceCreater defines methods for creating registry namespaces.
type NamespaceCreater interface {
	CreateNamespace(CreateNamespaceParams) error
	CreateNamespaces([]CreateNamespaceParams) error
}

// NamespaceDeleter defines methods for deleting registry namespaces.
type NamespaceDeleter interface {
	DeleteNamespace(name string) error
	DeleteNamespacesWithLabel(key, value string) (deleted []string, err error)
	DeleteAllNamespaces() error
}

// NamespaceGetter defines methods for fetching registry namespaces.
type NamespaceGetter interface {
	GetNamespace(name string) (*Namespace, error)
	GetNamespacesWithLabel(key, value string) (*[]Namespace, error)
	GetAllNamespaces() (*[]Namespace, error)
}

// CreateNamespaceParams defines parameters for creating a registry namespace.
type CreateNamespaceParams struct {
	Name         string
	StorageLimit uint64
	RepoLimit    uint64
	Labels       map[string]string
}

// Namespace contains information about registry namespaces.
type Namespace struct {
	Name         string
	StorageLimit uint64
	StorageUsed  uint64
	RepoLimit    uint64
	RepoCount    uint64
	Labels       map[string]string
	Status       string
}

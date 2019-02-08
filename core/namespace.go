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

import (
	"bytes"
)

// ErrNamespaceExists is the struct returned when a client attempts to create
// a new namespace with a name that already exists.
type ErrNamespaceExists struct {
	Namespace string
}

func (e ErrNamespaceExists) Error() string {
	buf := new(bytes.Buffer)
	buf.WriteString("namespace ")
	buf.WriteString(e.Namespace)
	buf.WriteString(" already exists")
	return buf.String()
}

// NamespaceBackend defines methods for working with registry namespaces.
type NamespaceBackend interface {
	NamespaceCreater
	NamespaceDeleter
	NamespaceFinder
}

// NamespaceCreater defines methods for creating registry namespaces.
type NamespaceCreater interface {
	NewNamespace(NewNamespaceP) NewNamespaceR
	BatchNewNamespace(BatchNewNamespaceP) BatchNewNamespaceR
}

// NewNamespaceP defines parameters for creating a registry namespace.
type NewNamespaceP struct {
	Name         string
	StorageLimit uint64
	RepoLimit    uint64
	Labels       map[string]string
}

// NewNamespaceR defines the response returned from NewNamespace.
type NewNamespaceR struct {
	Error error
}

// BatchNewNamespaceP is a wrapper around CreateNamespaceParams.
type BatchNewNamespaceP struct {
	Namespaces []NewNamespaceP
}

// BatchNewNamespaceR defines the response returned from BatchNewNamespace.
type BatchNewNamespaceR struct {
	Errors []error
}

// NamespaceDeleter defines methods for deleting registry namespaces.
type NamespaceDeleter interface {
	DelNamespace(DelNamespaceP) DelNamespaceR
	BatchDelNamespace(BatchDelNamespaceP) BatchDelNamespaceR
	PurgeNamespaces() PurgeNamespacesR
}

// DelNamespaceP defines parameters for DelNamespace.
type DelNamespaceP struct {
	Name string
}

// DelNamespaceR defines the response returned from DelNamespace.
type DelNamespaceR struct {
	Error error
}

// BatchDelNamespaceP defines parameters for BatchDelNamespace.
type BatchDelNamespaceP []DelNamespaceP

// BatchDelNamespaceR defines the response returned from BatchDelNamespace.
type BatchDelNamespaceR struct {
	Error error
}

// PurgeNamespacesR defines the response returned from PurgeNamespace.
type PurgeNamespacesR struct {
	Error error
}

// NamespaceFinder defines methods for fetching registry namespaces.
type NamespaceFinder interface {
	FindNamespace(FindNamespaceP) FindNamespaceR
	FindNamespaces(FindNamespacesP) FindNamespacesR
	AllNamespaces() AllNamespacesR
}

// FindNamespaceP defines parameters for FindNamespace.
type FindNamespaceP struct {
	Name string
}

// FindNamespaceR defines the response returned from FindNamespace.
type FindNamespaceR struct {
	Namespace Namespace
	Error     error
}

// FindNamespacesP defines parameters for FindNamespaces.
type FindNamespacesP struct {
	LabelKey   string
	LabelValue string
}

// FindNamespacesR defines the response returned from FindNamespaces.
type FindNamespacesR struct {
	Namespaces []Namespace
	Error      error
}

// AllNamespacesR defines the response returned from AllNamespaces.
type AllNamespacesR struct {
	Namespaces []Namespace
	Error      error
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

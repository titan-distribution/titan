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

// RepoBackend defines methods for working wih repositories.
type RepoBackend interface {
	RepoCreater
	RepoDeleter
}

// RepoCreater defines methods for creating repositories.
type RepoCreater interface {
	CreateRepo(*RepoConfig) error
	CreateRepos(*[]RepoConfig) error
}

// RepoDeleter defines methods for deleting repositories.
type RepoDeleter interface {
	DeleteRepo(namespace, repo string) error
	DeleteAllRepos(namespace string) error
	DeleteReposByLabel(namespace, key, value string) error
}

// RepoGetter defines methods for getting repositories.
type RepoGetter interface {
	GetRepo(namespace, repo string) (*Repo, error)
	GetReposWithLabel(namespace, key, value string) (*[]Repo, error)
	GetAllRepos(namespace string) (*[]Repo, error)
}

// Repo contains information about a repository.
type Repo struct {
	Name         string
	Namespace    string
	StorageLimit uint64
	StorageUsed  uint64
	Labels       map[string]string
	Status       string
}

// RepoConfig defines parameters for creating a repository.
type RepoConfig struct {
	Name         string
	Namespace    string
	StorageLimit uint64
	StorageUsed  uint64
	Labels       map[string]string
}

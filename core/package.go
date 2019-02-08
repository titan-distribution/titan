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
	"time"

	"github.com/opencontainers/go-digest"
)

// PackageUploader defines methods for uploading packages to a registry.
type PackageUploader interface {
	InitPackageUpload(InitPackageUploadParams) InitPackageUploadResp
	UploadBlobChunk(UploadBlobChunkParams)
}

// InitPackageUploadParams defines parameters for initiating the upload of
// a package.
type InitPackageUploadParams struct {
	Namespace string `json:"namespace"`
	Repo      string `json:"repo"`
	Blobs     []struct {
		Digest    digest.Digest `json:"digest"`
		MediaType string        `json:"mediaType"`
	} `json:"blobs"`
}

// InitPackageUploadResp defines the response returned from InitPackageUpload.
type InitPackageUploadResp struct {
	MaxChunkSize uint64 `json:"maxChunkSize"`
	Error        error  `json:"-"`
	UploadIDs    []struct {
		Namespace string        `json:"namespace"`
		Repo      string        `json:"repo"`
		Digest    digest.Digest `json:"digest"`
		UploadID  string        `json:"uploadID"`
		Expires   *time.Time    `json:"expires"`
	} `json:"uploadIDs"`
}

type UploadBlobChunkParams struct {
	Part uint64
}

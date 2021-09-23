// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// ThroughputRequestOptions includes options for throughput operations.
type ThroughputRequestOptions struct {
	IfMatchEtag     string
	IfNoneMatchEtag string
}

func (options *ThroughputRequestOptions) toHeaders() *map[string]string {
	if options.IfMatchEtag == "" && options.IfNoneMatchEtag == "" {
		return nil
	}

	headers := make(map[string]string)
	if options.IfMatchEtag != "" {
		headers[azcore.HeaderIfMatch] = options.IfMatchEtag
	}
	if options.IfNoneMatchEtag != "" {
		headers[azcore.HeaderIfNoneMatch] = options.IfNoneMatchEtag
	}
	return &headers
}
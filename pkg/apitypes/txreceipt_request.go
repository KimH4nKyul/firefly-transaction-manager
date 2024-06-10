// Copyright © 2024 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
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

package apitypes

import (
	"github.com/hyperledger/firefly-transaction-manager/pkg/ffcapi"
)

// TransactionReceiptRequest is the request payload to query for a receipt
type TransactionReceiptRequest struct {
	Headers RequestHeaders `json:"headers"`
	ffcapi.TransactionReceiptRequest
}

// TransactionReceiptResponse is the response payload for a query
type TransactionReceiptResponse struct {
	ffcapi.TransactionReceiptResponseBase
	Events []*EventWithContext `json:"events,omitempty"` // this is the serialization format for events (historical complexity)
}

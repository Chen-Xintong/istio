// Copyright Istio Authors
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

// nolint: golint // Avoid it complaining about the Fuzz function name; it is required
package fuzz

import (
	"istio.io/istio/pilot/pkg/config/kube/crd"
	"istio.io/istio/pilot/pkg/model"
)

func FuzzParseInputs(data []byte) int {
	_, _, err := crd.ParseInputs(string(data))
	if err != nil {
		return 0
	}
	return 1
}

func proxyValid(p *model.Proxy) bool {
	return len(p.IPAddresses) != 0
}

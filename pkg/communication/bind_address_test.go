/*
Copyright 2025 The llm-d-inference-sim Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package communication

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListenAddress", func() {
	DescribeTable("joins the configured bind address with the port",
		func(bindAddress string, port int, want string) {
			Expect(ListenAddress(bindAddress, port)).To(Equal(want))
		},
		Entry("empty bind-address keeps all interfaces", "", 8000, ":8000"),
		Entry("IPv4 loopback", "127.0.0.1", 8000, "127.0.0.1:8000"),
		Entry("IPv6 loopback is bracketed", "::1", 8000, "[::1]:8000"),
		Entry("non-default port", "127.0.0.1", 12345, "127.0.0.1:12345"),
	)
})

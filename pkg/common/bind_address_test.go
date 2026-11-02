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

package common

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BindAddress", func() {
	validConfig := func(bindAddress string) *Configuration {
		c := NewConfig()
		c.Model = "test-model"
		c.Mode = ModeEcho
		c.Port = 8000
		c.BindAddress = bindAddress
		return c
	}

	It("accepts an empty bind-address (all interfaces)", func() {
		Expect(validConfig("").validate()).To(Succeed())
	})

	DescribeTable("accepts valid IP literals", func(addr string) {
		Expect(validConfig(addr).validate()).To(Succeed())
	},
		Entry("IPv4 loopback", "127.0.0.1"),
		Entry("IPv4 any", "0.0.0.0"),
		Entry("IPv6 loopback", "::1"),
	)

	DescribeTable("rejects anything that is not an IP literal", func(addr string) {
		err := validConfig(addr).validate()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("invalid bind-address"))
	},
		Entry("hostname", "localhost"),
		Entry("dotted garbage", "127.0.0"),
		Entry("host:port", "127.0.0.1:8000"),
	)
})

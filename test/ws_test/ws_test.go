/*
 * Copyright 2018 The Service Manager Authors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package ws_test

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"

	"github.com/Peripli/service-manager/test/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWsConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Websocket test suite")
}

var _ = Describe("WS", func() {
	var ctx *common.TestContext

	BeforeSuite(func() {
		ctx = common.DefaultTestContext()
	})

	AfterSuite(func() {
		ctx.Cleanup()
	})

	Describe("test", func() {
		It("should work", func() {
			platform := ctx.RegisterPlatform()
			user := platform.Credentials.Basic.Username
			password := platform.Credentials.Basic.Password
			auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))

			smURL := ctx.Servers[common.SMServer].URL()
			smEndpoint, _ := url.Parse(smURL)

			conn, resp, err := websocket.DefaultDialer.Dial("ws://"+smEndpoint.Host+"/v1/notifications",
				http.Header{
					"Authorization": []string{auth},
				})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusSwitchingProtocols))
			var msg interface{}
			err = conn.ReadJSON(&msg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(msg).To(Equal("hello"))
		})
	})
})

/*
 * Copyright 2018 The Service Manager Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package basic

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	httpsec "github.com/Peripli/service-manager/pkg/security/http"

	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/util"
	"github.com/Peripli/service-manager/storage/storagefakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Basic Authenticator")
}

var _ = Describe("Basic Authenticator", func() {
	credentialsStorage := &storagefakes.FakeCredentials{}
	user := "user"
	password := "password"
	basicHeader := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, password)))
	var credentials *types.Credentials
	var authenticator httpsec.Authenticator

	BeforeEach(func() {
		credentials = &types.Credentials{
			Basic: &types.Basic{
				Username: user,
				Password: password,
			},
		}
		authenticator = &basicAuthenticator{CredentialStorage: credentialsStorage}
	})

	Describe("Authenticate", func() {
		var request *http.Request
		BeforeEach(func() {
			var err error
			request, err = http.NewRequest(http.MethodGet, "https://example.com", nil)
			Expect(err).ShouldNot(HaveOccurred())
		})
		Context("When authorization is not basic", func() {
			It("Should abstain", func() {
				request.Header.Add("Authorization", "Bearer token")
				user, decision, err := authenticator.Authenticate(request)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).To(BeNil())
				Expect(decision).To(Equal(httpsec.Abstain))
			})
		})
		Context("When user is not found", func() {
			BeforeEach(func() {
				credentialsStorage.GetReturns(nil, util.ErrNotFoundInStorage)
			})
			It("Should deny", func() {
				request.Header.Add("Authorization", "Basic "+basicHeader)
				user, decision, err := authenticator.Authenticate(request)
				Expect(err).To(Equal(util.ErrNotFoundInStorage))
				Expect(user).To(BeNil())
				Expect(decision).To(Equal(httpsec.Deny))
			})
		})

		Context("When getting credentials from storage results in error", func() {
			expectedError := fmt.Errorf("Error when fetching credentials from storage")
			BeforeEach(func() {
				credentialsStorage.GetReturns(nil, expectedError)
			})
			It("Should abstain with error", func() {
				request.Header.Add("Authorization", "Basic "+basicHeader)
				user, decision, err := authenticator.Authenticate(request)
				Expect(err.Error()).To(ContainSubstring(expectedError.Error()))
				Expect(user).To(BeNil())
				Expect(decision).To(Equal(httpsec.Abstain))
			})
		})

		Context("When passwords do not match", func() {
			BeforeEach(func() {
				credentialsStorage.GetReturns(credentials, nil)
			})
			It("Should deny", func() {
				request.Header.Add("Authorization", "Basic "+basicHeader)
				credentials.Basic.Password = "not-matching"
				user, decision, err := authenticator.Authenticate(request)
				Expect(err).To(BeNil())
				Expect(user).To(BeNil())
				Expect(decision).To(Equal(httpsec.Deny))
			})
		})

		Context("When passwords match", func() {
			BeforeEach(func() {
				credentialsStorage.GetReturns(credentials, nil)
			})
			It("Should allow", func() {
				request.Header.Add("Authorization", "Basic "+basicHeader)
				user, decision, err := authenticator.Authenticate(request)
				Expect(err).To(BeNil())
				Expect(user).To(Not(BeNil()))
				Expect(decision).To(Equal(httpsec.Allow))
			})
		})
	})
})

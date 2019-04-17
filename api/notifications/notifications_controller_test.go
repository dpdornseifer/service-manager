package notifications_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Notifications Controller Suite")
}

var _ = Describe("Notifications Controller", func() {

	It("should work", func() {
		fmt.Println("Work ...")
	})

})

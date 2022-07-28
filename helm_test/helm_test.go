package helm_test

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func findWord(key string, words []string) bool {
	for _, word := range words {
		if key == word {
			return true
		}
	}
	return false
}

func throwQuery(url string) []string {
	resp, err := http.Get(url)
	Expect(err).ShouldNot(HaveOccurred())
	defer resp.Body.Close()
	Expect(resp.StatusCode).To(Equal(http.StatusOK))
	bodyByte, err := io.ReadAll(resp.Body)
	Expect(err).ShouldNot(HaveOccurred())
	return strings.Split(strings.TrimSpace(string(bodyByte)), "\n")
}

var _ = Describe("Helm", func() {
	Describe("HTTP client can communicate with the wolper web server installed by Helm", func() {
		Context("With proper settings", func() {
			port := 8080
			It("Should get the correct result", func() {
				// Throw http requests
				root_url := "http://localhost:" + strconv.Itoa(port) + "/query"

				url := root_url + "?key=...l."
				words := throwQuery(url)
				Expect(findWord("world", words)).To(BeTrue())
				Expect(findWord("would", words)).To(BeTrue())
				Expect(findWord("hello", words)).To(BeFalse())
			})
		})
	})
})

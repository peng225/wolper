package web_test

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/peng225/wolper/dictionary"
	"github.com/peng225/wolper/pb"
	"github.com/peng225/wolper/service"
	"github.com/peng225/wolper/web"
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
	return []string{strings.TrimSpace(string(bodyByte))}
}

var _ = Describe("Web interface", func() {
	var (
		dict      dictionary.Dictionary
		registrar *grpc.Server
	)

	BeforeEach(func() {
		dict = dictionary.Dictionary{}
	})

	AfterEach(func() {
		registrar.GracefulStop()
	})

	Describe("HTTP client and wolper server communicate successfully via web server", func() {
		Context("With proper settings", func() {
			port := 8090
			wolper_port := 8091
			It("should get the correct result", func() {
				test_dict_name := "test_dict.txt"
				dict.Build("test_input", test_dict_name)
				// Start wolper server
				go func() {
					listenPort, err := net.Listen("tcp4", fmt.Sprintf(":%d", wolper_port))
					if err != nil {
						fmt.Println("failed to listen:", err)
					}
					registrar = grpc.NewServer()
					wssi := service.WolperServiceServerImpl{}
					wssi.Init(test_dict_name)
					pb.RegisterWolperServiceServer(registrar, &wssi)
					reflection.Register(registrar)
					registrar.Serve(listenPort)
				}()

				// Start web server
				go func() {
					web.Start(port, "localhost:"+strconv.Itoa(wolper_port), ".")
				}()

				time.Sleep(time.Second * 1)

				// Throw http requests
				root_url := "http://localhost:" + strconv.Itoa(port) + "/query"

				url := root_url + "?key=crane&include=&exclude="
				words := throwQuery(url)
				Expect(findWord("crane", words)).To(BeTrue())
				Expect(findWord("hello", words)).To(BeFalse())

				url = root_url + "?key=.r..e&include=a&exclude=x"
				words = throwQuery(url)
				Expect(findWord("crane", words)).To(BeTrue())

				url = root_url + "?key=.oda.&include=&exclude="
				words = throwQuery(url)
				Expect(findWord("today", words)).To(BeTrue())
			})
		})
	})
})

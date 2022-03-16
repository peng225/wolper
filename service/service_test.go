package service_test

import (
	"fmt"
	"net"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/peng225/wolper/dictionary"
	"github.com/peng225/wolper/pb"
	"github.com/peng225/wolper/service"
)

func findWord(key string, words []string) bool {
	for _, word := range words {
		if key == word {
			return true
		}
	}
	return false
}

var _ = Describe("Service", func() {
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

	Describe("Client and server communicate successfully", func() {
		Context("With proper settings", func() {
			port := 8080
			It("should get the correct result", func() {
				test_dict_name := "test_dict.txt"
				dict.Build("test_input", test_dict_name)
				// Start server
				go func() {
					listenPort, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
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

				// Start client
				words := service.ClientQuery("localhost:"+strconv.Itoa(port),
					"crane", "", "", false)
				Expect(findWord("crane", words)).To(BeTrue())
				Expect(findWord("hello", words)).To(BeFalse())

				words = service.ClientQuery("localhost:"+strconv.Itoa(port),
					".r..e", "a", "x", false)
				Expect(findWord("crane", words)).To(BeTrue())

				words = service.ClientQuery("localhost:"+strconv.Itoa(port),
					".oda.", "", "", false)
				Expect(findWord("today", words)).To(BeTrue())
			})
		})
	})

	Describe("Client and server fails to communicate", func() {
		Context("With wrong port number", func() {
			serverPort := 8080
			clientPort := 8081
			It("should fail to connect", func() {
				test_dict_name := "test_dict.txt"
				dict.Build("test_input", test_dict_name)
				// Start server
				go func() {
					listenPort, err := net.Listen("tcp4", fmt.Sprintf(":%d", serverPort))
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

				// Start client
				words := service.ClientQuery("localhost:"+strconv.Itoa(clientPort),
					"crane", "", "", false)

				Expect(words).To(BeNil())
			})
		})
	})
})

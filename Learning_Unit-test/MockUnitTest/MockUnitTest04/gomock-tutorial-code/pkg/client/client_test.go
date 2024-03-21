package client_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/walmartdigital/gomock-tutorial-code/pkg/client"
	"github.com/walmartdigital/gomock-tutorial-code/pkg/mocks"
)

var ctrl *gomock.Controller


func TestAll(t *testing.T){
	ctrl = gomock.NewController(t)
	defer ctrl.Finish()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Client tests")
}

var _ = Describe("Read message", func(){
	var (
		fakeHTTPClient *mocks.MockHTTPClient
		fakeHTTPClientFactory *mocks.MockHTTPClientFactory
		zooClient *client.ZooClient
	)

	BeforeEach(func(){
		fakeHTTPClient= mocks.NewMockHTTPClient(ctrl)
		fakeHTTPClientFactory = mocks.NewMockHTTPClientFactory(ctrl)
		fakeHTTPClientFactory.EXPECT().Create().Return(
			fakeHTTPClient,
		).Times(1)

		zooClient = client.NewZooClient(fakeHTTPClientFactory)
	})

	It("should read a message from the server", func(){
		fakeHTTPClient.EXPECT().Get("http://localhost:8080/dogs").Return(
			200,
			[]byte ("Hi there, I love dogs!"),
			nil,
		).Times(1)

		msg := zooClient.ReadMessage("dogs")
		Expect(msg).To(Equal("Hi there, I love dogs!"))
	})


	It("should answer that is doesn't know the provided animals", func(){
		fakeHTTPClient.EXPECT().Get("http://localhost:8080/elephants").Return(
			404,
			[]byte ("Not found"),
			nil,
		).Times(1)

		msg := zooClient.ReadMessage("elephants")
		Expect(msg).To(Equal("Hi there, what is elephant!"))
	})

	It("should answer that the zoo is closed", func(){
		fakeHTTPClient.EXPECT().Get("http://localhost:8080/dogs").Return(
			-1,
			nil,
			errors.New("Could not connect to the server"),
		).Times(1)

		msg := zooClient.ReadMessage("dogs")
		Expect(msg).To(Equal("Hi there, the zoo is closed!"))
	})
})

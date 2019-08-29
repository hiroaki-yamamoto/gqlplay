package gqlplay_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoGqlPlay(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoGqlPlay Suite")
}

package gqlplay_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoGqlPlay(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoGqlPlay Suite")
}

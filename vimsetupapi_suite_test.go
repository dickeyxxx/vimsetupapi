package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVimsetupApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vimsetup Api Suite")
}

package sql_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSQL(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SQL Suite")
}

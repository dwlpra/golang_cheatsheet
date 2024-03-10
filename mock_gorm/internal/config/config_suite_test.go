package config_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

type MockViper struct {
	Config map[string]string
}

func (m *MockViper) GetString(key string) string {
	return m.Config[key]
}

type MockOpener struct {
	mock.Mock
}

func (m *MockOpener) Open(dsn string) (*gorm.DB, error) {
	args := m.Called(dsn)
	return args.Get(0).(*gorm.DB), args.Error(1)
}

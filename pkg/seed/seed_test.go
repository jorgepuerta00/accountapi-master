package seed

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSeedTestSuite(t *testing.T) {
	suite.Run(t, new(SeedTestSuite))
}

type SeedTestSuite struct {
	suite.Suite
	seed *Seeding
}

func (t *SeedTestSuite) SetupTest() {
	logger := logrus.New()
	t.seed = NewSeeding(logger, os.Getenv("BASE_URL"))
}

func (t *SeedTestSuite) Test_Sedding() {
	data, err := t.seed.Seeding(2)
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), 2, len(data))
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

type MainTestSuite struct {
	suite.Suite
}

func (t *MainTestSuite) SetupTest() {
}

func (t *MainTestSuite) Test_Repo_Get() {
	assert.Equal(t.T(), true, true)
}

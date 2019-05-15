package main

import (
	"github.com/alibaba/pouch/test/environment"

	"github.com/go-check/check"
)

// PouchSearchSuite is the test suite for search CLI.
type PouchSearchSuite struct{}

func init() {
	check.Suite(&PouchSearchSuite{})
}

// SetUpSuite does common setup in the beginning of each test suite.
func (suite *PouchSearchSuite) SetUpSuite(c *check.C) {
	SkipIfFalse(c, environment.IsLinux)
}

// TestSearchWorks tests "pouch search" work.
func (suite *PouchSearchSuite) TestSearchWorks(c *check.C) {
	// TODO: here we need to add some test case
}

// TestSearchInSpecificRegistry test search in specific registry
func (suite *PouchSearchSuite) TestSearchInSpecificRegistry(c *check.C) {
	// TODO: Verification specific registry?
}

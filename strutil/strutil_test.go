package strutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "gopkg.in/check.v1"
	"testing"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type StrUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&StrUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *StrUtilSuite) TestConcat(c *C) {
	s1 := "abcdef"
	s2 := "123456"
	s3 := "ABCDEF"
	s4 := "!@#$%^"

	c.Assert(Concat(s1, s2, s3, s4), Equals, s1+s2+s3+s4)
}

func (s *StrUtilSuite) TestSubstr(c *C) {
	c.Assert(Substr("test1234TEST", 0, 8), Equals, "test1234")
	c.Assert(Substr("test1234TEST", 4, 8), Equals, "1234")
	c.Assert(Substr("test1234TEST", 8, 16), Equals, "TEST")
	c.Assert(Substr("test1234TEST", -1, 4), Equals, "test")
}

func (s *StrUtilSuite) TestEllipsis(c *C) {
	c.Assert(Ellipsis("Test1234", 8), Equals, "Test1234")
	c.Assert(Ellipsis("Test1234test", 8), Equals, "Test1...")
}

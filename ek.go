// +build !windows

// Package ek is set of auxiliary packages
package ek

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"golang.org/x/crypto/bcrypt"
	"pkg.re/essentialkaos/go-linenoise.v3"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// worthless is used as dependency fix
func worthless() {
	linenoise.Clear()
	bcrypt.Cost(nil)
}

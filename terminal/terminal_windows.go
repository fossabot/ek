// +build windows, !linux, !darwin

package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Prompt is prompt string
var Prompt = "> "

// MaskSymbol is symbol used for masking passwords
var MaskSymbol = "*"

// ErrKillSignal is error type when user cancel input
var ErrKillSignal = errors.New("")

// ////////////////////////////////////////////////////////////////////////////////// //

func ReadUI(title string, nonEmpty bool) (string, error) {
	return "", nil
}

func ReadAnswer(title, defaultAnswer string) (bool, error) {
	return true, nil
}

func ReadPassword(title string, nonEmpty bool) (string, error) {
	return "", nil
}

func PrintErrorMessage(message string, args ...interface{}) {
	return
}

func PrintWarnMessage(message string, args ...interface{}) {
	return
}

func PrintActionMessage(message string) {
	return
}

func PrintActionStatus(status int) {
	return
}

func AddHstory(ui string) {
	return
}

func SetCompletionHandler(compfunc func(in string) []string) {
	return
}

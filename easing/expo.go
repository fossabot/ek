package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ExpoIn Accelerating from zero velocity
func ExpoIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c*math.Pow(2, 10*(t/d-1)) + b
}

// ExpoOut Decelerating to zero velocity
func ExpoOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c*(-math.Pow(2, -10*t/d)+1) + b
}

// ExpoInOut Acceleration until halfway, then deceleration
func ExpoInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d / 2

	if t < 1 {
		return c/2*math.Pow(2, 10*(t-1)) + b
	}

	t--

	return c/2*(-math.Pow(2, -10*t)+2) + b
}

package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// QuadIn Accelerating from zero velocity
func QuadIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d

	return c*t*t + b
}

// QuadOut Decelerating to zero velocity
func QuadOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d

	return -c*t*(t-2) + b
}

// QuadInOut Acceleration until halfway, then deceleration
func QuadInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d / 2

	if t < 1 {
		return c/2*t*t + b
	}

	t--

	return -c/2*(t*(t-2)-1) + b
}

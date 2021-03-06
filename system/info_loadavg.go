// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v9/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with LA info in procfs
var procLoadAvgFile = "/proc/loadavg"

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[LOC,ABC]

// GetLA return loadavg
func GetLA() (*LoadAvg, error) {
	fd, err := os.OpenFile(procLoadAvgFile, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	text, _ := r.ReadString('\n')

	if len(text) < 20 {
		return nil, errors.New("Can't parse file " + procLoadAvgFile)
	}

	la := &LoadAvg{}

	la.Min1, err = strconv.ParseFloat(strutil.ReadField(text, 0, true), 64)

	if err != nil {
		return nil, errors.New("Can't parse field 0 as float number in " + procLoadAvgFile)
	}

	la.Min5, err = strconv.ParseFloat(strutil.ReadField(text, 1, true), 64)

	if err != nil {
		return nil, errors.New("Can't parse field 1 as float number in " + procLoadAvgFile)
	}

	la.Min15, err = strconv.ParseFloat(strutil.ReadField(text, 2, true), 64)

	if err != nil {
		return nil, errors.New("Can't parse field 2 as float number in " + procLoadAvgFile)
	}

	procs := strutil.ReadField(text, 3, true)
	delimPosition := strings.IndexRune(procs, '/')

	if delimPosition == -1 {
		return nil, errors.New("Can't parse field 3 in " + procLoadAvgFile)
	}

	la.RProc, err = strconv.Atoi(procs[:delimPosition])

	if err != nil {
		return nil, errors.New("Can't parse processes number in " + procLoadAvgFile)
	}

	la.TProc, err = strconv.Atoi(procs[delimPosition+1:])

	if err != nil {
		return nil, errors.New("Can't parse processes number in " + procLoadAvgFile)
	}

	return la, nil
}

// codebeat:enable[LOC,ABC]

package directio

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	BLOCK_SIZE = 4096 // Minimal block size
	ALIGN_SIZE = 0    // Align size
)

// ////////////////////////////////////////////////////////////////////////////////// //

func openFile(file string, flag int, perm os.FileMode) (*os.File, error) {
	fd, err := os.OpenFile(file, flag, perm)

	if err != nil {
		return nil, err
	}

	_, _, e := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd.Fd()), syscall.F_NOCACHE, 1)

	if e != 0 {
		fd.Close()
		return nil, errors.New("Can't set F_NOCACHE for given file")
	}

	return fd, nil
}

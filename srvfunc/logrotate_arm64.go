package srvfunc

import (
	"os"
	"syscall"
)

// LogRotate переоткрывает указанный файл и подменяем stdout/stderr вывод на этот файл
func LogRotate(prevLogFd *os.File, fname string) (newLogFd *os.File, err error) {
	if prevLogFd != nil {
		prevLogFd.Close()
		prevLogFd = nil
	}

	flag := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	newLogFd, err = os.OpenFile(fname, flag, os.FileMode(0644))
	if err != nil {
		return nil, err
	}

	syscall.Dup3(int(newLogFd.Fd()), syscall.Stdout, 0)
	syscall.Dup3(int(newLogFd.Fd()), syscall.Stderr, 0)

	return newLogFd, nil
}

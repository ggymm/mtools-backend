package utils

import "syscall"

// GetUsername 当前用户名
func GetUsername() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	user, err := syscall.UTF16PtrFromString("USERNAME")
	if err != nil {
		return "unknow"
	}
	r, err := syscall.GetEnvironmentVariable(user, &buffer[0], size)
	if err != nil {
		return "unknow"
	}
	return syscall.UTF16ToString(buffer[:r])
}

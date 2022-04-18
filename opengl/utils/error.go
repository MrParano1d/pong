package utils

import (
	"errors"
	"unsafe"
)

type GetGlParam func(uint32, uint32, *int32)
type GetInfoLog func(uint32, int32, *int32, *uint8)

func CheckGlError(glObject uint32, errorParam uint32, getParamFn GetGlParam,
	getInfoLogFn GetInfoLog) error {

	var success int32
	getParamFn(glObject, errorParam, &success)
	if success != 1 {
		var infoLog [512]byte
		getInfoLogFn(glObject, 512, nil, (*uint8)(unsafe.Pointer(&infoLog)))
		return errors.New(string(infoLog[:512]))
	}
	return nil
}

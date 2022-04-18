package utils

import "github.com/go-gl/mathgl/mgl32"

func GetMat4(val float32) mgl32.Mat4 {
	return mgl32.Mat4{val, val, val, val, val, val, val, val, val, val, val, val, val, val, val, val}
}

package opengl

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/mrparano1d/ecs"
	"go.uber.org/zap"
	"unsafe"
)

func DebugSystem() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			gl.Enable(gl.DEBUG_OUTPUT)
			gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS)
			logger := ecs.GetResource[*zap.Logger](resourceMap)
			gl.DebugMessageCallback(debugCallback(logger), nil)
		})
	}
}

func debugCallback(logger *zap.Logger) gl.DebugProc {
	return func(source uint32, gltype uint32, id uint32, severity uint32, length int32, message string, userParam unsafe.Pointer) {
		logger.Debug(
			message,
			zap.Uint32("source", source),
			zap.Uint32("gltype", gltype),
			zap.Uint32("id", id),
			zap.Uint32("severity", severity),
			zap.Int32("length", length),
		)
	}
}

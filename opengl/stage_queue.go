package opengl

import "github.com/mrparano1d/ecs"

const (
	StageQueue = "queue"
)

type QueueStage struct {
	ecs.Stage
}

var _ ecs.Stage = &QueueStage{}

func NewQueueStage() *QueueStage {
	return &QueueStage{
		Stage: ecs.NewDefaultStage(),
	}
}

func (QueueStage) Name() string {
	return StageQueue
}

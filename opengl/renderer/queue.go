package renderer

import "github.com/mrparano1d/ecs"

type RenderQueue []ecs.Entity

func (rq *RenderQueue) Push(entities ...ecs.Entity) {
	*rq = append(*rq, entities...)
}
func (rq *RenderQueue) Replace(entities ...ecs.Entity) {
	*rq = entities
}

type RenderQueueResource struct {
	Queue RenderQueue
}

package events

import "github.com/mrparano1d/ecs"

type CollisionSide string

const (
	CollisionLeft   CollisionSide = "left"
	CollisionRight  CollisionSide = "right"
	CollisionBottom CollisionSide = "bottom"
	CollisionTop    CollisionSide = "top"
)

type CollisionEvent struct {
	Entity1 ecs.Entity
	Entity2 ecs.Entity
	Side    CollisionSide
}

package camera

type Projections struct {
	Orthographic *OrthographicProjection
	Perspective  *PerspectiveProjection
}

type Cameras struct {
	Orthographic *Camera
	Perspective  *Camera
}

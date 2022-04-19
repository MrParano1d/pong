package camera

type ScalingMode string

const (
	ScalingModeNone            ScalingMode = "none"
	ScalingModeWindowSize                  = "window_size"
	ScalingModeFixedVertical               = "fixed_vertical"
	ScalingModeFixedHorizontal             = "fixed_horizontal"
)

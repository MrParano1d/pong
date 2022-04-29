package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mrparano1d/pong/ebiten_plugin/enums"
)

func ToEbitenFilter(filter enums.Filter) ebiten.Filter {
	switch filter {
	case enums.FilterNearest:
		return ebiten.FilterNearest
	default:
		return ebiten.FilterLinear
	}
}

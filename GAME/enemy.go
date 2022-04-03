package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySize = 111
)

type enemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newEnemy(renderer *sdl.Renderer, x, y float64) (e enemy) {
	e.tex = textureFromPNG(renderer, "assets/spaceship2.png")

	e.x = x
	e.y = y

	return e
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	X := e.x - enemySize/2
	Y := e.y - enemySize/2

	renderer.CopyEx(e.tex,
		&sdl.Rect{X: 0, Y: 0, W: enemySize, H: enemySize},
		&sdl.Rect{X: int32(X), Y: int32(Y), W: enemySize, H: enemySize},
		180,
		&sdl.Point{X: enemySize / 2, Y: enemySize / 2},
		sdl.FLIP_NONE)
}

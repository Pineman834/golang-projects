package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed   = 0.5
	playerSize    = 111
	shootCooldown = time.Millisecond * 250
)

type player struct {
	tex  *sdl.Texture
	x, y float64

	lastShot time.Time
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromPNG(renderer, "assets/spaceship.png")

	p.x = screenWidth / 2.0
	p.y = screenHeight - playerSize/2.0

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	X := p.x - playerSize/2
	Y := p.y - playerSize/2

	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(X), Y: int32(Y), W: playerSize, H: playerSize})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_A] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_D] == 1 {
		p.x += playerSpeed
	}
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) > shootCooldown {
			p.shoot()
			p.shoot()

			p.lastShot = time.Now()
		}
	}
}

func (p *player) shoot() {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.x = p.x
		bul.y = p.y
		bul.angle = 270 * (math.Pi / 180)
	}
}

package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 1.5
)

type bullet struct {
	tex   *sdl.Texture
	x, y  float64
	angle float64

	active bool
}

func newBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.tex = textureFromBMP(renderer, "sprites/player_bullet.bmp")

	return bul
}

func (bul *bullet) draw(renderer *sdl.Renderer) {
	if !bul.active {
		return
	}
	X := bul.x - bulletSize/2
	Y := bul.y - bulletSize/2

	renderer.Copy(bul.tex,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
		&sdl.Rect{X: int32(X), Y: int32(Y), W: bulletSize, H: bulletSize})
}

var bulletPool []*bullet

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, &bul)
	}
}

func bulletFromPool() (*bullet, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}

func (bul *bullet) update() {
	bul.x += bulletSpeed * math.Cos(bul.angle) // Ax = A cos θ	A = bulletSpeed	θ = angle
	bul.y += bulletSpeed * math.Sin(bul.angle) // Ay = A sin θ	A = bulletSpeed	θ = angle
}

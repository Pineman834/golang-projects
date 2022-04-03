package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Space Fighters",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	plr := newPlayer(renderer)

	var enemies []enemy

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (enemySize / 2.0)
			y := float64(j)*enemySize + (enemySize / 2.0)
			enemy := newEnemy(renderer, x, y)
			enemies = append(enemies, enemy)
		}
	}

	initBulletPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		plr.draw(renderer)
		plr.update()

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		for _, bul := range bulletPool {
			bul.draw(renderer)
			bul.update()
		}

		renderer.Present()
	}
}

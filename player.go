package main

import rl "github.com/gen2brain/raylib-go/raylib"

const PLAYER_HORIZONTAL_SPEED int = 5

type Player struct {
	position      rl.Vector2
	verticalSpeed float32
	moveRight     bool
	moveLeft      bool
}

func initPlayer() Player {
	return Player{
		position: rl.Vector2{
			X: WINDOW_WIDTH / 2,
			Y: WINDOW_HEIGHT/2 - 40,
		},
		verticalSpeed: 3,
		moveRight:     false,
		moveLeft:      false,
	}
}

func (p *Player) update() {
	if p.moveRight != p.moveLeft {
		if p.moveRight {
			p.position.X += float32(PLAYER_HORIZONTAL_SPEED)
		} else {
			p.position.X -= float32(PLAYER_HORIZONTAL_SPEED)
		}
	}

	p.position = rl.Vector2Clamp(p.position, rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: WINDOW_WIDTH - 16, Y: WINDOW_HEIGHT})
	p.moveLeft = false
	p.moveRight = false
}

func (p *Player) render() {
	rl.DrawRectangle(int32(p.position.X), int32(p.position.Y), 16, 16, rl.Blue)
}

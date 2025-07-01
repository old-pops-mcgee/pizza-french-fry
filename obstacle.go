package main

import rl "github.com/gen2brain/raylib-go/raylib"

const OBSTACLE_BUFFER int = -50

type Obstacle struct {
	texture  rl.Texture2D
	position rl.Vector2
}

func initObstacle(t rl.Texture2D, p rl.Vector2) Obstacle {
	return Obstacle{
		texture:  t,
		position: p,
	}
}

func (o *Obstacle) update(speed float32) {
	o.position.Y -= speed
}

func (o *Obstacle) render() {
	// Placeholder
	rl.DrawRectangle(int32(o.position.X), int32(o.position.Y), 16, 16, rl.Red)
}

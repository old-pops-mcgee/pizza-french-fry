package main

import rl "github.com/gen2brain/raylib-go/raylib"

const WINDOW_WIDTH float32 = 340
const WINDOW_HEIGHT float32 = 600

type Game struct {
	playerPosition rl.Vector2
}

func initGame() Game {
	return Game{
		playerPosition: rl.Vector2{X: WINDOW_WIDTH / 2, Y: 20},
	}
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)

	// TODO: Sprite sorting?

	// Render the background

	// Render the monster

	// Render the player
	rl.DrawCircle(int32(g.playerPosition.X), int32(g.playerPosition.Y), 10.0, rl.Red)

	// Render obstacles

	// Render the UI

	rl.EndDrawing()
}

func (g *Game) update() {
	// Update map
}

func (g *Game) handleInput() {
	// Handle touch controls
	if rl.GetTouchPointCount() > 0 {
		touchPos := rl.GetTouchPosition(0)
		if touchPos.X > g.playerPosition.X {
			g.playerPosition.X += float32(PLAYER_HORIZONTAL_SPEED)
		} else {
			g.playerPosition.X -= float32(PLAYER_HORIZONTAL_SPEED)
		}
	}

	// Handle arrow controls
	if rl.IsKeyDown(rl.KeyRight) {
		g.playerPosition.X += float32(PLAYER_HORIZONTAL_SPEED)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		g.playerPosition.X -= float32(PLAYER_HORIZONTAL_SPEED)
	}
}

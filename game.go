package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH float32 = 340
const WINDOW_HEIGHT float32 = 600

//go:embed assets/*
var ASSETS embed.FS

type Game struct {
	player       Player
	music        rl.Sound
	textureAtlas rl.Texture2D
}

func initGame() Game {
	return Game{
		player:       initPlayer(),
		music:        rl.LoadSound("assets/black_diamond.mp3"),
		textureAtlas: rl.LoadTexture("assets/tilemap_packed.png"),
	}
}

func (g *Game) unloadGame() {
	rl.UnloadSound(g.music)
	rl.UnloadTexture(g.textureAtlas)
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)

	// TODO: Sprite sorting?

	// Render the background

	// Render the monster

	// Render the player
	g.player.render()

	// Render obstacles

	// Render the UI

	rl.EndDrawing()
}

func (g *Game) update() {
	// Loop music if needed
	if !rl.IsSoundPlaying(g.music) {
		rl.PlaySound(g.music)
	}

	// Update map

	// Update player
	g.player.update()
}

func (g *Game) handleInput() {
	// Handle touch controls
	if rl.GetTouchPointCount() > 0 {
		touchPos := rl.GetTouchPosition(0)
		if touchPos.X > g.player.position.X {
			g.player.moveRight = true
		} else {
			g.player.moveLeft = true
		}
	}

	// Handle arrow controls
	if rl.IsKeyDown(rl.KeyRight) {
		g.player.moveRight = true
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		g.player.moveLeft = true
	}
}

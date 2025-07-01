package main

import (
	"embed"
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH float32 = 340
const WINDOW_HEIGHT float32 = 600
const MAX_SPEED_MULTIPLIER float32 = 2.2

//go:embed assets/*
var ASSETS embed.FS

type Game struct {
	player              Player
	obstacles           []Obstacle
	music               rl.Sound
	textureAtlas        rl.Texture2D
	score               int
	obstacleCounter     int
	difficultyCountdown int
	speedMultiplier     float32
}

func initGame() Game {
	return Game{
		player:              initPlayer(),
		obstacles:           []Obstacle{},
		music:               rl.LoadSound("assets/black_diamond.mp3"),
		textureAtlas:        rl.LoadTexture("assets/tilemap_packed.png"),
		score:               0,
		obstacleCounter:     60,
		difficultyCountdown: 600,
		speedMultiplier:     1.0,
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
	for _, o := range g.obstacles {
		o.render()
	}

	// Render the UI
	rl.DrawText(fmt.Sprintf("Score: %d", g.score), 10, 10, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Speed: %.1f", g.speedMultiplier), 10, 30, 20, rl.RayWhite)

	rl.EndDrawing()
}

func (g *Game) update() {
	// Loop music if needed - uncomment when not on wsl
	/*if !rl.IsSoundPlaying(g.music) {
		rl.PlaySound(g.music)
	}*/

	// Update map

	// Update player
	g.player.update()

	// Update obstacles
	newOList := []Obstacle{}
	for _, o := range g.obstacles {
		o.update(g.player.verticalSpeed)
		if o.position.Y >= float32(OBSTACLE_BUFFER) {
			newOList = append(newOList, o)
		}
	}
	// Determine if we should add a new obstacle
	g.obstacleCounter -= 1
	if g.obstacleCounter <= 0 {
		xPos := rand.Float32() * WINDOW_WIDTH

		// TODO: Add logic for different obstacle textures
		newOList = append(newOList, initObstacle(g.textureAtlas, rl.Vector2{X: xPos, Y: WINDOW_HEIGHT + 50}))
		g.obstacleCounter = rand.Intn(240) + 60
	}
	g.obstacles = newOList

	// Update the player speed as the game progresses to make things more difficult
	g.difficultyCountdown -= 1
	if g.difficultyCountdown <= 0 {
		g.difficultyCountdown = 600 // Ten seconds
		g.speedMultiplier = float32(math.Min(float64(g.speedMultiplier+0.2), float64(MAX_SPEED_MULTIPLIER)))
		g.player.verticalSpeed *= g.speedMultiplier
	}

	g.score += 1
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

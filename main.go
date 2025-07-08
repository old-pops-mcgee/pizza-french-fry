package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(int32(WINDOW_WIDTH), int32(WINDOW_HEIGHT), "Pizza French Fry")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)

	game := initGame()

	game.run()

	game.unloadGame()

	rl.CloseAudioDevice()
	rl.CloseWindow()
}

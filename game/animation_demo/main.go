package main

import (
	_ "image/png"
	"log"

	"github.com/Kirbstomper/animationdemo/dungeon"
	"github.com/Kirbstomper/animationdemo/maths"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenHeight = 600
	screenWidth  = 800
)

var (
	state       = 0
	runnerImage *ebiten.Image
)

var (
	player = dungeon.Player{X: 5, Y: 5}
	enemy  = dungeon.Enemy{X: 5, Y: 6}
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		player.Y = maths.Clamp(player.Y+1, 0, len(dungeon.World[0])-1)
		player.SetState(dungeon.DOWN)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		player.Y = maths.Clamp(player.Y-1, 0, len(dungeon.World[0])-1)

		player.SetState(dungeon.UP)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		player.X = maths.Clamp(player.X+1, 0, len(dungeon.World[0])-1)

		player.SetState(dungeon.RIGHT)

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		player.X = maths.Clamp(player.X-1, 0, len(dungeon.World[0])-1)

		player.SetState(dungeon.LEFT)
	}

	dungeon.World.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	dungeon.World.Draw(screen, g.count)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 640
}

func main() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	dungeon.World[5][5] = &player
	dungeon.World[5][6] = &enemy
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Animation Dungeon Demo")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

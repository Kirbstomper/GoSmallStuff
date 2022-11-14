package dungeon

import (
	"bytes"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	enemySprite *ebiten.Image
)

const (
	enemyFrameCount = 2
)

type Enemy struct {
	X, Y int
}

func init() {
	f, err := os.ReadFile("./enemy.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}
	enemySprite = ebiten.NewImageFromImage(img)

}

func (e *Enemy) getImg(c int) *ebiten.Image {
	i := (c / 20) % enemyFrameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	return enemySprite.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
}

func (e *Enemy) getXPos() int {
	return e.X
}
func (e *Enemy) getYPos() int {
	return e.Y
}

package dungeon

import (
	"bytes"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	state int
	X, Y  int
}

var (
	runnerImage *ebiten.Image
)

const (
	playerFrameNum = 4
)

func init() {
	f, err := os.ReadFile("./ethanwalk.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

}

func (p *Player) getImg(c int) *ebiten.Image {
	i := (c / 20) % playerFrameNum
	sx, sy := frameOX+i*frameWidth, frameOY+p.state
	return runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
}

func (p *Player) SetState(s int) {
	p.state = s
}
func (p *Player) getXPos() int {
	return p.X
}
func (p *Player) getYPos() int {
	return p.Y
}

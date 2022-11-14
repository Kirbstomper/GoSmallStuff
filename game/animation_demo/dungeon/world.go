package dungeon

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	worldx       = 10
	worldy       = 10
	spritewidth  = 64
	spriteheight = 64

	DOWN         = iota
	LEFT         = iota * 64
	RIGHT        = iota * 64
	UP           = iota * 64
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 0
	frameWidth  = 64
	frameHeight = 64
)

var (
	World world
)

type world [worldx][worldy]obj
type obj interface {
	getImg(c int) *ebiten.Image
	getXPos() int
	getYPos() int
}

func (w *world) Update() {
	for x := range w {
		for y, o := range w[x] {
			//Do whatever ya gotta do per turn
			if (o != nil) && ((x != o.getXPos()) || (y != o.getYPos())) {
				w[o.getXPos()][o.getYPos()] = o
				w[x][y] = nil
				log.Println("Stuff")
			}
		}
	}
}
func (w world) Draw(s *ebiten.Image, c int) {
	for x := range w {
		for y, o := range w[x] {
			if o != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*spritewidth), float64(y*spriteheight))
				s.DrawImage(o.getImg(c), op)
			}
		}
	}
}

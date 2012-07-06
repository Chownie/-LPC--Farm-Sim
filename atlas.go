package main

import (
	sdl "github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"os"
	"encoding/json"
)

type TextureMap struct {
	GameMap   [][]Tile
	Atlas     *sdl.Surface
	AtlasInfo map[string]Rect
}

type Tile struct {
	Under   Rect
	Mid     Rect
	Upper   Rect
	collide bool
}

type Rect struct {
	X int
	Y int
	H int
	W int
}

func LoadImage(name string) *sdl.Surface {
	image := sdl.Load(name)
	if image == nil {
		panic(sdl.GetError())
	}
	return image
}

func (mp *TextureMap) Init() {
	mp.Atlas = LoadImage("atlas.png")

}

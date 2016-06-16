package main

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/act"
	"engo.io/engo/common"
	"engo.io/engo/demos/demoutils"
)

type DefaultScene struct{}

var (
	scrollSpeed float32 = 700

	worldWidth  int = 800
	worldHeight int = 800
)

func (*DefaultScene) Preload() {}

// Setup is called before the main loop is started
func (*DefaultScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)
	w.AddSystem(&common.RenderSystem{})

	// Get the default axis id's
	vertAxis := engo.Axes.Id(engo.DefaultVerticalAxis)
	horiAxis := engo.Axes.Id(engo.DefaultHorizontalAxis)
	if 0 == vertAxis {
		log.Println("Default vertical axis not found, setting up fall back!")
		vertAxis = engo.Axes.SetNamed(engo.DefaultVerticalAxis, act.AxisPair{act.KeyW, act.KeyS}, act.AxisPair{act.KeyUp, act.KeyDown})
	}
	if 0 == horiAxis {
		log.Println("Default horizontal axis not found, setting up fall back!")
		horiAxis = engo.Axes.SetNamed(engo.DefaultHorizontalAxis, act.AxisPair{act.KeyA, act.KeyD}, act.AxisPair{act.KeyLeft, act.KeyRight})
	}

	// The most important line in this whole demo:
	w.AddSystem(common.NewKeyboardScroller(scrollSpeed, horiAxis, vertAxis))

	// Create the background; this way we'll see when we actually scroll
	demoutils.NewBackground(w, worldWidth, worldHeight, color.RGBA{102, 153, 0, 255}, color.RGBA{102, 173, 0, 255})
}

func (*DefaultScene) Type() string { return "Game" }

func main() {
	opts := engo.RunOptions{
		Title:          "KeyboardScroller Demo",
		Width:          worldWidth,
		Height:         worldHeight,
		StandardInputs: true,
	}

	engo.Run(opts, &DefaultScene{})
}

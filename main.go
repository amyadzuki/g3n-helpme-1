// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is a minimum G3N application showing how to create a window,
// a scene, add some 3D objects to the scene and render it.
// For more complete demos please see: https://github.com/g3n/g3nd
package main

import (
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
	"math"
	"runtime"
)

func main() {

	// Creates window and OpenGL context
	wmgr, err := window.Manager("glfw")
	if err != nil {
		panic(err)
	}
	win, err := wmgr.CreateWindow(800, 600, "Hello G3N", false)
	if err != nil {
		panic(err)
	}

	// OpenGL functions must be executed in the same thread where
	// the context was created (by CreateWindow())
	runtime.LockOSThread()

	// Create OpenGL state
	gs, err := gls.New()
	if err != nil {
		panic(err)
	}

	// Sets the OpenGL viewport size the same as the window size
	// This normally should be updated if the window is resized.
	width, height := win.Size()
	gs.Viewport(0, 0, int32(width), int32(height))

	// Creates scene for 3D objects
	scene := core.NewNode()

	// Adds white ambient light to the scene
	ambLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.5)
	scene.Add(ambLight)

	// Adds a perspective camera to the scene
	// The camera aspect ratio should be updated if the window is resized.
	aspect := float32(width) / float32(height)
	camera := camera.NewPerspective(65, aspect, 0.01, 1000)
	// CHANGE #1: instead of (0, 0, 5), use (0, -5, 0)
	camera.SetPosition(0, -5, 0)

	// CHANGE #2: make +Z the up direction
	camera.SetUp(&math32.Vector3{0, 0, 1})

	// Add an axis helper to the scene
	axis := graphic.NewAxisHelper(2)
	scene.Add(axis)

	// CHANGE #3: use a tapered cylinder instead of a sphere
	geom := geometry.NewCylinder(0.5, 1, 3, 16, 16, 0, math.Pi*2, true, false)
	mat := material.NewStandard(math32.NewColor("White"))
	mat.SetSide(material.SideDouble)
	mat.SetWireframe(true)
	cylinder := graphic.NewMesh(geom, mat)
	scene.Add(cylinder)

	// Creates a renderer and adds default shaders
	rend := renderer.NewRenderer(gs)
	err = rend.AddDefaultShaders()
	if err != nil {
		panic(err)
	}
	rend.SetScene(scene)

	// Sets window background color
	gs.ClearColor(0, 0, 0, 1.0)

	// Render loop
	for !win.ShouldClose() {

		// Rotates the cylinder a bit around the Z axis (up)
		cylinder.AddRotationY(0.005)

		// Render the scene using the specified camera
		rend.Render(camera)

		// Update window and checks for I/O events
		win.SwapBuffers()
		wmgr.PollEvents()
	}
}

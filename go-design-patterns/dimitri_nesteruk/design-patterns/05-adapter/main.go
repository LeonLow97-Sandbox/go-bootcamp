package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

// consuming an API for constructing graphical images
// Problem: we are given a VectorImage, but the DrawPoints func expects a RasterImage, solution: use an adapter

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

// given interface, using lines to construct image
func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{
		[]Line{
			{0, 0, width, 0},
			{0, 0, 0, height},
			{width, 0, width, height},
			{0, height, width, height},
		},
	}
}

// interface we have, using points to construct image
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// solution: using an adapter
type vectorToRasterAdapter struct {
	points []Point
}

var pointCache = map[[16]byte][]Point{}

// addLine converts lines to points
// func (a *vectorToRasterAdapter) addLine(line Line) {
// 	left, right := minmax(line.X1, line.X2)
// 	top, bottom := minmax(line.Y1, line.Y2)
// 	dx := right - left
// 	dy := line.Y2 - line.Y1

// 	if dx == 0 {
// 		for y := top; y <= bottom; y++ {
// 			a.points = append(a.points, Point{left, y})
// 		}
// 	} else if dy == 0 {
// 		for x := left; x <= right; x++ {
// 			a.points = append(a.points, Point{x, top})
// 		}
// 	}

// 	fmt.Println("we have", len(a.points), "points")
// }

// addLine converts lines to points with caching
func (a *vectorToRasterAdapter) addLineCached(line Line) {
	// calculating the line's hash
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}
	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		a.points = append(a.points, pts...)
		return
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	// add to cache
	pointCache[h] = a.points
	fmt.Println("we have", len(a.points), "points")
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	// magic happens here
	for _, line := range vi.Lines {
		// adapter.addLine(line)
		adapter.addLineCached(line)
	}

	return adapter // as RasterImage interface
}

func main() {
	rc := NewRectangle(6, 4)

	// using an adapter to convert VectorImage to RasterImage
	a := VectorToRaster(rc)
	_ = VectorToRaster(rc) // 2nd adapter

	fmt.Println(DrawPoints(a))
}

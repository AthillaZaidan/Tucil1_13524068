package imageprocessor

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func LoadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// cek pixel gelap apa ngga (buat detect border)
func isDarkPixel(c color.Color) bool {
	r, g, b, _ := c.RGBA()

	// convert ke 0-255
	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)

	avg := (int(r8) + int(g8) + int(b8)) / 3
	return avg < 50
}

// auto detect ukuran kotak dari jarak antar garis
func DetectCellSize(img image.Image) int {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	distances := make([]int, 0)

	// scan 3 baris horizontal buat detect border
	scanRows := [3]int{height / 4, height / 2, height * 3 / 4}

	for s := 0; s < 3; s++ {
		y := scanRows[s]

		// cari border stretches (kumpulan dark pixel yang nyambung)
		borderCenters := make([]int, 0)
		inBorder := false
		borderStart := 0

		for x := 0; x < width; x++ {
			dark := isDarkPixel(img.At(x, y))
			if dark && !inBorder {
				borderStart = x
				inBorder = true
			} else if !dark && inBorder {
				center := (borderStart + x - 1) / 2
				borderCenters = append(borderCenters, center)
				inBorder = false
			}
		}
		if inBorder {
			center := (borderStart + width - 1) / 2
			borderCenters = append(borderCenters, center)
		}

		// hitung jarak antar border center
		for i := 1; i < len(borderCenters); i++ {
			d := borderCenters[i] - borderCenters[i-1]
			if d > 10 {
				distances = append(distances, d)
			}
		}
	}

	if len(distances) == 0 {
		return width / 6
	}

	// sort distances (insertion sort)
	for i := 1; i < len(distances); i++ {
		key := distances[i]
		j := i - 1
		for j >= 0 && distances[j] > key {
			distances[j+1] = distances[j]
			j--
		}
		distances[j+1] = key
	}

	// return median
	return distances[len(distances)/2]
}

func colorsSimilar(c1, c2 color.Color, tolerance uint32) bool {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	dr := int(r1) - int(r2)
	dg := int(g1) - int(g2)
	db := int(b1) - int(b2)

	if dr < 0 {
		dr = -dr
	}
	if dg < 0 {
		dg = -dg
	}
	if db < 0 {
		db = -db
	}

	distance := uint32(dr + dg + db)
	return distance < tolerance
}

func ImageToGrid(img image.Image, cellSize int) ([][]byte, int, int) {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	// hitung berapa kotak
	cols := width / cellSize
	rows := height / cellSize

	// nyimpen warna tiap cell
	colorGrid := make([][]color.Color, rows)
	for i := 0; i < rows; i++ {
		colorGrid[i] = make([]color.Color, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			centerX := j*cellSize + cellSize/2
			centerY := i*cellSize + cellSize/2

			colorGrid[i][j] = img.At(centerX, centerY)
		}
	}

	// bikin grid hasil
	grid := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]byte, cols)
	}

	uniqueColors := make([]color.Color, 0)
	colorToLetter := make(map[color.Color]byte)

	tolerance := uint32(10000)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			currentColor := colorGrid[i][j]

			// skip border
			if isDarkPixel(currentColor) {
				grid[i][j] = '.'
				continue
			}

			// cari warna yang udah ada
			found := false
			var matchedColor color.Color

			for k := 0; k < len(uniqueColors); k++ {
				if colorsSimilar(currentColor, uniqueColors[k], tolerance) {
					found = true
					matchedColor = uniqueColors[k]
					break
				}
			}

			if found {
				grid[i][j] = colorToLetter[matchedColor]
			} else {
				// warna baru, kasih huruf baru
				newLetter := byte('A' + len(uniqueColors))
				uniqueColors = append(uniqueColors, currentColor)
				colorToLetter[currentColor] = newLetter
				grid[i][j] = newLetter
			}
		}
	}

	return grid, rows, cols
}

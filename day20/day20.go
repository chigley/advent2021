package day20

func Part1(algo []Pixel, img *Image) int {
	for i := 0; i < 2; i++ {
		img.Enhance(algo)
	}
	return img.LitPixels()
}

func Part2(algo []Pixel, img *Image) int {
	for i := 0; i < 50; i++ {
		img.Enhance(algo)
	}
	return img.LitPixels()
}

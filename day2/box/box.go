package box

type Box struct {
	Length int
	Width  int
	Height int
}

func (b Box) GetSmallestArea() int {
	sideOne := b.Length * b.Width
	sideTwo := b.Length * b.Height
	sideThree := b.Width * b.Height

	return min(sideOne, sideTwo, sideThree)
}

func (b Box) GetSquareFeet() int {
	sideOne := b.Length * b.Width * 2
	sideTwo := b.Length * b.Height * 2
	sideThree := b.Width * b.Height * 2

	return sideOne + sideTwo + sideThree
}

func (b Box) GetBowFeet() int {
	return b.Length * b.Width * b.Height

}

func (b Box) GetRibbonFeet() int {
	sideOne := b.Length + b.Width
	sideTwo := b.Length + b.Height
	sideThree := b.Width + b.Height

	return 2 * min(sideOne, sideTwo, sideThree)
}

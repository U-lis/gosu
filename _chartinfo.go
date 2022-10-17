package gosu

// Background brightness at Song select: 60% (153 / 255), confirmed.
// Score box color: Gray128 with 50% transparent
// Hovered Score box color: Gray96 with 50% transparent
func (c ChartInfo) NewChartBoard() draws.Box {
	var (
		ws = []float64{140, 360, 140}
		hs = []float64{24, 36, 32, 24}
	)
	boxs := [][]draws.Box{
		{
			{
				Inner: draws.NewLabel(c.MusicSource, Face12, color.White),
				Align: draws.AtMin,
			},
			{
				Inner: draws.NewLabel(c.Artist, Face12, color.White),
				Align: draws.ModeXY{X: draws.ModeMid, Y: draws.ModeMin},
			},
			{
				Inner: draws.NewLabel(c.NoteCountString(), Face12, color.White),
				Align: draws.ModeXY{X: draws.ModeMax, Y: draws.ModeMin},
			},
		},
		{
			{
				// Inner: draws.NewRectangle(draws.Pt(ws[0], hs[1])),
			},
			{
				Inner: draws.NewLabel(c.MusicName, Face20, color.White),
				Align: draws.ModeXY{X: draws.ModeMid, Y: draws.ModeMax},
			},
			{
				// Inner: draws.NewRectangle(draws.Pt(ws[2], hs[1])),
			},
		},
		{
			{
				Inner: draws.NewLabel(c.TimeString(), Face16, color.White),
				Align: draws.AtMin,
			},
			{
				Inner: draws.NewLabel(c.ChartName, Face16, color.White),
				Align: draws.ModeXY{X: draws.ModeMid, Y: draws.ModeMin},
			},
			{
				// Inner: draws.NewRectangle(draws.Pt(ws[2], hs[2])),
			},
		},
		{
			{
				Inner: draws.NewLabel(c.BPMString(), Face16, color.White),
				Align: draws.ModeXY{X: draws.ModeMin, Y: draws.ModeMax},
			},
			{
				Inner: draws.NewLabel(c.Charter, Face12, color.White),
				Align: draws.ModeXY{X: draws.ModeMid, Y: draws.ModeMax},
			},
			{ // Todo: ranked status
				// Inner: draws.NewRectangle(draws.Pt(ws[2], hs[3])),
				Align: draws.ModeXY{X: draws.ModeMax, Y: draws.ModeMax},
			},
		},
	}
	for i, row := range boxs {
		for j := range row {
			boxs[i][j].Outer = draws.NewRectangle(draws.Pt(ws[j], hs[i]))
			// fmt.Printf("%d %d: %+v\n", i, j, box)
		}
		// fmt.Println()
	}
	// boxs = make([][]draws.Box, 3)
	// for i := range boxs {
	// 	boxs[i] = make([]draws.Box, 3)
	// 	for j := range boxs[i] {
	// 		boxs[i][j] = draws.Box{
	// 			Inner: labels[i][j],
	// 			Align: draws.ModeXY{i, j},
	// 		}
	// 	}
	// }

	board := draws.Box{
		Inner:   draws.NewGrid(boxs, ws, hs, draws.Point{}),
		Pad:     draws.Pt(10, 10),
		Point:   draws.Pt(screenSizeX/2, 150),
		Origin2: draws.AtMid,
		Align:   draws.AtMid,
	}
	board.Outer = &draws.Rectangle{
		Size_: board.OuterSize(),
		Color: color.NRGBA{128, 128, 128, 128},
	}
	// box.Outer = draws.NewRectangle(box.OuterSize(), gray)
	// outerImage := ebiten.NewImage(box.OuterSize().XYInt())
	// outerImage.Fill(gray)
	// box.Outer = draws.NewSprite3FromImage(outerImage)
	return board
}

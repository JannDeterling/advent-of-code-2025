package day7

func FindAllPossibleBeamPaths(lines []string, startPos int) int {
	lenght := len(lines)
	lineLenght := len(lines[0])

	beamPaths := make([][]int, lenght)
	for i := range beamPaths {
		beamPaths[i] = make([]int, lineLenght)
	}

	for linePosition := 0; linePosition < lineLenght; linePosition++ {
		if lines[lenght-1][linePosition] == '|' {
			beamPaths[lenght-1][linePosition] = 1
		}else {
			beamPaths[lenght-1][linePosition] = 0
		}
	}

	for lineIndex := lenght - 2; lineIndex >= 0; lineIndex-- {
		for linePosition := 0; linePosition < lineLenght; linePosition++ {
			char := lines[lineIndex][linePosition]
			switch char {
			case '|':
				beamPaths[lineIndex][linePosition] = beamPaths[lineIndex+1][linePosition]
			case '^':
				leftPath := 0
				rightPath := 0
				if linePosition-1 >=0 {
					leftPath = beamPaths[lineIndex+1][linePosition-1]
				}
				if linePosition+1 < lineLenght {
					rightPath = beamPaths[lineIndex+1][linePosition+1]
				}
				beamPaths[lineIndex][linePosition]= leftPath + rightPath
			case '.':
				beamPaths[lineIndex][linePosition] = 0
			}
		}
	}	

	return beamPaths[0][startPos]
}

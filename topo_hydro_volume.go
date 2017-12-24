package interview

// FindMaxRainVolume determines the maximum capacity for rain that a given topological
// map can hold.
func FindMaxRainVolume(topology []int) int {

	// We need at least 3 points to create a valley.
	if len(topology) < 3 {
		return 0
	}

	var (
		// Two indexes, one for the first peak, another for the second.
		idx    = 0
		runner = 1

		totalVolume = 0
	)

	for ; runner < len(topology); runner++ {

		// A possible valley
		if topology[idx] < topology[runner] {

			// If we've encountered the end of a valley
			if idx != runner-1 {

				totalVolume += rainVolumeForValley(topology, idx, runner)
			}

			idx = runner
		}
	}

	// If there is a valley at the end of the list
	if idx != runner-1 {

		totalVolume += rainVolumeForValley(topology, idx, runner-1)
	}

	return totalVolume
}

// rainVolumeForValley finds the maximum capacity for rain for a single valley.
func rainVolumeForValley(topology []int, start, end int) int {

	minPeakHeight := topology[end]
	if topology[start] < minPeakHeight {

		minPeakHeight = topology[start]
	}

	var volume int
	for i := start; i < end; i++ {

		if topology[i] < minPeakHeight {

			volume += intAbs(topology[i] - minPeakHeight)
		}
	}

	return volume
}

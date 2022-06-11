package fstat

import (
	"errors"
	"fmt"
	"io"
	"os"

	"infowatchtest/histogram"
)

func GatherStats(filepath string) (*histogram.DiscreteHistogram[rune], error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("couldnt open file: %w", err)
	}
	defer f.Close()

	// https://stackoverflow.com/questions/236861/how-do-you-determine-the-ideal-buffer-size-when-using-fileinputstream
	buf := make([]byte, 8192)

	hist := histogram.NewDiscreteHistogram[rune]()

	for {
		n, err := f.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("couldnt read file: %w", err)
		}

		for _, char := range buf[:n] {
			hist.Add(rune(char))
		}
	}

	return hist, nil
}

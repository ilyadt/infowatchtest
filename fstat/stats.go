package fstat

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"infowatchtest/histogram"
)

func GatherStats(filepath string) (*histogram.DiscreteHistogram[byte], error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("couldnt open file: %w", err)
	}
	defer f.Close()

	// https://stackoverflow.com/questions/236861/how-do-you-determine-the-ideal-buffer-size-when-using-fileinputstream
	reader := bufio.NewReaderSize(f, 8192)

	hist := histogram.NewDiscreteHistogram[byte](256)

	for {
		b, err := reader.ReadByte()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("couldnt read file: %w", err)
		}

		hist.Add(b)
	}

	return hist, nil
}

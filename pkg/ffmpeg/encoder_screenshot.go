package ffmpeg

import "fmt"

type ScreenshotOptions struct {
	OutputPath string
	Quality    int
	Time       float64
	Width      int
	Verbosity  string
}

func (e *Encoder) Screenshot(probeResult VideoFile, options ScreenshotOptions) {
	if options.Verbosity == "" {
		options.Verbosity = "error"
	}
	if options.Quality == 0 {
		options.Quality = 1
	}
	args := []string{
		"-v", options.Verbosity,
		"-ss", fmt.Sprintf("%v", options.Time),
		"-y",
		"-i", probeResult.Path, // TODO: Wrap in quotes?
		"-vframes", "1",
		"-q:v", fmt.Sprintf("%v", options.Quality),
		"-vf", fmt.Sprintf("scale=%v:-1", options.Width),
		"-f", "image2",
		options.OutputPath,
	}
	_, _ = e.run(probeResult, args)
}

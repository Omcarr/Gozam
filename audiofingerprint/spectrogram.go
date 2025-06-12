package audiofingerprint

import (
	"errors"
	"fmt"
	"math"
	"math/cmplx"
)

const (
	dspRatio    = 4
	freqBinSize = 1024   //2^10
	maxFreq     = 5000.0 // 5kHz
	hopSize     = freqBinSize / 32
)

// remove low frequencies, downsample then perform STFT
func Spectrogram(sample []float64, sampleRate int) ([][]complex128, error) {
	filteredSample := LowPassFilter(maxFreq, float64(sampleRate), sample)

	downsampledSample, err := Downsample(filteredSample, sampleRate, sampleRate/dspRatio)
	if err != nil {
		return nil, fmt.Errorf("couldn't downsample audio sample: %v", err)
	}

	numOfWindows := len(downsampledSample) / (freqBinSize - hopSize)
	spectrogram := make([][]complex128, numOfWindows)

	window := make([]float64, freqBinSize)
	for i := range window {
		window[i] = 0.54 - 0.46*math.Cos(2*math.Pi*float64(i)/(float64(freqBinSize)-1))
	}

	// Perform STFT
	for i := 0; i < numOfWindows; i++ {
		start := i * hopSize
		end := start + freqBinSize
		if end > len(downsampledSample) {
			end = len(downsampledSample)
		}

		bin := make([]float64, freqBinSize)
		copy(bin, downsampledSample[start:end])

		// Apply Hamming window
		for j := range window {
			bin[j] *= window[j]
		}

		spectrogram[i] = FFT(bin)
	}

	return spectrogram, nil
}

// LowPassFilter is a first-order low-pass filter with
// transfer function H(s) = 1 / (1 + sRC)
func LowPassFilter(cutoffFrequency, sampleRate float64, input []float64) []float64 {
	rc := 1.0 / (2 * math.Pi * cutoffFrequency)
	dt := 1.0 / sampleRate
	alpha := dt / (rc + dt)

	filteredSignal := make([]float64, len(input))
	var prevOutput float64 = 0

	for i, x := range input {
		if i == 0 {
			filteredSignal[i] = x * alpha
		} else {

			filteredSignal[i] = alpha*x + (1-alpha)*prevOutput
		}
		prevOutput = filteredSignal[i]
	}
	return filteredSignal
}

// Downsample downsamples the input audio from originalSampleRate to targetSampleRate
func Downsample(input []float64, originalSampleRate, targetSampleRate int) ([]float64, error) {
	if targetSampleRate <= 0 || originalSampleRate <= 0 {
		return nil, errors.New("sample rates must be positive")
	}
	if targetSampleRate > originalSampleRate {
		return nil, errors.New("target sample rate must be less than or equal to original sample rate")
	}

	ratio := originalSampleRate / targetSampleRate
	if ratio <= 0 {
		return nil, errors.New("invalid ratio calculated from sample rates")
	}

	var resampled []float64
	for i := 0; i < len(input); i += ratio {
		end := i + ratio
		if end > len(input) {
			end = len(input)
		}

		sum := 0.0
		for j := i; j < end; j++ {
			sum += input[j]
		}
		avg := sum / float64(end-i)
		resampled = append(resampled, avg)
	}

	return resampled, nil
}

// to get magnitude of the spectogram
func MagnitudeSpectrogram(spec [][]complex128) ([][]float64, error) {
	if len(spec) == 0 {
		return nil, errors.New("input spectrogram is empty")
	}

	rowLen := len(spec[0])
	for _, row := range spec {
		if len(row) != rowLen {
			return nil, errors.New("inconsistent row lengths in spectrogram")
		}
	}

	magSpec := make([][]float64, len(spec))
	for i := range spec {
		magSpec[i] = make([]float64, len(spec[i]))
		for j, val := range spec[i] {
			mag := cmplx.Abs(val)
			//dB scale (20 * log10(mag))
			magSpec[i][j] = 20 * math.Log10(mag+1e-6) // avoid log(0)
		}
	}
	return magSpec, nil
}

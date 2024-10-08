# MP3 to 3CX WAV Converter

This Go program converts MP3 audio files to WAV format with specific output parameters. It's designed for use cases where lower quality audio is acceptable or preferred, such as voice recordings or applications with bandwidth constraints.

## Features

- Converts MP3 files to WAV format
- Downsamples audio to 8000 Hz
- Converts stereo to mono
- Sets bit depth to 16-bit

## Prerequisites

Before running this program, make sure you have Go installed on your system. You'll also need to install the following dependencies:

```
go get github.com/hajimehoshi/go-mp3
go get github.com/go-audio/audio
go get github.com/go-audio/wav
```

## Usage

1. Clone this repository or download the `mp3-to-wave-converter.go` file.

2. Open a terminal and navigate to the directory containing the Go file.

3. Run the program with the following command:

   ```
   go run mp3-to-wave-converter.go <input.mp3> <output.wav>
   ```

   Replace `<input.mp3>` with the path to your input MP3 file, and `<output.wav>` with the desired path for the output WAV file.

4. The program will process the file and display "Conversion completed successfully." when finished.

## Example

```
go run mp3-to-wave-converter.go input/mysong.mp3 output/mysong_converted.wav
```

## Notes

- The output WAV file will have a sample rate of 8000 Hz, 16-bit depth, and will be mono (single channel).
- This conversion will result in a loss of audio quality, which is intentional for specific use cases requiring smaller file sizes or lower bandwidth.
- Large input files may take some time to process, depending on your system's performance.

## License

[Include your chosen license here]

## Contributing

[Include guidelines for contributing to the project, if applicable]

## Acknowledgments

This program uses the following open-source libraries:
- [go-mp3](https://github.com/hajimehoshi/go-mp3)
- [go-audio](https://github.com/go-audio/audio)
- [go-audio/wav](https://github.com/go-audio/wav)
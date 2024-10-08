/*

MP3 to 3CX WAV Converter

  MIT License

  Copyright (c) `2024` `Leonid Semenenko`
  https://github.com/lsemenenko/mp3-to-3cx-wav-converter

  Permission is hereby granted, free of charge, to any person obtaining a copy
  of this software and associated documentation files (the "Software"), to deal
  in the Software without restriction, including without limitation the rights
  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  copies of the Software, and to permit persons to whom the Software is
  furnished to do so, subject to the following conditions:

  The above copyright notice and this permission notice shall be included in all
  copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
  SOFTWARE.

*/

package main

import (
    "fmt"
    "io"
    "log"
    "os"

    "github.com/hajimehoshi/go-mp3"
    "github.com/go-audio/audio"
    "github.com/go-audio/wav"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: mp3-to-3cx-wav-converter <input.mp3> <output.wav>")
        os.Exit(1)
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    // Open the MP3 file
    file, err := os.Open(inputFile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Decode the MP3
    decoder, err := mp3.NewDecoder(file)
    if err != nil {
        log.Fatal(err)
    }

    // Set output parameters
    outputSampleRate := 8000
    outputNumChannels := 1 // mono

    // Create the output WAV file
    out, err := os.Create(outputFile)
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    // Create a new WAV encoder with the desired output format
    enc := wav.NewEncoder(out, outputSampleRate, 16, outputNumChannels, 1)
    defer enc.Close()

    // Create an audio buffer
    audioBuf := make([]int, 8192)
    buf := &audio.IntBuffer{
        Data: audioBuf,
        Format: &audio.Format{
            NumChannels: outputNumChannels,
            SampleRate:  outputSampleRate,
        },
    }

    // Read and convert
    downsampleRatio := float64(decoder.SampleRate()) / float64(outputSampleRate)
    sampleSum := 0
    sampleCount := 0
    outputIndex := 0

    for {
        frame := make([]byte, 8192)
        n, err := decoder.Read(frame)
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        // Process samples
        for i := 0; i < n-3; i += 4 {
            // Convert bytes to int (left and right channel)
            left := int(int16(frame[i]) | int16(frame[i+1])<<8)
            right := int(int16(frame[i+2]) | int16(frame[i+3])<<8)

            // Mix channels to mono and accumulate
            sampleSum += (left + right) / 2
            sampleCount++

            // Check if we have accumulated enough samples for downsampling
            if float64(sampleCount) >= downsampleRatio {
                // Average the accumulated samples
                avgSample := sampleSum / sampleCount

                // Store the downsampled value
                audioBuf[outputIndex] = avgSample
                outputIndex++

                // Reset accumulators
                sampleSum = 0
                sampleCount = 0
            }
        }

        // Write the downsampled data
        buf.Data = audioBuf[:outputIndex]
        if err := enc.Write(buf); err != nil {
            log.Fatal(err)
        }

        // Reset output index
        outputIndex = 0
    }

    fmt.Println("Conversion completed successfully.")
}
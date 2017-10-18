package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gordonklaus/portaudio"
)

type deviceInfo portaudio.DeviceInfo

func (d *deviceInfo) Set(s string) error {
	devices, err := portaudio.Devices()
	if err != nil {
		log.Fatal(err)
	}
	for _, device := range devices {
		if device.Name != s {
			continue
		}
		*d = deviceInfo(*device)
		return nil
	}
	return fmt.Errorf("could not find device %v", s)
}

func (d deviceInfo) String() string {
	return d.Name
}

func main() {
	if err := portaudio.Initialize(); err != nil {
		log.Fatal(err)
	}
	defer portaudio.Terminate()

	var in, out deviceInfo
	flag.Var(&in, "in", "Input device")
	flag.Var(&out, "out", "Input device")
	sampleRate := flag.Float64("samplerate", 44100, "Sample Rate")
	bufferSize := flag.Int("buffersize", 8092, "Frames per Buffer")
	list := flag.Bool("devices", false, "List devices")
	flag.Parse()

	devices, err := portaudio.Devices()
	if err != nil {
		log.Fatal(err)
	}
	if *list {
		for _, d := range devices {
			fmt.Println(d.Name)
		}
		return
	}

	ina := portaudio.DeviceInfo(in)
	outa := portaudio.DeviceInfo(out)
	p := portaudio.HighLatencyParameters(&ina, &outa)
	p.Input.Channels = 1
	p.Output.Channels = 1
	p.SampleRate = *sampleRate
	p.FramesPerBuffer = *bufferSize

	inBuffer := make([]int32, *bufferSize)
	outBuffer := make([]int32, *bufferSize)
	stream, err := portaudio.OpenStream(p, &inBuffer, &outBuffer)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := stream.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}

	for {
		if err := stream.Read(); err != nil {
			log.Fatal(err)
		}
		outBuffer = inBuffer
		if err := stream.Write(); err != nil {
			log.Fatal(err)
		}
	}
}

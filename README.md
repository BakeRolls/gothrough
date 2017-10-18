# gothrough

Pass audio from any input to any output. I needed something to send audio from one computer to another.

```bash
$ # install
$ go get github.com/BakeRolls/gothrough
```

```bash
$ gothrough -help
Usage of gothrough:
  -buffersize int
    Frames per Buffer (default 8092)
  -devices
    List devices
  -in value
    Input device
  -out value
    Input device
  -samplerate float
    Sample Rate (default 44100)
```

```bash
$ # list devices
$ gothrough -devices
Built-in Microphone
Built-in Output
Sound Blaster Tactic(3D) Alpha
BoomAudio
```

```bash
$ # pass input from "Sound Blaster" to the Built-in Output:
$ gothrough -in "Sound Blaster Tactic(3D) Alpha" -out "Built-in Output"
$ # If a device has an in- and an output, you can pass it to itself:
$ gothrough -in "Sound Blaster Tactic(3D) Alpha" -out "Sound Blaster Tactic(3D) Alpha"
```

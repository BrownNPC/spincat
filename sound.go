package main

import (
	"bytes"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/wav"

	_ "embed"
)

//go:embed spin.wav
var SpinWav []byte

type Audio struct {
	stream beep.StreamSeekCloser
	ctrl   *beep.Ctrl
	mute   bool
}

func (a *Audio) SetMute(v bool) {
	if a.mute == v {
		return
	}
	if v {
		a.mute = true
		a.Stop()
	} else {
		a.mute = false
	}
}
func (a *Audio) PlayLoop() {
	if a.mute {
		return
	}
	speaker.Lock()
	a.ctrl.Paused = false
	speaker.Unlock()
}
func (a *Audio) Stop() {
	if a.mute {
		return
	}
	speaker.Lock()
	a.stream.Seek(0)
	a.ctrl.Paused = true
	speaker.Unlock()
}

func LoadAudio() Audio {
	stream, format, err := wav.Decode(bytes.NewReader(SpinWav))
	if err != nil {
		panic("failed to decode sound: " + err.Error())
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Millisecond*100))

	loop, _ := beep.Loop2(stream)
	ctrl := &beep.Ctrl{
		Streamer: loop,
		Paused:   true,
	}
	speaker.Play(ctrl)
	return Audio{
		stream: stream,
		ctrl:   ctrl,
	}
}

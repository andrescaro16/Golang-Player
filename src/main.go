package main

import (
	"os"
	"time"
	"sync"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

func player(wg *sync.WaitGroup, otoCtx *oto.Context, file string) {
	defer wg.Done()

	fileBytes, err := os.Open(file)
	if err != nil {
		panic("reading " + file + " failed: " + err.Error())
	}

	decodedMp3, err := mp3.NewDecoder(fileBytes)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	player := otoCtx.NewPlayer(decodedMp3)

	player.Play()

	for player.IsPlaying() {
	    time.Sleep(time.Millisecond)
	}

	err = player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}

	fileBytes.Close()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	samplingRate := 48000
	numOfChannels := 2
	audioBitDepth := 2
	otoCtx, readyChan, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	<-readyChan

	file1 := "../audios/NF - LOST ft. Hopsin.mp3"
	file2 := "../audios/NF - HOPE.mp3"
	go player(&wg, otoCtx, file1)
	go player(&wg, otoCtx, file2)
	wg.Wait()
}

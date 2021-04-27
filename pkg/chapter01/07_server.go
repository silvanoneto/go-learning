// Package chapter01 07_server contains the code studied on section 07
package chapter01

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Server1 is a simple server that gives back the URL path
func Server1() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Server2 is a simple server and counter that gives back the URL path
func Server2() {
	var mu sync.Mutex
	var count int

	handler := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}
	counter := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count %d\n", count)
		mu.Unlock()
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Server3 echoes an HTTP request
func Server3() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Exercise12 is a lissajous server
func Exercise12() {
	rand.Seed(time.Now().UTC().UnixNano())

	type lissajousConfig struct {
		Palette []color.Color
		Cycles  int
		Res     float64
		Size    int
		NFrames int
		Delay   int
	}

	newLissajousConfig := func() lissajousConfig {
		return lissajousConfig{
			Palette: []color.Color{
				color.White,
				color.RGBA{0xff, 0x00, 0x00, 0xff},
				color.RGBA{0x00, 0xff, 0x00, 0xff},
				color.RGBA{0x00, 0x00, 0xff, 0xff},
			},
			Cycles:  5,
			Res:     0.001,
			Size:    100,
			NFrames: 64,
			Delay:   8,
		}
	}

	lissajousAlternative2 := func(out io.Writer, config lissajousConfig) {
		freq := rand.Float64() * 3.0
		anim := gif.GIF{LoopCount: config.NFrames}
		phase := 0.0
		for i := 0; i < config.NFrames; i++ {
			rect := image.Rect(0, 0, 2*config.Size+1, 2*config.Size+1)
			img := image.NewPaletted(rect, config.Palette)
			max := float64(config.Cycles) * 2 * math.Pi
			for t := 0.0; t < max; t += config.Res {
				x := math.Sin(t)
				y := math.Sin(t*freq + phase)
				img.SetColorIndex(config.Size+int(x*float64(config.Size)+0.5),
					config.Size+int(y*float64(config.Size)+0.5),
					uint8(rand.Intn(len(config.Palette)-1)+1))
			}
			phase += 0.1
			anim.Delay = append(anim.Delay, config.Delay)
			anim.Image = append(anim.Image, img)
		}
		gif.EncodeAll(out, &anim)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}

		getValuesFromRequest := func(paramName string, w http.ResponseWriter,
			r *http.Request) (int, error) {

			param := r.URL.Query().Get(paramName)
			if param == "" {
				errorMsg := paramName + " param has no value"
				return 0, errors.New(errorMsg)
			}
			value, err := strconv.Atoi(param)
			if err != nil {
				errorMsg := paramName + " param value is invalid."
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w)
				return 0, errors.New(errorMsg)
			}
			return value, nil
		}

		config := newLissajousConfig()
		if cycles, err := getValuesFromRequest("cycles", w, r); err == nil {
			config.Cycles = cycles
		}
		if size, err := getValuesFromRequest("size", w, r); err == nil {
			config.Size = size
		}
		if nframes, err := getValuesFromRequest("nframes", w, r); err == nil {
			config.NFrames = nframes
		}
		if delay, err := getValuesFromRequest("delay", w, r); err == nil {
			config.Delay = delay
		}
		lissajousAlternative2(w, config)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

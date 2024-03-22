package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func (w *Weather) checkStatus() (resWater string, resWind string) {
	switch {
	case w.Water < 5:
		resWater = "aman"
	case w.Water >= 6 && w.Water <= 8:
		resWater = "siaga"
	case w.Water > 8:
		resWater = "bahaya"
	}

	switch {
	case w.Wind < 6:
		resWind = "Aman"
	case w.Wind >= 7 && w.Wind <= 15:
		resWind = "Siaga"
	case w.Wind > 15:
		resWind = "Bahaya"
	}
	return resWater, resWind
}

func generateJSON(weather Weather) {

	res, err := json.Marshal(weather)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%s\n", res)
	resWater, resWind := weather.checkStatus()
	fmt.Printf("status water : %s\n", resWater)
	fmt.Printf("status wind : %s\n", resWind)
}

func main() {
	weather := Weather{}

	for {
		time.Sleep(5 * time.Second)

		weather.Water = rand.Intn(20)
		weather.Wind = rand.Intn(20)
		generateJSON(weather)
	}
}

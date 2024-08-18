package main

import "fmt"

var ErrInvalidStationNumber = fmt.Errorf("invalid weather station number")

func setTemp(stationsList *[10]float64, station int, temp float64) error {
	if station < 0 || station >= len(*stationsList) {
		return ErrInvalidStationNumber
	}
	(*stationsList)[station] = temp
	return nil
}

func avgTemp(stations [10]float64) float64 {
	temp := 0.0
	for _, t := range stations {
		temp += t
	}
	return temp / float64(len(stations))
}

func printTemperature(stations [10]float64) {
	for i, t := range stations {
		fmt.Printf("st: %d - %.1f °C\n", i, t)
	}
}

func main() {
	var weatherStations [10]float64

	var (
		ws int
		t  float64
	)

	for {
		avgTemp := avgTemp(weatherStations)
		printTemperature(weatherStations)
		fmt.Printf("Avg temp: %.1f°C\n", avgTemp)

		fmt.Printf("\nEnter weather station number [0-9] (or -1 to exit): ")
		_, err := fmt.Scanln(&ws)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if ws == -1 {
			break
		}

		fmt.Printf("Enter temperature: ")
		_, err = fmt.Scanln(&t)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = setTemp(&weatherStations, ws, t)
		if err != nil {
			fmt.Println(err)
		}
	}
}

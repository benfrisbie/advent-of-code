package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

type Almanac struct {
	Seeds                 []int
	SeedToSoil            RangeMapping
	SoilToFertilizer      RangeMapping
	FertilizerToWater     RangeMapping
	WaterToLight          RangeMapping
	LightToTemperature    RangeMapping
	TemperatureToHumidity RangeMapping
	HumidityToLocation    RangeMapping
}

type RangeMapping struct {
	DestinationRangeStart []int
	SourceRangeStart      []int
	RangeLength           []int
}

func (rm RangeMapping) Map(source int) int {
	for i := range rm.SourceRangeStart {
		if source >= rm.SourceRangeStart[i] && source < rm.SourceRangeStart[i]+rm.RangeLength[i] {
			return rm.DestinationRangeStart[i] + (source - rm.SourceRangeStart[i])
		}
	}
	return source
}

func (rm *RangeMapping) AddRange(s string) {
	split := strings.Split(s, " ")
	if len(split) != 3 {
		fmt.Println(s, split)
		panic("invalid input")
	}
	destRangeStart, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	sourceRangeStart, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	rangeLength, err := strconv.Atoi(split[2])
	if err != nil {
		panic(err)
	}
	rm.DestinationRangeStart = append(rm.DestinationRangeStart, destRangeStart)
	rm.SourceRangeStart = append(rm.SourceRangeStart, sourceRangeStart)
	rm.RangeLength = append(rm.RangeLength, rangeLength)
}

func ReadInput(inputPath string) Almanac {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := Almanac{
		SeedToSoil:            RangeMapping{},
		SoilToFertilizer:      RangeMapping{},
		FertilizerToWater:     RangeMapping{},
		WaterToLight:          RangeMapping{},
		LightToTemperature:    RangeMapping{},
		TemperatureToHumidity: RangeMapping{},
		HumidityToLocation:    RangeMapping{},
	}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		if s == "" {
			i++
			scanner.Scan()
			continue
		}

		if i == 0 {
			for _, seedString := range strings.Split(s, " ")[1:] {
				seed, err := strconv.Atoi(seedString)
				if err != nil {
					panic(err)
				}
				data.Seeds = append(data.Seeds, seed)
			}
		} else if i == 1 {
			data.SeedToSoil.AddRange(s)
		} else if i == 2 {
			data.SoilToFertilizer.AddRange(s)
		} else if i == 3 {
			data.FertilizerToWater.AddRange(s)
		} else if i == 4 {
			data.WaterToLight.AddRange(s)
		} else if i == 5 {
			data.LightToTemperature.AddRange(s)
		} else if i == 6 {
			data.TemperatureToHumidity.AddRange(s)
		} else if i == 7 {
			data.HumidityToLocation.AddRange(s)
		}
	}

	return data
}

func Part1(data Almanac) int {
	min := -1

	for _, seed := range data.Seeds {
		soil := data.SeedToSoil.Map(seed)
		fertilizer := data.SoilToFertilizer.Map(soil)
		water := data.FertilizerToWater.Map(fertilizer)
		light := data.WaterToLight.Map(water)
		temperature := data.LightToTemperature.Map(light)
		humidity := data.TemperatureToHumidity.Map(temperature)
		location := data.HumidityToLocation.Map(humidity)

		if min == -1 || location < min {
			min = location
		}
	}

	return min
}

func Part2(data Almanac) int {
	// TODO: optimize?
	min := -1
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(data.Seeds) / 2)

	for i := 0; i < len(data.Seeds); i += 2 {
		go func(seedStart int, seedEnd int) {
			for seed := seedStart; seed <= seedEnd; seed++ {
				soil := data.SeedToSoil.Map(seed)
				fertilizer := data.SoilToFertilizer.Map(soil)
				water := data.FertilizerToWater.Map(fertilizer)
				light := data.WaterToLight.Map(water)
				temperature := data.LightToTemperature.Map(light)
				humidity := data.TemperatureToHumidity.Map(temperature)
				location := data.HumidityToLocation.Map(humidity)

				mutex.Lock()
				if min == -1 || location < min {
					min = location
				}
				mutex.Unlock()
			}
			wg.Done()
		}(data.Seeds[i], data.Seeds[i]+data.Seeds[i+1])
	}

	wg.Wait()

	return min
}

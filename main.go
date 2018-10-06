package main

import (
	"ApproxPI/tables"
	"math"
	"math/rand"
	"strconv"
)

const xCenter float64 = 0.5
const yCenter float64 = 0.5
const radius float64 = 0.5
const repetitions int = 1000000

type Point struct {
	x float64
	y float64
}

func main() {
	var point Point
	count := 0
	hit := 0
	var track map[int]float64
	track = make(map[int]float64)
	for i:=1 ; i <= repetitions ; i++ {
		point = getRandomPoint()
		count ++;
		if(isInsideTheCircle(point)){
			hit ++;
		}
		if i% 10000 == 0{
			track[i] = float64(hit)/float64(count)/(math.Pow(radius,2.0))
		}
	}

	table := tables.New([]string{"Iteration", "Evaluation"})

	for key,val := range track{
		iterString := strconv.Itoa(key)
		evalString := strconv.FormatFloat(val, 'f', -1, 64)
		table.AddRow(map[string]interface{}{"Iteration": iterString , "Evaluation": evalString })
	}

	table.Print()
	
}

func getRandFloat() float64{
	return rand.Float64()
}

func getRandomPoint() Point{
	return Point{x:getRandFloat(),y:getRandFloat()}
}

func isInsideTheCircle(p Point) bool{
	return math.Pow(p.x-xCenter,2.0) + math.Pow(p.y-yCenter,2.0) < math.Pow(radius,2)
}
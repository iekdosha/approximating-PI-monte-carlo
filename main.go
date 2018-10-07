package main

import (
	"ApproxPI/tables"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const xCenter float64 = 0.5
const yCenter float64 = 0.5
const radius float64 = 0.5
const repetitions int = 1000000
const trackInterval int = 10000


type Point struct {
	x float64
	y float64
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var point Point
	hit := 0
	miss := 0
	var track []float64
	// create random points
	for i:=1 ; i <= repetitions ; i++ {
		//generate a random point
		point = getRandomPoint()
		// count hit or miss points
		if(isInsideTheCircle(point)){
			hit ++
		} else{
			miss ++
		}
		// enter a result to the track in preset intervals
		if i%trackInterval == 0{
			track = append(track,calcHitRatio(hit,miss)/(math.Pow(radius,2.0)))
		}
	}

	// create a table
	table := tables.New([]string{"Iteration", "Evaluation"})
	// insert rows into table
	for i,val := range track{
		iterString := strconv.Itoa((i+1)* trackInterval)
		evalString := strconv.FormatFloat(val, 'f', -1, 64)
		table.AddRow(map[string]interface{}{"Iteration": iterString , "Evaluation": evalString })
	}
	// print table
	table.Print()

}

func calcHitRatio(hit int , miss int) float64{
	return float64(hit)/float64(hit + miss)
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
package main

import (
  "fmt"
  "io/ioutil"
  "time"
  "strconv"
  "bytes"
)

func main() {

  fmt.Println("Day 2")
	start := time.Now()
  
	dat, err := ioutil.ReadFile("day2.txt")
	if err != nil {
		panic(err)
  }
  dat = dat[:len(dat) - 1]
    
  lines := bytes.Split(dat, []byte("\n"))

  vertical := 0
  horizontal := 0

  part2vertical := 0
  part2horizontal := 0
  aim := 0

  for _, value := range lines {
    
    splitInstruction := bytes.Split(value, []byte(" "))

    direction := string(splitInstruction[0])
    scale, _ := strconv.Atoi(string(splitInstruction[1]))
    
    if direction == "forward" {
      horizontal += scale
      part2horizontal += scale
      part2vertical += aim * scale
    } else if direction == "up" {
      vertical -= scale
      aim -= scale
    } else if direction == "down" {
      vertical += scale
      aim += scale
    }

  }

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %v \n", elapsed)
  fmt.Printf("Part 1: %v \n", vertical * horizontal) 
  fmt.Printf("Part 2: %v \n", part2vertical * part2horizontal) 

}

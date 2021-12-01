package main

import (
  "fmt"
  "io/ioutil"
  "time"
  "strconv"
  "bytes"
)

func main() {
  
  fmt.Println("Day 1")
	start := time.Now()
  
	dat, err := ioutil.ReadFile("day1.txt")
	if err != nil {
		panic(err)
	}
  
  lines := bytes.Split(dat, []byte("\n"))

  increase := 0
  increaseWindow := 0

  for i := 0; i < len(lines) - 1; i++ {
    
    num1, _ := strconv.Atoi(string(lines[i]))
    num2, _ := strconv.Atoi(string(lines[i + 1]))

    if num1 < num2 {
      increase++
    }

  }

  for i := 0; i < len(lines) - 3; i++ {
    
    num1, _ := strconv.Atoi(string(lines[i]))
    num2, _ := strconv.Atoi(string(lines[i + 1]))
    num3, _ := strconv.Atoi(string(lines[i + 2]))
    num4, _ := strconv.Atoi(string(lines[i + 3]))

    if num1 + num2 + num3 < num2 + num3 + num4 {
      increaseWindow++
    }

  }

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %v \n", elapsed)
  fmt.Printf("Part 1: %v \n", increase) 
  fmt.Printf("Part 2: %v \n", increaseWindow) 

}

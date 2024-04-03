package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "math/rand/v2"
    "time"
)

const length = 80
const height  = 30
const startInit = 800

type liveField [height][length]int

func clearOutput() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  err := cmd.Run()

  if err != nil {
      log.Fatal(err)
  }
}

func printField(field *liveField){
  for _, line := range field {
    // fmt.Print(i)
    for _, column := range line {
      if column == 1 {
        fmt.Print("*")
      } else {
        fmt.Print(".")
      }

    }
    fmt.Println()
  }
}



func randomInit(field *liveField){
  for i := 0; i < startInit; i++ {
    field[rand.IntN(height)][rand.IntN(length)] = 1
  }
}

func calculateNewGender(oldGender *liveField, newGender *liveField){
  for i, v := range newGender {
    for j,_ := range v {
      newGender[i][j] = newValueForField(oldGender, i, j)
    }
  }
}

func newValueForField(oldGender *liveField, i int, j int) int {

  var summ int

  upI := i - 1
  centrI := i
  downI := i + 1

  leftJ := j - 1
  centrJ := j
  rightJ := j + 1

  if upI < 0 {
    upI = height - 1
  }
  if downI == height {
    downI = 0
  }

  if leftJ < 0 {
    leftJ = length - 1
  }

  if rightJ == length {
    rightJ = 0
  }

  summ =
    oldGender[upI][leftJ] + oldGender[upI][centrJ] + oldGender[upI][rightJ] +
    oldGender[centrI][leftJ] + oldGender[centrI][rightJ] +
    oldGender[downI][leftJ] + oldGender[downI][centrJ] + oldGender[downI][rightJ]

  if oldGender[i][j] == 1 {
    if (summ == 2 || summ == 3) {
      return 1
    } else {
      return 0
    }
  } else if oldGender[i][j] == 0 {
    if summ == 3 {
      return 1
    } else {
      return 0
    }
  } else {
    fmt.Println("Ошибочка вышла")
    return 412
  }
}

func main() {
  var field liveField
  var newGender liveField

  randomInit(&field)

  for true {
    clearOutput()
    calculateNewGender(&field, &newGender)
    field = newGender

    printField(&field)

    time.Sleep(100 * time.Millisecond)
  }
}

package main

import (
  "fmt"
  "sort"
)

func min(a, b int) int {
  if a <= b {
    return a
  }
  return b
}

func max(a, b int) int {
  if a>= b {
    return a
  }
  return b
}

func sum(array []int) int {
  var result = 0
  for _, t := range array {
    result += t
  }
  return result
}


func main()  {
  var arrNumber = []int{1,2,33,44,55,33,44,22,44,33,2,77,66,1,2,3,4,5,6,7,89,3,3,8,9,75,4,3,2,2,1,3}
  var limit = 11

  for i := 0; i < len(arrNumber); i += limit {
    var groupNumber = arrNumber[i:min(i+limit, len(arrNumber))]
    sort.Sort(sort.Reverse(sort.IntSlice(groupNumber)))
    total := sum(groupNumber)
    avg := (float32(total)/float32(len(groupNumber)))

    minVal := groupNumber[0]
    maxVal := groupNumber[0]

    for _, value := range groupNumber {
      if value < minVal {
        minVal = value
      }
      if value > maxVal {
        maxVal = value
      }
    }

    fmt.Println()
    fmt.Println(groupNumber)
    fmt.Println("Total jumlah =",total)
    fmt.Println("Rata-rata =", avg)
    fmt.Println("Nilai minimal =",minVal," Nilai maksimal=",maxVal)
  }


}

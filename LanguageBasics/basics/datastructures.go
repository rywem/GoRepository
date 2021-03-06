package basics
import "fmt"
//array
//slices
//map
//struct

func ArrayExample() {
  var x [58] string
  fmt.Println(x)
  fmt.Println(len(x))
  fmt.Println(x[42])

  for i := 65; i <= 122; i++ {
    x[i-65] = string(i)
  }
  fmt.Println(x)
  fmt.Println(len(x))
  fmt.Println(x[42])
}

func SliceExample() {
  mySlice := make([]int, 0, 5)
  fmt.Println("-----")
  fmt.Println(mySlice)
  fmt.Println(len(mySlice))
  fmt.Println(cap(mySlice))

  for i := 0; i < 80; i++ {
    mySlice = append(mySlice, i)
    fmt.Println("Len: ", len(mySlice), "Capacity: ", cap(mySlice), "Value: " , mySlice[i])
  }


}

func MapExample() {
  m := make(map[string] int)
  m["k1"] = 7
  m["k2"] = 12

  fmt.Println("map: ", m)
  v1 := m["k1"]
  fmt.Println("v1: ", v1)
  fmt.Println("len: ", len(m))
  delete(m, "k2")
  fmt.Println("map: ", m)
  _, prs := m["k2"]
  fmt.Println("prs: ", prs)
}

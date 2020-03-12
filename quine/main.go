package main
import `fmt`
func main() {
  q := string(34)
  l := []string {
"package main",
"import `fmt`",
"func main() {",
"  q := string(34)",
"  l := []string {",
"  }",
"  for i := 0; i < 5; i++ {",
"    fmt.Println(l[i])",
"  }",
"  for i := 0; i < len(l); i++ {",
"    fmt.Println(q + l[i] + q + string(','))",
"  }",
"  for i := 5; i < len(l); i++ {",
"    fmt.Println(l[i])",
"  }",
"}",
  }
  for i := 0; i < 5; i++ {
    fmt.Println(l[i])
  }
  for i := 0; i < len(l); i++ {
    fmt.Println(q + l[i] + q + string(','))
  }
  for i := 5; i < len(l); i++ {
    fmt.Println(l[i])
  }
}
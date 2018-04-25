package main

import (
	"bufio"
	"os"
	"fmt"
)

//func main() {
//	s, sep := "", ""
//	for _, arg := range os.Args {
//		s += sep + arg
//		sep = ","
//	}
//	fmt.Println("Fuck you!!")
//	fmt.Println(strings.Join(os.Args, ";"))
//	fmt.Println(s)
//	fmt.Println(os.Args[0:])
//}

var commands = map[string]func(){
	"dup": dup,
}

func main() {
	command := os.Args[1]
	commands[command]()
}

func animateGif() {

}

func dup() {
	counts := make(map[string]int)
	files := os.Args[2:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

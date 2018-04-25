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

var commands map[string]func()

func main() {
	commands = map[string]func(){
		"dup":      dup,
		"lissajou": lissajou,
		"help":     showHelp,
		"fetch":    fetch,
		"server":   server,
	}

	command := os.Args[1]
	commands[command]()
}

func showHelp() {
	fmt.Println("Experimental application.")
	fmt.Println("Usage: awesomProject [command] (params...)")
	fmt.Println(fmt.Sprintf("Awailable commands: %v", mapKeys(commands)))
}

func mapKeys(m map[string]func()) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
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

package application

import (
	"go.uber.org/zap"
	"awesomeProject/server"
	"awesomeProject/queue"
	"gopkg.in/cheggaaa/pb.v1"
	"time"
	"os"
	"fmt"
	"bufio"
)

type App struct {
	zaps     *zap.SugaredLogger
	commands map[string]func()
}

func (a App) ShowHelp() {
	fmt.Println("Experimental application.")
	fmt.Println("Usage: awesomProject [command] (params...)")
	fmt.Println(fmt.Sprintf("Awailable commands: %v", mapKeys(a.commands)))
}

// Process command by it's name
func (a App) Process(command string) {
	a.commands[command]()
}

// Creating new application. Just some test application.
func New() *App {
	a := new(App)
	a.zaps = zap.NewExample().Sugar()
	defer a.zaps.Sync()

	a.zaps.Infow("Application initalized",
		"in", "10 s")

	a.commands = map[string]func(){
		"dup":        dup,
		"lissajou":   lissajou,
		"help":       a.ShowHelp,
		"fetch":      fetch,
		"server":     server.Server,
		"pbar":       pbar,
		"send:queue": queue.SendQueue,
		"read:queue": queue.ReadQueue,
	}

	return a
}

func pbar() {
	count := 100000
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("The End!")
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

func mapKeys(m map[string]func()) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

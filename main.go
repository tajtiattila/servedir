package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	nobr := flag.Bool("n", false, "don't start browser")
	flag.Parse()

	if flag.NArg() > 1 {
		log.Fatal("at most path argument is needed")
	}

	dir := "."
	if flag.NArg() == 1 {
		dir = flag.Arg(0)
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))

	if !*nobr {
		go func() {
			time.Sleep(time.Second)
			host, port, err := net.SplitHostPort(*addr)
			if err != nil {
				log.Printf("couldn't parse %q: %v\n", *addr, err)
				return
			}
			if host == "" {
				host = "localhost"
			}
			err = openbrowser(fmt.Sprintf("http://%s:%s", host, port))
			if err != nil {
				log.Println("couldn't start browser:", err)
			}
		}()
	}

	fmt.Println("listening on", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func openbrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}

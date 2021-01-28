package aliver

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const(
	windowWidth = 1920
	windowHeight = 1080
	banner = `
                                                
                                                                                                    
                                               ++ossyh+                                             
                                    +oosyyhdmNNNMMMMMMo                                             
                        ++ossyhddmNNNMMMMMMMMMMMMMMMMMo                                             
                      oNNNMMMMMMMMMMMMMMMMMMMMMMMMMMMMo                                             
                      oMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMo                                             
                      oMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMo hNNNmho                                     
                      oMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMo NMMMMMNs     oyhhyo                         
                      oMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMo NMMMMMMd   omMMMMMNd+                       
                      oMMMMMMMMmmmdddddhhhyyyhMMMMMMMMo NMMMMMNo   NMMMMMMMMh                       
                      oMMMMMMMM+             oMMMMMMMMo hmmmdy+    mMMMMMMMMy                       
                      oMMMMMMMMoooss    syyyydMMMMMMMMo            +dNMMMMmy                        
                      oMMMMMMMMMMMMM+   NMMMMMMMMMMMMMo yyyyyyyyy    +osoo                          
                      oMMMMMMMMMMMMM+   NMMMMMMMMMMMMMo NMMMMMMMM+ ++++++++++                       
                      oMMMMMMMMMMMMM+   NMMMMMMMMMMMMMo NMMMMMMMM+ NNNNNNNNNh                       
                      oMMMMMMMMMMMMM+   NMMMMMMMMMMMMMo NMMMMMMMM+ MMMMMMMMMd                       
                      oMMMMMMMMMMMMM+   NMMMMMMMMMMMMMo NMMMMMMMM+ MMMMMMMMMd                       
                      oMMMMMMMMMMMMM+   NMMMMMMMMMMMMMo NMMMMMMMM+ MMMMMMMMMd                       
                      oMMMMMMMMMMMMM+   NMMMMMMMMMMMMMo NMMMMMMMM+ MMMMMMMMMd                       
                      oMMMMMMMMMMMMMhyyyNMMMMMMMMMMMMMo NMMMMMMMM+ MMMMMMMMMs                       
                      oMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMo NMMMMMMMM+ MMMMMNNdo                        
                      oMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMo NMMMMMMMN  ooooo+                           
                      oMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMo NMMMMMNmo                                   
                      oMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMo syyyys+                                     
                      oMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMo                                             
                      +hddmNNNMMMMMMMMMMMMMMMMMMMMMMMMo                                             
                             +oosyyhdmNNNMMMMMMMMMMMMMo                                             
                                        ++ossyhdmmNNMMo                                             
                                                    +o                                              
                                                                                                    
                             Microsoft Teams
                             Presence Bypassing Utility v0.1

                             Supported systems:
                             -Linux only with xdotool module
`
)
// it requires
// sudo apt-get install xdotool -y
func Start() error {
	fmt.Println(banner)
	// add support to detect Ctrl+C interrupts
	go detectInterrupt()
	// start mouse fake clicker routine
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)
	for {
		select {
		case <-done:
		case t := <-ticker.C:
			fmt.Println("Tick at", t.Format("02 Jan 06 15:04 -0700"))
			x, y := randomBetween(0, windowWidth), randomBetween(0, windowHeight)
			if err := move(x, y); err != nil {
				// todo send error to a channel
			}
		}
	}
}

// detectInterrupt creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func detectInterrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func leftClick(x, y int) error {
	// example command to move the mouse
	// xdotool click 1 mousemove_relative 0 100
	query := fmt.Sprintf("mousemove %d %d click 1", x, y)
	cmd := exec.Command("xdotool", strings.Split(query, " ")...)
	return cmd.Run()
}

func move(x, y int) error {
	// example command to move the mouse
	// xdotool click 1 mousemove_relative 0 100
	query := fmt.Sprintf("mousemove %d %d", x, y)
	cmd := exec.Command("xdotool", strings.Split(query, " ")...)
	return cmd.Run()
}

// randomBetween creates a random number between min nad max values
func randomBetween(min, max int) int {
	return rand.Intn(max - min) + min
}
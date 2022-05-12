package main

import (
	"fmt"
	"os"
	"os/user"
)

const art = `
⠀⠀⠀⠀⠀⣠⣴⣖⢶⣶⣦⣀⢰⣿⣿⣿⣿⠏⠀⠀⣐⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣸⣿⣿⣿⣯⣻⣿⣿⢹⣿⣿⣿⡟⠠⠨⠪⠂⣰⣦⣄⠀⠀⠀ 	%s
⠀⠀⠀⢠⢯⣶⣮⢿⣿⣷⡽⣿⡘⣿⡿⣿⠁⡑⠡⠃⣰⣿⣿⣿⣦⠀⠀	  %s
⠀⠀⠀⠀⠀⠈⠉⠑⠛⢿⣿⣿⣇⣿⠇⠁⡐⠀⢀⣬⣿⣿⣿⣿⣿⡄⠀	  %s
⠀⠀⠀⠀⠀⠀⣠⣴⣶⣴⣽⢿⣯⠸⡀⡐⠀⠐⢱⣙⣮⣭⣿⣿⠿⠷⠀
⠀⠀⠀⠀⣠⣻⣿⠿⣿⣿⣟⣩⡑⠀⢠⠁⠀⠮⠿⠶⠶⠦⣌⡁⠀⠀⠀
⠀⠀⠀⢈⣿⣿⣿⣿⣷⣿⣿⢟⣍⣤⣨⡴⣄⢎⢍⡖⡿⠽⣿⣟⣷⣄⠀
⠀⠀⠀⠸⣿⣿⣿⣿⢿⣯⣮⣟⣾⣿⠋⠘⠉⠑⣳⡻⣦⣯⣽⣿⣿⣿⡀
⠀⠀⠀⠀⠘⢿⣷⣽⣿⡿⡎⠛⠋⠁⠀⠀⠀⠀⢳⢿⣽⣷⣿⣝⢿⡿⠀	  %s
⠀⠀⠀⠀⠀⠀⠉⠙⠻⠋⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣿⣽⣾⡿⠟⠁⠀
`


const RED = "\033[31m"
const GREEN = "\033[32m"
const YELLOW = "\033[33m"
const BLUE = "\033[34m"
const PURPLE = "\033[35m"
const CYAN = "\033[36m"
const RESET = "\033[0m"

func getCurrentUsername() string {
	username, err := user.Current()
	if err != nil { return "unknown" }
	return username.Username
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil { return "unknown" }
	return hostname
}

func main() {
	shell := os.Getenv("SHELL")
	home := os.Getenv("HOME")
	username := getCurrentUsername()
	hostname := getHostname()

	formattedHostname := fmt.Sprintf("%s%s@%s%s%s",YELLOW, username, PURPLE,hostname, RESET)
	formattedHome := fmt.Sprintf("%s%s%s%s", "home:	",GREEN, home, RESET)
	formattedShell := fmt.Sprintf("%s%s%s%s", "shell:	",BLUE, shell, RESET)
	formattedColors := fmt.Sprintf("%s███%s███%s███%s███%s███%s███ %s", RED, GREEN, YELLOW, BLUE, PURPLE, CYAN, RESET)

	fmt.Printf(art, formattedHostname, formattedHome, formattedShell, formattedColors)
}
package main

import "github.com/0xBl4nk/FuzzSwarm2/cmd"

func main() {
  showBanner()
	cmd.Execute()
}

func showBanner() {
  banner := `
                           █████▒█    ██ ▒███████▒▒███████▒
                       ▓██   ▒ ██  ▓██▒▒ ▒ ▒ ▄▀░▒ ▒ ▒ ▄▀░
                       ▒████ ░▓██  ▒██░░ ▒ ▄▀▒░ ░ ▒ ▄▀▒░ 
                       ░▓█▒  ░▓▓█  ░██░  ▄▀▒   ░  ▄▀▒   ░
                       ░▒█░   ▒▒█████▓ ▒███████▒▒███████▒
                        ▒ ░   ░▒▓▒ ▒ ▒ ░▒▒ ▓░▒░▒░▒▒ ▓░▒░▒
                        ░     ░░▒░ ░ ░ ░░▒ ▒ ░ ▒░░▒ ▒ ░ ▒
                        ░ ░    ░░░ ░ ░ ░ ░ ░ ░ ░░ ░ ░ ░ ░
                                 ░       ░ ░      ░ ░    
                                       ░        ░        
                  ██████  █     █░ ▄▄▄       ██▀███   ███▄ ▄███▓
                ▒██    ▒ ▓█░ █ ░█░▒████▄    ▓██ ▒ ██▒▓██▒▀█▀ ██▒
                ░ ▓██▄   ▒█░ █ ░█ ▒██  ▀█▄  ▓██ ░▄█ ▒▓██    ▓██░
                  ▒   ██▒░█░ █ ░█ ░██▄▄▄▄██ ▒██▀▀█▄  ▒██    ▒██ 
                ▒██████▒▒░░██▒██▓  ▓█   ▓██▒░██▓ ▒██▒▒██▒   ░██▒
                ▒ ▒▓▒ ▒ ░░ ▓░▒ ▒   ▒▒   ▓▒█░░ ▒▓ ░▒▓░░ ▒░   ░  ░
                ░ ░▒  ░ ░  ▒ ░ ░    ▒   ▒▒ ░  ░▒ ░ ▒░░  ░      ░
                ░  ░  ░    ░   ░    ░   ▒     ░░   ░ ░      ░   
                      ░      ░          ░  ░   ░            ░   

  `
  println(banner)
}

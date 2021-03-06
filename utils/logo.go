package utils

import "github.com/fatih/color"

var (
	logo = `
 ___  __        ___  ___      ________      _______       ___  __        ___      _________   
|\  \|\  \     |\  \|\  \    |\   __  \    |\  ___ \     |\  \|\  \     |\  \    |\___   ___\ 
\ \  \/  /|_   \ \  \\\  \   \ \  \|\ /_   \ \   __/|    \ \  \/  /|_   \ \  \   \|___ \  \_| 
 \ \   ___  \   \ \  \\\  \   \ \   __  \   \ \  \_|/__   \ \   ___  \   \ \  \       \ \  \  
  \ \  \\ \  \   \ \  \\\  \   \ \  \|\  \   \ \  \_|\ \   \ \  \\ \  \   \ \  \       \ \  \ 
   \ \__\\ \__\   \ \_______\   \ \_______\   \ \_______\   \ \__\\ \__\   \ \__\       \ \__\
    \|__| \|__|    \|_______|    \|_______|    \|_______|    \|__| \|__|    \|__|        \|__|


KubeKit V%s ⓒ  OrientSoft 2017

`

	CheckSymbol = "\u2714 "
	CrossSymbol = "\u2716 "
)

func DisplayLogo(version string) {
	color.Green(logo, version)
}

# go-storyteller
A Golang reader for JSON representations of non-linear interactive texts/books/stories (CYOA).

(It will read and display any non-linear story in a JSON file.)

# About the app
go-storyteller is an innovative, small, single file application written in Go(lang), able to read and display non-linear stories/texts - also called CYOA (Choose Your Own Adventure).

It looks and behaves like a desktop application, but actually runs like a web app over a local server - no need for an internet connection -, and uses the computer's default browser as GUI.

![story - sm](https://github.com/manuelcaeiro/go-storyteller/blob/master/screenshots/story_sm.png)

# How to test it
If you have golang installed on your computer and know how to run a .go file (and you happen to not have a JSON file with a CYOA story at hand) just download the files gideon.json (1) and story.go onto the same directory, and run it or build your own executable.

Else, you can download an executable file suitable for your OS from [here](https://github.com/manuelcaeiro/go-storyteller/tree/master/downloads)...

- storyteller\_win.exe for Windows OS since version 7
- storyteller\_ubu for Ubuntu/Mint Linux since version 16.04/18.3
- storyteller\_deb for Debian Linux since version 9 (or any Debian 9 & up based distro)

... and the file gideon.json (1) to an empty folder and follow the instructions on the [READ2RUN.md](https://github.com/manuelcaeiro/go-storyteller/blob/master/READ2RUN.md) file.

(1) which contains a sample story; right click over the name of the file on the repository above and choose Save link as...

![story - ie](https://github.com/manuelcaeiro/go-storyteller/blob/master/screenshots/story_ie.JPG)

# Licences
Copyright [2019] [J. Manuel Caeiro D. P.]

The code in this repository is under the _MIT License_. You may obtain a copy of the _MIT License_ [here](https://opensource.org/licenses/MIT)

The story in the JSON file has the same copyright and is under the _Creative Commons License_ [Attribution-NonCommercial-NoDerivatives 4.0 International](https://creativecommons.org/licenses/by-nc-nd/4.0/)

# TO-DO
- Upon request (plus a suitable donation): Add a first page that displays a choice list between several (when available) stories.

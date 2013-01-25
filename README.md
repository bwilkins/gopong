# GoPong #

## Code ownership ##
I've based (read: mercilessly stolen) my code on [krepa098's sfmlPong sample](https://bitbucket.org/krepa098/gosfml2-samples).

I've started Go-ifying it however I can, and I'm using this starting point as a way to learn more about go, how to write games, and how to use SFML.

I haven't put a license on this code yet, as the original code didn't come with a license. I've done what I can to preserve commits from the original author, so that you can see it is not all my work.

## Playing the game ##
- When requested, press space to start playing, or escape to quit.
- You may press escape to quit at any time - you will not be prompted for confirmation
- Press up and down to move the paddle. You are the left paddle.

## Building the game ##
You will need:
- Your OS's build tools
- [SFML]( http://github.com/LaurentGomila/SFML ) and [CSFML]( http://github.com/LaurentGomila/CSFML ) (both need to be installed, SFML before CSFML)
- `$ go install bitbucket.org/krepa098/gosfml2`
- `$ go run main.go`

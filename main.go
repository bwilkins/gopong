/*
*********************************
*								*
*		GOSFML2					*
*		SFML Examples:	 Pong	*
*		Ported from C++ to Go	*
*********************************
 */

package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.LockOSThread()

	const (
		paddleSpeed = float32(400)
		ballSpeed   = float32(400)
	)

	var (
		gameWidth  uint        = 800
		gameHeight uint        = 600
		paddleSize sf.Vector2f = sf.Vector2f{25, 100}
		ballRadius float32     = 10
	)

	ticker := time.NewTicker(time.Second / 60)
	AITicker := time.NewTicker(time.Second / 10)
	rand.Seed(time.Now().UnixNano())

	renderWindow := sf.NewRenderWindow(sf.VideoMode{gameWidth, gameHeight, 32}, "Pong (GoSFML2)", sf.Style_DefaultStyle, nil)

	// Load the sounds used in the game
	buffer := sf.NewSoundBufferFromFile("resources/ball.wav")
	ballSound := sf.NewSound(buffer)

	// Create the left paddle
	leftPaddle := sf.NewRectangleShape()
	leftPaddle.SetSize(sf.Vector2f{paddleSize.X - 3, paddleSize.Y - 3})
	leftPaddle.SetOutlineThickness(3)
	leftPaddle.SetOutlineColor(sf.Color_Black)
	leftPaddle.SetFillColor(sf.Color{100, 100, 200, 255})
	leftPaddle.SetOrigin(sf.Vector2f{paddleSize.X / 2, paddleSize.Y / 2})

	// Create the right paddle
	rightPaddle := sf.NewRectangleShape()
	rightPaddle.SetSize(sf.Vector2f{paddleSize.X - 3, paddleSize.Y - 3})
	rightPaddle.SetOutlineThickness(3)
	rightPaddle.SetOutlineColor(sf.Color_Black)
	rightPaddle.SetFillColor(sf.Color{200, 100, 100, 255})
	rightPaddle.SetOrigin(sf.Vector2f{paddleSize.X / 2, paddleSize.Y / 2})

	// Create the ball
	ball := sf.NewCircleShape(ballRadius - 3)
	ball.SetOutlineThickness(3)
	ball.SetOutlineColor(sf.Color_Black)
	ball.SetFillColor(sf.Color_White)
	ball.SetOrigin(sf.Vector2f{ballRadius / 2, ballRadius / 2})

	// Load the text font
	font := sf.NewFontFromFile("resources/sansation.ttf")

	// Initialize the pause message
	pauseMessage := sf.NewText()
	pauseMessage.SetCharacterSize(40)
	pauseMessage.SetPosition(sf.Vector2f{170, 150})
	pauseMessage.SetColor(sf.Color_White)
	pauseMessage.SetString("Welcome to SFML pong!\nPress space to start the game")
	pauseMessage.SetFont(font)

	var (
		rightPaddleSpeed float32 = 0
		ballAngle        float32 = 0
		isPlaying        bool    = false
	)

	for renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			//poll events
			for event, eventType := renderWindow.PollEvent(); event != nil; event, eventType = renderWindow.PollEvent() {
				switch event.(type) {
				case *sf.KeyEvent:
					switch event.(*sf.KeyEvent).Code {
					case sf.Key_Escape:
						renderWindow.Close()
					case sf.Key_Space:
						if !isPlaying {
							// (re)start the game
							isPlaying = true

							// reset position of the paddles and ball
							leftPaddle.SetPosition(sf.Vector2f{10 + paddleSize.X/2, float32(gameHeight) / 2})
							rightPaddle.SetPosition(sf.Vector2f{float32(gameWidth) - 10 - paddleSize.X/2, float32(gameHeight) / 2})
							ball.SetPosition(sf.Vector2f{float32(gameWidth) / 2, float32(gameHeight) / 2})

							// reset the ball angle
							for {
								// Make sure the ball initial angle is not too much vertical
								ballAngle = rand.Float32() * math.Pi * 2
								if math.Abs(math.Cos(float64(ballAngle))) > 0.7 {
									break
								}
							}
						}
					}
				}
				if eventType == sf.Event_Closed {
					renderWindow.Close()
				}
			}

			//playing
			if isPlaying {
				deltaTime := time.Second / 60

				// Move the player's paddle
				if sf.Keyboard_IsKeyPressed(sf.Key_Up) && leftPaddle.GetPosition().Y-paddleSize.Y/2 > 5 {
					leftPaddle.Move(sf.Vector2f{0, -paddleSpeed * float32(deltaTime.Seconds())})
				}

				if sf.Keyboard_IsKeyPressed(sf.Key_Down) && leftPaddle.GetPosition().Y+paddleSize.Y/2 < float32(gameHeight)-5 {
					leftPaddle.Move(sf.Vector2f{0, paddleSpeed * float32(deltaTime.Seconds())})
				}

				// Move the computer's paddle
				if (rightPaddleSpeed < 0 && rightPaddle.GetPosition().Y-paddleSize.Y/2 > 5) || (rightPaddleSpeed > 0 && rightPaddle.GetPosition().Y+paddleSize.Y/2 < float32(gameHeight)-5) {
					rightPaddle.Move(sf.Vector2f{0, rightPaddleSpeed * float32(deltaTime.Seconds())})
				}

				// Move the ball
				factor := ballSpeed * float32(deltaTime.Seconds())
				ball.Move(sf.Vector2f{float32(math.Cos(float64(ballAngle))) * factor, float32(math.Sin(float64(ballAngle))) * factor})

				// Check collisions between the ball and the screen
				if ball.GetPosition().X-ballRadius < 0 {
					isPlaying = false
					pauseMessage.SetString("You lost !\nPress space to restart or\nescape to exit")
				}

				if ball.GetPosition().X+ballRadius > float32(gameWidth) {
					isPlaying = false
					pauseMessage.SetString("You won !\nPress space to restart or\nescape to exit")
				}

				if ball.GetPosition().Y-ballRadius < 0 {
					ballAngle = -ballAngle
					ball.SetPosition(sf.Vector2f{ball.GetPosition().X, ballRadius + 0.1})
					ballSound.Play()
				}

				if ball.GetPosition().Y+ballRadius > float32(gameHeight) {
					ballAngle = -ballAngle
					ball.SetPosition(sf.Vector2f{ball.GetPosition().X, float32(gameHeight) - ballRadius - 0.1})
					ballSound.Play()
				}

				// Check the collisions between the ball and the paddles
				// Left Paddle
				if ball.GetPosition().X-ballRadius < leftPaddle.GetPosition().X+paddleSize.X/2 &&
					ball.GetPosition().X-ballRadius > leftPaddle.GetPosition().X &&
					ball.GetPosition().Y+ballRadius >= leftPaddle.GetPosition().Y-paddleSize.Y/2 &&
					ball.GetPosition().Y-ballRadius <= leftPaddle.GetPosition().Y+paddleSize.Y/2 {

					if ball.GetPosition().Y > leftPaddle.GetPosition().Y {
						ballAngle = math.Pi - ballAngle + rand.Float32()*math.Pi*0.2
					} else {
						ballAngle = math.Pi - ballAngle - rand.Float32()*math.Pi*0.2
					}

					ball.SetPosition(sf.Vector2f{leftPaddle.GetPosition().X + ballRadius + paddleSize.X/2 + 0.1, ball.GetPosition().Y})
					ballSound.Play()
				}

				// Right Paddle
				if ball.GetPosition().X+ballRadius > rightPaddle.GetPosition().X-paddleSize.X/2 &&
					ball.GetPosition().X+ballRadius < rightPaddle.GetPosition().X &&
					ball.GetPosition().Y+ballRadius >= rightPaddle.GetPosition().Y-paddleSize.Y/2 &&
					ball.GetPosition().Y-ballRadius <= rightPaddle.GetPosition().Y+paddleSize.Y/2 {

					if ball.GetPosition().Y > rightPaddle.GetPosition().Y {
						ballAngle = math.Pi - ballAngle + rand.Float32()*math.Pi*0.2
					} else {
						ballAngle = math.Pi - ballAngle - rand.Float32()*math.Pi*0.2
					}

					ball.SetPosition(sf.Vector2f{rightPaddle.GetPosition().X - ballRadius - paddleSize.X/2 - 0.1, ball.GetPosition().Y})
					ballSound.Play()
				}
			}

			// Clear the window
			renderWindow.Clear(sf.Color{50, 200, 50, 0})

			if isPlaying {
				renderWindow.Draw(leftPaddle, nil)
				renderWindow.Draw(rightPaddle, nil)
				renderWindow.Draw(ball, nil)
			} else {
				renderWindow.Draw(pauseMessage, nil)
			}

			// Display things on screen
			renderWindow.Display()
		case <-AITicker.C:
			if ball.GetPosition().Y+ballRadius > rightPaddle.GetPosition().Y+paddleSize.Y/2 {
				rightPaddleSpeed = paddleSpeed
			} else if ball.GetPosition().Y-ballRadius < rightPaddle.GetPosition().Y-paddleSize.Y/2 {
				rightPaddleSpeed = -paddleSpeed
			} else {
				rightPaddleSpeed = 0
			}
		}
	}
}
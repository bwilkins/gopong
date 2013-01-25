package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"github.com/bwilkins/gopong/ball"
	"github.com/bwilkins/gopong/paddle"
	"math"
	"math/rand"
	"time"
)

func main() {
	//Define some variables for the game
	//This block mostly will not change
	paddleMaxSpeed := float32(400.0)
	paddleDefaultSize := sf.Vector2f{25, 100}
	isPlaying := false

	gameWidth := uint(800)
	gameHeight := uint(600)
	bitDepth := uint(32)

	ballMaxSpeed := float32(400.0)
	ballRadius := float32(10.0)

	//These are a little more special... guess what they do!
	ticker := time.NewTicker(time.Second / 60)
	aiTicker := time.NewTicker(time.Second / 10)
	rand.Seed(time.Now().UnixNano())

	//Instantiate the render window for SFML
	renderWindow := sf.NewRenderWindow(sf.VideoMode{gameWidth, gameHeight, bitDepth}, "Pong (Brett's Go test)", sf.Style_DefaultStyle, nil)

	//Create the left paddle
	leftPaddle := paddle.NewPaddle(paddleMaxSpeed, paddleMaxSpeed, paddleDefaultSize, sf.Color{100, 100, 200, 255})

	//Create the right paddle
	rightPaddle := paddle.NewPaddle(0, paddleMaxSpeed, paddleDefaultSize, sf.Color{200, 100, 100, 255})

	//Create the ball
	ball := ball.NewBall(ballMaxSpeed, ballMaxSpeed, ballRadius, "resources/ball.wa")

	//Load font
	font, _ := sf.NewFontFromFile("resources/sansation.ttf")

	//Init the pause message
	pauseMessage := sf.NewText(font)
	pauseMessage.SetCharacterSize(40)
	pauseMessage.SetPosition(sf.Vector2f{170, 150})
	pauseMessage.SetColor(sf.ColorWhite())
	pauseMessage.SetString("Welcome to Brett's SFML Pong!\nPress space to start the game.")

	for renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			//poll for events
			for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
				switch ev := event.(type) {
				case sf.EventKeyReleased:
					switch ev.Code {
					case sf.Key_Escape:
						renderWindow.Close()
					case sf.Key_Space:
						if !isPlaying {
							//restart the game
							isPlaying = true
							leftPaddle.shape.SetPosition(sf.Vector2f{10 + leftPaddle.size.X/2, float32(gameHeight) / 2})
							rightPaddle.shape.SetPosition(sf.Vector2f{float32(gameWidth) - 10 - rightPaddle.size.X/2, float32(gameHeight) / 2})
							ball.shape.SetPosition(sf.Vector2f{float32(gameWidth) / 2, float32(gameHeight) / 2})

							//ensure the ball angle isn't too vertical
							for {
								ball.angle = rand.Float32() * math.Pi * 2

								if math.Abs(math.Cos(float64(ball.angle))) > 0.7 {
									break
								}
							}
						}
					}
				case sf.EventClosed:
					renderWindow.Close()
				}
			}

			if isPlaying {
				deltaTime := time.Second / 60

				//Move the player's paddle
				if sf.KeyboardIsKeyPressed(sf.Key_Up) && leftPaddle.TopLeft().Y > 5 {
					leftPaddle.shape.Move(sf.Vector2f{0, -leftPaddle.speed * float32(deltaTime.Seconds())})
				}
				if sf.KeyboardIsKeyPressed(sf.Key_Down) && leftPaddle.BottomRight().Y < float32(gameHeight)-5 {
					leftPaddle.shape.Move(sf.Vector2f{0, leftPaddle.speed * float32(deltaTime.Seconds())})
				}

				//Move the ai's paddle
				if (rightPaddle.speed < 0 && rightPaddle.TopLeft().Y > 5) || (rightPaddle.speed > 0 && rightPaddle.BottomRight().Y < float32(gameHeight)-5) {
					rightPaddle.shape.Move(sf.Vector2f{0, rightPaddle.speed * float32(deltaTime.Seconds())})
				}

				//Move ze ball
				factor := ball.speed * float32(deltaTime.Seconds())
				ball.shape.Move(sf.Vector2f{float32(math.Cos(float64(ball.angle))) * factor, float32(math.Sin(float64(ball.angle))) * factor})

				//Check collisions between ball and screen edge
				if ball.TopLeft().X < 0 {
					isPlaying = false
					pauseMessage.SetString("You lost!\nPress space to restart or\nescape to quit")
				}

				if ball.BottomRight().X > float32(gameWidth) {
					isPlaying = false
					pauseMessage.SetString("You won!\nPress space to play again or\nescape to quit")
				}

				if ball.TopLeft().Y < 0 {
					ball.angle = -ball.angle
					ball.shape.SetPosition(sf.Vector2f{ball.Center().X, ball.radius + 0.1})
					ball.sound.Play()
				}

				if ball.BottomRight().Y > float32(gameHeight) {
					ball.angle = -ball.angle
					ball.shape.SetPosition(sf.Vector2f{ball.Center().X, float32(gameHeight) - ball.radius - 0.1})
					ball.sound.Play()
				}

				//Check collisions between the ball and the left paddle
				if ball.TopLeft().X < leftPaddle.BottomRight().X &&
					ball.TopLeft().X > leftPaddle.Center().X &&
					ball.BottomRight().Y >= leftPaddle.TopLeft().Y &&
					ball.TopLeft().Y <= leftPaddle.BottomRight().Y {

					if ball.Center().Y > leftPaddle.Center().Y {
						ball.angle = math.Pi - ball.angle + rand.Float32()*math.Pi*0.2
					} else {
						ball.angle = math.Pi - ball.angle - rand.Float32()*math.Pi*0.2
					}

					ball.shape.SetPosition(sf.Vector2f{leftPaddle.Center().X + ball.radius + leftPaddle.size.X/2 + 0.1, ball.Center().Y})
					ball.sound.Play()
				}

				//Check collisions between the ball and the right paddle
				if ball.BottomRight().X > rightPaddle.TopLeft().X &&
					ball.BottomRight().X < rightPaddle.Center().X &&
					ball.BottomRight().Y >= rightPaddle.TopLeft().Y &&
					ball.TopLeft().Y <= rightPaddle.BottomRight().Y {

					if ball.Center().Y > rightPaddle.Center().Y {
						ball.angle = math.Pi - ball.angle + rand.Float32()*math.Pi*0.2
					} else {
						ball.angle = math.Pi - ball.angle - rand.Float32()*math.Pi*0.2
					}

					ball.shape.SetPosition(sf.Vector2f{rightPaddle.Center().X - ball.radius - rightPaddle.size.X/2 - 0.1, ball.Center().Y})
					ball.sound.Play()
				}
			}

			//Clear the window
			renderWindow.Clear(sf.Color{50, 200, 50, 0})

			//Draw some shit
			if isPlaying {
				renderWindow.Draw(leftPaddle.shape, nil)
				renderWindow.Draw(rightPaddle.shape, nil)
				renderWindow.Draw(ball.shape, nil)
			} else {
				renderWindow.Draw(pauseMessage, nil)
			}

			//Draw everything to the screen
			renderWindow.Display()

		case <-aiTicker.C:
			if ball.BottomRight().Y > rightPaddle.BottomRight().Y {
				rightPaddle.speed = rightPaddle.max_speed
			} else if ball.TopLeft().Y < rightPaddle.TopLeft().Y {
				rightPaddle.speed = -rightPaddle.max_speed
			} else {
				rightPaddle.speed = 0
			}
		}
	}

}

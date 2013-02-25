package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"github.com/bwilkins/gopong/ball"
	"github.com/bwilkins/gopong/paddle"
	"math"
	"math/rand"
	"time"
)

type Collider interface {
	TopLeft() sf.Vector2f
	BottomRight() sf.Vector2f
	Center() sf.Vector2f
}

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
	ball := ball.NewBall(ballMaxSpeed, ballMaxSpeed, ballRadius, "resources/ball.wav")

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
							leftPaddle.Shape.SetPosition(sf.Vector2f{10 + leftPaddle.Size.X/2, float32(gameHeight) / 2})
							rightPaddle.Shape.SetPosition(sf.Vector2f{float32(gameWidth) - 10 - rightPaddle.Size.X/2, float32(gameHeight) / 2})
							ball.Shape.SetPosition(sf.Vector2f{float32(gameWidth) / 2, float32(gameHeight) / 2})

							//ensure the ball angle isn't too vertical
							for {
								ball.Angle = rand.Float32() * math.Pi * 2

								if math.Abs(math.Cos(float64(ball.Angle))) > 0.7 {
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
					leftPaddle.Shape.Move(sf.Vector2f{0, -leftPaddle.Speed * float32(deltaTime.Seconds())})
				}
				if sf.KeyboardIsKeyPressed(sf.Key_Down) && leftPaddle.BottomRight().Y < float32(gameHeight)-5 {
					leftPaddle.Shape.Move(sf.Vector2f{0, leftPaddle.Speed * float32(deltaTime.Seconds())})
				}

				//Move the ai's paddle
				if (rightPaddle.Speed < 0 && rightPaddle.TopLeft().Y > 5) || (rightPaddle.Speed > 0 && rightPaddle.BottomRight().Y < float32(gameHeight)-5) {
					rightPaddle.Shape.Move(sf.Vector2f{0, rightPaddle.Speed * float32(deltaTime.Seconds())})
				}

				//Move ze ball
				factor := ball.Speed * float32(deltaTime.Seconds())
				ball.Shape.Move(sf.Vector2f{float32(math.Cos(float64(ball.Angle))) * factor, float32(math.Sin(float64(ball.Angle))) * factor})

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
					ball.Angle = -ball.Angle
					ball.Shape.SetPosition(sf.Vector2f{ball.Center().X, ball.Radius + 0.1})
					ball.Sound.Play()
				}

				if ball.BottomRight().Y > float32(gameHeight) {
					ball.Angle = -ball.Angle
					ball.Shape.SetPosition(sf.Vector2f{ball.Center().X, float32(gameHeight) - ball.Radius - 0.1})
					ball.Sound.Play()
				}

				//Check collisions between the ball and the left paddle
				if leftPaddle.CollideRight(ball) {

					if ball.Center().Y > leftPaddle.Center().Y {
						ball.Angle = math.Pi - ball.Angle + rand.Float32()*math.Pi*0.2
					} else {
						ball.Angle = math.Pi - ball.Angle - rand.Float32()*math.Pi*0.2
					}

					ball.Shape.SetPosition(sf.Vector2f{leftPaddle.Center().X + ball.Radius + leftPaddle.Size.X/2 + 0.1, ball.Center().Y})
					ball.Sound.Play()
				}

				//Check collisions between the ball and the right paddle
				if rightPaddle.CollideLeft(ball) {

					if ball.Center().Y > rightPaddle.Center().Y {
						ball.Angle = math.Pi - ball.Angle + rand.Float32()*math.Pi*0.2
					} else {
						ball.Angle = math.Pi - ball.Angle - rand.Float32()*math.Pi*0.2
					}

					ball.Shape.SetPosition(sf.Vector2f{rightPaddle.Center().X - ball.Radius - rightPaddle.Size.X/2 - 0.1, ball.Center().Y})
					ball.Sound.Play()
				}
			}

			//Clear the window
			renderWindow.Clear(sf.Color{50, 200, 50, 0})

			//Draw some shit
			if isPlaying {
				renderWindow.Draw(leftPaddle.Shape, nil)
				renderWindow.Draw(rightPaddle.Shape, nil)
				renderWindow.Draw(ball.Shape, nil)
			} else {
				renderWindow.Draw(pauseMessage, nil)
			}

			//Draw everything to the screen
			renderWindow.Display()

		case <-aiTicker.C:
			if ball.BottomRight().Y > rightPaddle.BottomRight().Y {
				rightPaddle.Speed = rightPaddle.MaxSpeed
			} else if ball.TopLeft().Y < rightPaddle.TopLeft().Y {
				rightPaddle.Speed = -rightPaddle.MaxSpeed
			} else {
				rightPaddle.Speed = 0
			}
		}
	}

}

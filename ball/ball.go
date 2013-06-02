package ball

import sf "bitbucket.org/krepa098/gosfml2"

type Ball struct {
	Speed    float32
	MaxSpeed float32
	Angle    float32
	Radius   float32
	Shape    *sf.CircleShape
	Sound    *sf.Sound
}

func NewBall(speed, max_speed, radius float32, sound_file string) *Ball {
	//Once again, accounting for outline thickness
	shape := sf.NewCircleShape(radius - 3)
	shape.SetOutlineThickness(3)
	shape.SetOutlineColor(sf.ColorBlack())
	shape.SetFillColor(sf.ColorWhite())
	shape.SetOrigin(sf.Vector2f{radius / 2, radius / 2})

	buffer, _ := sf.NewSoundBufferFromFile(sound_file)
	sound := sf.NewSound(buffer)

	return &Ball{speed, max_speed, float32(0), radius, shape, sound}
}

func (b *Ball) Center() sf.Vector2f {
	return b.Shape.GetPosition()
}

func (b *Ball) TopLeft() sf.Vector2f {
	return sf.Vector2f{b.Center().X - b.Radius, b.Center().Y - b.Radius}
}

func (b *Ball) BottomRight() sf.Vector2f {
	return sf.Vector2f{b.Center().X + b.Radius, b.Center().Y + b.Radius}
}

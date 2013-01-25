package ball

import sf "bitbucket.org/krepa098/gosfml2"

type Ball struct {
	speed     float32
	max_speed float32
	angle     float32
	radius    float32
	shape     *sf.CircleShape
	sound     *sf.Sound
}

func NewBall(speed, max_speed, radius float32, sound_file string) *Ball {
	//Once again, accounting for outline thickness
	shape := sf.NewCircleShape(radius - 3)
	shape.SetOutlineThickness(3)
	shape.SetOutlineColor(sf.ColorBlack())
	shape.SetFillColor(sf.ColorWhite())
	shape.SetOrigin(sf.Vector2f{radius / 2, radius / 2})

	buffer := sf.NewSoundBufferFromFile(sound_file)
	sound := sf.NewSound(buffer)

	return &Ball{speed, max_speed, float32(0), radius, shape, sound}
}

func (b *Ball) TopLeft() sf.Vector2f {
	return sf.Vector2f{b.shape.GetPosition().X - b.radius, b.shape.GetPosition().Y - b.radius}
}

func (b *Ball) BottomRight() sf.Vector2f {
	return sf.Vector2f{b.shape.GetPosition().X + b.radius, b.shape.GetPosition().Y + b.radius}
}

func (b *Ball) Center() sf.Vector2f {
	return b.shape.GetPosition()
}

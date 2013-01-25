package paddle

import sf "bitbucket.org/krepa098/gosfml2"

type Paddle struct {
	speed     float32
	max_speed float32
	size      sf.Vector2f
	shape     *sf.RectangleShape
}

func NewPaddle(speed, max_speed float32, size sf.Vector2f, color sf.Color) *Paddle {
	shape := sf.NewRectangleShape()
	//Take 3 off each edge to account for outline thickness
	shape.SetSize(sf.Vector2f{size.X - 3, size.Y - 3})
	shape.SetOutlineThickness(3)
	shape.SetOutlineColor(sf.ColorBlack())
	shape.SetFillColor(color)
	shape.SetOrigin(sf.Vector2f{size.X / 2, size.Y / 2})

	return &Paddle{speed, max_speed, size, shape}
}

func (p *Paddle) TopLeft() sf.Vector2f {
	return sf.Vector2f{p.shape.GetPosition().X - p.size.X/2, p.shape.GetPosition().Y - p.size.Y/2}
}

func (p *Paddle) BottomRight() sf.Vector2f {
	return sf.Vector2f{p.shape.GetPosition().X + p.size.X/2, p.shape.GetPosition().Y + p.size.Y/2}
}

func (p *Paddle) Center() sf.Vector2f {
	return p.shape.GetPosition()
}

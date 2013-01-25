package paddle

import sf "bitbucket.org/krepa098/gosfml2"

type Paddle struct {
	Speed    float32
	MaxSpeed float32
	Size     sf.Vector2f
	Shape    *sf.RectangleShape
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
	return sf.Vector2f{p.Shape.GetPosition().X - p.Size.X/2, p.Shape.GetPosition().Y - p.Size.Y/2}
}

func (p *Paddle) BottomRight() sf.Vector2f {
	return sf.Vector2f{p.Shape.GetPosition().X + p.Size.X/2, p.Shape.GetPosition().Y + p.Size.Y/2}
}

func (p *Paddle) Center() sf.Vector2f {
	return p.Shape.GetPosition()
}

package paddle

import sf "bitbucket.org/krepa098/gosfml2"
import "github.com/bwilkins/gopong/collider"

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

func (p *Paddle) Center() sf.Vector2f {
	return p.Shape.GetPosition()
}

func (p *Paddle) TopLeft() sf.Vector2f {
	return sf.Vector2f{p.Center().X - p.Size.X/2, p.Center().Y - p.Size.Y/2}
}

func (p *Paddle) BottomRight() sf.Vector2f {
	return sf.Vector2f{p.Center().X + p.Size.X/2, p.Center().Y + p.Size.Y/2}
}

func (p *Paddle) CollideLeft(c collider.Collider) bool {
	c_br := c.BottomRight()
	c_tl := c.TopLeft()
	p_br := p.BottomRight()
	p_tl := p.TopLeft()
	if c_br.X > p_tl.X {
		if c_br.Y < p_br.Y && c_tl.Y > p_tl.Y {
			return true
		}
	}
	return false
}

func (p *Paddle) CollideRight(c collider.Collider) bool {
	c_br := c.BottomRight()
	c_tl := c.TopLeft()
	p_br := p.BottomRight()
	p_tl := p.TopLeft()
	if c_tl.X < p_br.X {
		if c_br.Y < p_br.Y && c_tl.Y > p_tl.Y {
			return true
		}
	}
	return false
}

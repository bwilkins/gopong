package collider

import sf "bitbucket.org/krepa098/gosfml2"

type Collider interface {
	TopLeft() sf.Vector2f
	BottomRight() sf.Vector2f
	Center() sf.Vector2f
}

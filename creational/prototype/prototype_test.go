package prototype

import "testing"

func TestPrototype(t *testing.T) {
	circle := cloneRegistry[CircleShape]
	rectangle := cloneRegistry[RectangleShape]

	c, ok := circle.(Circle)
	if !ok {
		t.Errorf("Circle is not a circle")
	}

	c.Radius = 10
	if c.Radius != 10 {
		t.Errorf("Circle radius is not 10")
	}
	
	r, ok := rectangle.(Rectangle)
	if !ok {
		t.Errorf("Rectangle is not a rectangle")
	}

	r.Length = 10
	r.Width = 20
	if r.Length != 10 {
		t.Errorf("Rectangle length is not 10")
	}

	if r.Width != 20 {
		t.Errorf("Rectangle width is not 20")
	}

	if c.ShapeType() != CircleShape {
		t.Errorf("Circle clone is not a circle")
	}

	if r.ShapeType() != RectangleShape {
		t.Errorf("Rectangle clone is not a rectangle")
	}

	c.Display()
	r.Display()
}
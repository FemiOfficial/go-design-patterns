package prototype

import "testing"

func TestPrototype(t *testing.T) {
	circleClone := cloneRegistry[CircleShape]
	rectangleClone := cloneRegistry[RectangleShape]

	circle, ok := circleClone.(Circle)
	if !ok {
		t.Errorf("Circle is not a circle")
	}

	circle.Radius = 10
	if circle.Radius != 10 {
		t.Errorf("Circle radius is not 10")
	}
	
	rectangle, ok := rectangleClone.(Rectangle)
	if !ok {
		t.Errorf("Rectangle is not a rectangle")
	}

	rectangle.Length = 10
	rectangle.Width = 20
	if rectangle.Length != 10 {
		t.Errorf("Rectangle length is not 10")
	}

	if rectangle.Width != 20 {
		t.Errorf("Rectangle width is not 20")
	}

	if circle.ShapeType() != CircleShape {
		t.Errorf("Circle clone is not a circle")
	}

	if rectangle.ShapeType() != RectangleShape {
		t.Errorf("Rectangle clone is not a rectangle")
	}

	circle.Display()
	rectangle.Display()
}
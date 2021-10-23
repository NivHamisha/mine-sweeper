package utils

type Cell struct {
	value int
	isPressed bool
}

func (c *Cell) IsMine()bool{
	if c.value == -1 {
		return true
	}
	return false
}

func (c *Cell) UpdateValue(newValue int){
	c.value = newValue
}

func (c *Cell) Press(){
	c.isPressed = true
}


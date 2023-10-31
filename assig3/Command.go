package main

import "fmt"

type Command interface {
	Execute()
}

type Light struct {
	On bool
}

func (l *Light) TurnOn() {
	l.On = true
	fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
	l.On = false
	fmt.Println("Light is off")
}

type TurnOnCommand struct {
	Light *Light
}

func (c *TurnOnCommand) Execute() {
	c.Light.TurnOn()
}

type TurnOffCommand struct {
	Light *Light
}

func (c *TurnOffCommand) Execute() {
	c.Light.TurnOff()
}

type RemoteControl struct {
	Command Command
}

func (r *RemoteControl) PressButton() {
	r.Command.Execute()
}

func main() {
	light := &Light{}
	remote := &RemoteControl{}

	turnOnCommand := &TurnOnCommand{Light: light}
	turnOffCommand := &TurnOffCommand{Light: light}

	remote.Command = turnOnCommand
	remote.PressButton()

	remote.Command = turnOffCommand
	remote.PressButton()
}

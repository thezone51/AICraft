package main

import (
  "fmt"
)

// just for convenience
type Id int

// что бы не копировать всю структуру целиком в connect(), пришлось сделать указатель на указатель. наверное неэффективно по времени, но по памяти точно
type Neuron struct {
  Input  []**Neuron
  Output []**Neuron
  W float64
  State float64
}


func (N *Neuron) SetWeight(v float64) {
  N.W = v
}

func (N *Neuron) StateIs(v float64) {
  N.State = v
}

func (N *Neuron) Connect(C *Neuron) {
  if N.Output[len(N.Output)-1] != nil {}
  N.Output[0] = &C
  C.Input[0] = &N
}


var n1, n2 Neuron

func main() {
  n1.OutputSize(1)
  n2.InputSize(1)
  n1.Connect(&n2)

  fmt.Println(n1)
  fmt.Println()
  fmt.Println(n2)

  n1.OutputSize(2)
  n2.InputSize(2)

  fmt.Println(n1)
  fmt.Println()
  fmt.Println(n2)
}

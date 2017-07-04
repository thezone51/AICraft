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
  N.Output = append(N.Output, &C)
  C.Input  = append(C.Input, &N)
}

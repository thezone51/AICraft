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
  if N == C {panic("neuron was self connected")}
  var haveNconnect, haveCconnect bool
  if (len(N.Output) !=0) && (len(C.Input) != 0) {
    for i := 0; i < len(N.Output); i++ {
      if N.Output[i] == C {haveNconnect = true}
    }
    if !haveNconnect {N.Output  = append(N.Output, C)}
    for i := 0; i < len(N.Output); i++ {
      if C.Input[i] == N {haveCconnect = true}
    }
    if !haveCconnect {C.Input  = append(C.Input, N)}
  } else {
    N.Output  = append(N.Output, C)
    C.Input  = append(C.Input, N)
  }
}

package gonet

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

func sigmoid(x float64) float64{
  return 1.0 / (1.0 + math.Exp(x*-1))
}

func mse(perfect, actual float64) float64{
  return (perfect*perfect - 2*perfect*actual + actual*actual)
}

func erroutput(actual, expected float64) float64{
  return (actual - expected)
}

func errhidden(weight, delta float64) float64 {
  return (weight*delta)
}

func delta(err,x float64) float64{
  return err*sigmoid(x)*(1-sigmoid(x))
}

func wieght_change(weight,output,delta float64) float64 {
  return (LEARNING_RATE*0.1 + weight - output * delta * LEARNING_RATE)
}


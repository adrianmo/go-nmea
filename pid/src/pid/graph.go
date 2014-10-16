
package pid

import (
  "code.google.com/p/plotinum/plot"
  "code.google.com/p/plotinum/plotter"
  "image/color"
  "io"
)

const (
  Title = "Vessel temperature"
  XLabel = "Seconds"
  YLabel = "Celcius"
)

type Graph struct {
  Plot *plot.Plot
  inputs plotter.XYs
  outputs plotter.XYs
}

type Point struct {
  X, Y float64
}

func NewGraph() (*Graph, error) {
  g := &Graph{}
  var err error
  g.Plot, err = plot.New()
  if err != nil {
    return nil, err
  }
  g.Plot.Add(plotter.NewGrid())
  g.Plot.Title.Text = Title
  g.Plot.X.Label.Text = XLabel
  g.Plot.Y.Label.Text = YLabel
  g.inputs = make(plotter.XYs, 100)
  g.outputs = make(plotter.XYs, 100)
  return g, nil
}

func (g *Graph) AddInput(x, y float64) {
  point := Point{}
  point.X = x
  point.Y = y
  g.inputs = append(g.inputs, point)
}

func (g *Graph) AddOutput(x, y float64) {
  point := Point{}
  point.X = x
  point.Y = y
  g.outputs = append(g.outputs, point)
}

func (g *Graph) Draw() error {
  l, err := plotter.NewLine(g.inputs)
  if err != nil {
    return err
  }
  l.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
  g.Plot.Add(l)

  p, errp := plotter.NewLine(g.outputs)
  if errp != nil {
    return errp
  }
  p.LineStyle.Color = color.RGBA{R: 64, G: 64, B: 196, A: 196}
  g.Plot.Add(p)

  g.Plot.Legend.Add("Temperature", l)
  return nil
}


func (g *Graph) PngWriter() io.WriterTo {
  return g.Plot.PngWriter(6, 3)
}

func (g *Graph) Save() error {
  if err := g.Plot.Display(12, 6) ; err != nil {
    return err
  }
  return nil
}


package pid

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"image/color"
	"io"
)

const (
	// The title on the graph.
	Title = "Vessel temperature"
	// The label of the X axis.
	XLabel = "Seconds"
	// The label of the Y axis.
	YLabel = "Celcius"
)

// A Graph is a plot of the PID system.
type Graph struct {
	// Plot is the plot struct.
	Plot *plot.Plot
	// inputs are the values for the PID input.
	inputs plotter.XYs
	// outputs are the values for the PID output.
	outputs plotter.XYs
}

// A Point is a point on the graph.
type Point struct {
	// The co-ordinates of the point.
	X, Y float64
}

// NewGraph creates a new Graph object.
func NewGraph() (*Graph, error) {
	g := &Graph{}

	g.inputs = make(plotter.XYs, 100)
	g.outputs = make(plotter.XYs, 100)
	return g, nil
}

// AddInput adds a new point to the 'input' plot.
func (g *Graph) AddInput(x, y float64) {
	point := Point{}
	point.X = x
	point.Y = y
	g.inputs = append(g.inputs, point)
}

// AddOutput adds a new point to the 'output' plot.
func (g *Graph) AddOutput(x, y float64) {
	point := Point{}
	point.X = x
	point.Y = y
	g.outputs = append(g.outputs, point)
}

// Draw renders the graph.
func (g *Graph) Draw() error {
	var err error
	g.Plot, err = plot.New()
	if err != nil {
		return err
	}
	g.Plot.Add(plotter.NewGrid())
	g.Plot.Title.Text = Title
	g.Plot.X.Label.Text = XLabel
	g.Plot.Y.Label.Text = YLabel

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

// PngWriter returns a writer for a PNG rendering of the graph.
func (g *Graph) PngWriter() io.WriterTo {
	return g.Plot.PngWriter(6, 3)
}

// Save displays the graph.
func (g *Graph) Save() error {
	if err := g.Plot.Display(12, 6); err != nil {
		return err
	}
	return nil
}

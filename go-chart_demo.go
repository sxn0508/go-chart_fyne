package main

import (
	"fmt"
	"image"
	"math/rand"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func PieChart() *image.Image {
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	collector := &chart.ImageWriter{}
	pie.Render(chart.PNG, collector)

	image, err := collector.Image()

	if err != nil {
		panic(err)
	}
	return &image
}
func xvalues() []time.Time {
	rawx := []string{"2015-07-17", "2015-07-20", "2015-07-21", "2015-07-22", "2015-07-23", "2015-07-24", "2015-07-27", "2015-07-28", "2015-07-29", "2015-07-30", "2015-07-31", "2015-08-03", "2015-08-04", "2015-08-05", "2015-08-06", "2015-08-07", "2015-08-10", "2015-08-11", "2015-08-12", "2015-08-13", "2015-08-14", "2015-08-17", "2015-08-18", "2015-08-19", "2015-08-20", "2015-08-21", "2015-08-24", "2015-08-25", "2015-08-26", "2015-08-27", "2015-08-28", "2015-08-31", "2015-09-01", "2015-09-02", "2015-09-03", "2015-09-04", "2015-09-08", "2015-09-09", "2015-09-10", "2015-09-11", "2015-09-14", "2015-09-15", "2015-09-16", "2015-09-17", "2015-09-18", "2015-09-21", "2015-09-22", "2015-09-23", "2015-09-24", "2015-09-25", "2015-09-28", "2015-09-29", "2015-09-30", "2015-10-01", "2015-10-02", "2015-10-05", "2015-10-06", "2015-10-07", "2015-10-08", "2015-10-09", "2015-10-12", "2015-10-13", "2015-10-14", "2015-10-15", "2015-10-16", "2015-10-19", "2015-10-20", "2015-10-21", "2015-10-22", "2015-10-23", "2015-10-26", "2015-10-27", "2015-10-28", "2015-10-29", "2015-10-30", "2015-11-02", "2015-11-03", "2015-11-04", "2015-11-05", "2015-11-06", "2015-11-09", "2015-11-10", "2015-11-11", "2015-11-12", "2015-11-13", "2015-11-16", "2015-11-17", "2015-11-18", "2015-11-19", "2015-11-20", "2015-11-23", "2015-11-24", "2015-11-25", "2015-11-27", "2015-11-30", "2015-12-01", "2015-12-02", "2015-12-03", "2015-12-04", "2015-12-07", "2015-12-08", "2015-12-09", "2015-12-10", "2015-12-11", "2015-12-14", "2015-12-15", "2015-12-16", "2015-12-17", "2015-12-18", "2015-12-21", "2015-12-22", "2015-12-23", "2015-12-24", "2015-12-28", "2015-12-29", "2015-12-30", "2015-12-31", "2016-01-04", "2016-01-05", "2016-01-06", "2016-01-07", "2016-01-08", "2016-01-11", "2016-01-12", "2016-01-13", "2016-01-14", "2016-01-15", "2016-01-19", "2016-01-20", "2016-01-21", "2016-01-22", "2016-01-25", "2016-01-26", "2016-01-27", "2016-01-28", "2016-01-29", "2016-02-01", "2016-02-02", "2016-02-03", "2016-02-04", "2016-02-05", "2016-02-08", "2016-02-09", "2016-02-10", "2016-02-11", "2016-02-12", "2016-02-16", "2016-02-17", "2016-02-18", "2016-02-19", "2016-02-22", "2016-02-23", "2016-02-24", "2016-02-25", "2016-02-26", "2016-02-29", "2016-03-01", "2016-03-02", "2016-03-03", "2016-03-04", "2016-03-07", "2016-03-08", "2016-03-09", "2016-03-10", "2016-03-11", "2016-03-14", "2016-03-15", "2016-03-16", "2016-03-17", "2016-03-18", "2016-03-21", "2016-03-22", "2016-03-23", "2016-03-24", "2016-03-28", "2016-03-29", "2016-03-30", "2016-03-31", "2016-04-01", "2016-04-04", "2016-04-05", "2016-04-06", "2016-04-07", "2016-04-08", "2016-04-11", "2016-04-12", "2016-04-13", "2016-04-14", "2016-04-15", "2016-04-18", "2016-04-19", "2016-04-20", "2016-04-21", "2016-04-22", "2016-04-25", "2016-04-26", "2016-04-27", "2016-04-28", "2016-04-29", "2016-05-02", "2016-05-03", "2016-05-04", "2016-05-05", "2016-05-06", "2016-05-09", "2016-05-10", "2016-05-11", "2016-05-12", "2016-05-13", "2016-05-16", "2016-05-17", "2016-05-18", "2016-05-19", "2016-05-20", "2016-05-23", "2016-05-24", "2016-05-25", "2016-05-26", "2016-05-27", "2016-05-31", "2016-06-01", "2016-06-02", "2016-06-03", "2016-06-06", "2016-06-07", "2016-06-08", "2016-06-09", "2016-06-10", "2016-06-13", "2016-06-14", "2016-06-15", "2016-06-16", "2016-06-17", "2016-06-20", "2016-06-21", "2016-06-22", "2016-06-23", "2016-06-24", "2016-06-27", "2016-06-28", "2016-06-29", "2016-06-30", "2016-07-01", "2016-07-05", "2016-07-06", "2016-07-07", "2016-07-08", "2016-07-11", "2016-07-12", "2016-07-13", "2016-07-14", "2016-07-15"}

	var dates []time.Time
	for _, ts := range rawx {
		parsed, _ := time.Parse(chart.DefaultDateFormat, ts)
		dates = append(dates, parsed)
	}
	return dates
}

func yvalues() []float64 {
	return []float64{212.47, 212.59, 211.76, 211.37, 210.18, 208.00, 206.79, 209.33, 210.77, 210.82, 210.50, 209.79, 209.38, 210.07, 208.35, 207.95, 210.57, 208.66, 208.92, 208.66, 209.42, 210.59, 209.98, 208.32, 203.97, 197.83, 189.50, 187.27, 194.46, 199.27, 199.28, 197.67, 191.77, 195.41, 195.55, 192.59, 197.43, 194.79, 195.85, 196.74, 196.01, 198.45, 200.18, 199.73, 195.45, 196.46, 193.90, 193.60, 192.90, 192.87, 188.01, 188.12, 191.63, 192.13, 195.00, 198.47, 197.79, 199.41, 201.21, 201.33, 201.52, 200.25, 199.29, 202.35, 203.27, 203.37, 203.11, 201.85, 205.26, 207.51, 207.00, 206.60, 208.95, 208.83, 207.93, 210.39, 211.00, 210.36, 210.15, 210.04, 208.08, 208.56, 207.74, 204.84, 202.54, 205.62, 205.47, 208.73, 208.55, 209.31, 209.07, 209.35, 209.32, 209.56, 208.69, 210.68, 208.53, 205.61, 209.62, 208.35, 206.95, 205.34, 205.87, 201.88, 202.90, 205.03, 208.03, 204.86, 200.02, 201.67, 203.50, 206.02, 205.68, 205.21, 207.40, 205.93, 203.87, 201.02, 201.36, 198.82, 194.05, 191.92, 192.11, 193.66, 188.83, 191.93, 187.81, 188.06, 185.65, 186.69, 190.52, 187.64, 190.20, 188.13, 189.11, 193.72, 193.65, 190.16, 191.30, 191.60, 187.95, 185.42, 185.43, 185.27, 182.86, 186.63, 189.78, 192.88, 192.09, 192.00, 194.78, 192.32, 193.20, 195.54, 195.09, 193.56, 198.11, 199.00, 199.78, 200.43, 200.59, 198.40, 199.38, 199.54, 202.76, 202.50, 202.17, 203.34, 204.63, 204.38, 204.67, 204.56, 203.21, 203.12, 203.24, 205.12, 206.02, 205.52, 206.92, 206.25, 204.19, 206.42, 203.95, 204.50, 204.02, 205.92, 208.00, 208.01, 207.78, 209.24, 209.90, 210.10, 208.97, 208.97, 208.61, 208.92, 209.35, 207.45, 206.33, 207.97, 206.16, 205.01, 204.97, 205.72, 205.89, 208.45, 206.50, 206.56, 204.76, 206.78, 204.85, 204.91, 204.20, 205.49, 205.21, 207.87, 209.28, 209.34, 210.24, 209.84, 210.27, 210.91, 210.28, 211.35, 211.68, 212.37, 212.08, 210.07, 208.45, 208.04, 207.75, 208.37, 206.52, 207.85, 208.44, 208.10, 210.81, 203.24, 199.60, 203.20, 206.66, 209.48, 209.92, 208.41, 209.66, 209.53, 212.65, 213.40, 214.95, 214.92, 216.12, 215.83}
}

func StockAnalysis() *image.Image {
	xv, yv := xvalues(), yvalues()
	priceSeries := chart.TimeSeries{
		Name: "SPY",
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: xv,
		YValues: yv,
	}

	smaSeries := chart.SMASeries{
		Name: "SPY - SMA",
		Style: chart.Style{
			StrokeColor:     drawing.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: priceSeries,
	}

	bbSeries := &chart.BollingerBandsSeries{
		Name: "SPY - Bol. Bands",
		Style: chart.Style{
			StrokeColor: drawing.ColorFromHex("efefef"),
			FillColor:   drawing.ColorFromHex("efefef").WithAlpha(64),
		},
		InnerSeries: priceSeries,
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			TickPosition: chart.TickPositionBetweenTicks,
		},
		YAxis: chart.YAxis{
			Range: &chart.ContinuousRange{
				Max: 220.0,
				Min: 180.0,
			},
		},
		Series: []chart.Series{
			bbSeries,
			priceSeries,
			smaSeries,
		},
	}
	collector := &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	image, err := collector.Image()

	if err != nil {
		panic(err)
	}
	return &image
}

func StackedBar() *image.Image {
	sbc := chart.StackedBarChart{
		Title: "Test Stacked Bar Chart",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height: 512,
		Bars: []chart.StackedBar{
			{
				Name: "This is a very long string to test word break wrapping.",
				Values: []chart.Value{
					{Value: 5, Label: "Blue"},
					{Value: 5, Label: "Green"},
					{Value: 4, Label: "Gray"},
					{Value: 3, Label: "Orange"},
					{Value: 3, Label: "Test"},
					{Value: 2, Label: "??"},
					{Value: 1, Label: "!!"},
				},
			},
			{
				Name: "Test",
				Values: []chart.Value{
					{Value: 10, Label: "Blue"},
					{Value: 5, Label: "Green"},
					{Value: 1, Label: "Gray"},
				},
			},
			{
				Name: "Test 2",
				Values: []chart.Value{
					{Value: 10, Label: "Blue"},
					{Value: 6, Label: "Green"},
					{Value: 4, Label: "Gray"},
				},
			},
		},
	}
	collector := &chart.ImageWriter{}
	sbc.Render(chart.PNG, collector)

	image, err := collector.Image()

	if err != nil {
		panic(err)
	}
	return &image
}

func Annotation() *image.Image {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
			chart.AnnotationSeries{
				Annotations: []chart.Value2{
					{XValue: 1.0, YValue: 1.0, Label: "One"},
					{XValue: 2.0, YValue: 2.0, Label: "Two"},
					{XValue: 3.0, YValue: 3.0, Label: "Three"},
					{XValue: 4.0, YValue: 4.0, Label: "Four"},
					{XValue: 5.0, YValue: 5.0, Label: "Five"},
				},
			},
		},
	}

	collector := &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	image, err := collector.Image()

	if err != nil {
		panic(err)
	}
	return &image
}

func BarChart() *image.Image {
	graph := chart.BarChart{
		Title: "Test Bar Chart",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars: []chart.Value{
			{Value: 5.25, Label: "Blue"},
			{Value: 4.88, Label: "Green"},
			{Value: 4.74, Label: "Gray"},
			{Value: 3.22, Label: "Orange"},
			{Value: 3, Label: "Test"},
			{Value: 2.27, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	collector := &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	image, err := collector.Image()

	if err != nil {
		panic(err)
	}
	return &image
}

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func BenchmarkLine() *image.Image {

	numValues := 100
	numSeries := 20
	series := make([]chart.Series, numSeries)

	for i := 0; i < numSeries; i++ {
		xValues := make([]time.Time, numValues)
		yValues := make([]float64, numValues)

		for j := 0; j < numValues; j++ {
			xValues[j] = time.Now().AddDate(0, 0, (numValues-j)*-1)
			yValues[j] = random(float64(-500), float64(500))
		}

		series[i] = chart.TimeSeries{
			Name:    fmt.Sprintf("aaa.bbb.hostname-%v.ccc.ddd.eee.fff.ggg.hhh.iii.jjj.kkk.lll.mmm.nnn.value", i),
			XValues: xValues,
			YValues: yValues,
		}
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Time",
		},
		YAxis: chart.YAxis{
			Name: "Value",
		},
		Series: series,
	}

	collector := &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	image, err := collector.Image()

	if err != nil {
		panic(err)
	}
	return &image
}

func LinearRegression() *image.Image {

	mainSeries := chart.ContinuousSeries{
		Name:    "A test series",
		XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(),        //generates a []float64 from 1.0 to 100.0 in 1.0 step increments, or 100 elements.
		YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(0).WithMax(100)}.Values(), //generates a []float64 randomly from 0 to 100 with 100 elements.
	}

	// note we create a LinearRegressionSeries series by assignin the inner series.
	// we need to use a reference because `.Render()` needs to modify state within the series.
	linRegSeries := &chart.LinearRegressionSeries{
		InnerSeries: mainSeries,
	} // we can optionally set the `WindowSize` property which alters how the moving average is calculated.

	graph := chart.Chart{
		Series: []chart.Series{
			mainSeries,
			linRegSeries,
		},
	}
	collector := &chart.ImageWriter{}

	graph.Render(chart.PNG, collector)
	image, err := collector.Image()

	if err != nil {
		panic(err)
	}
	return &image
}

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	LIQUID = 1
	BOUND  = 3
	MASS   = 20.00e-13
	DENS   = 1000.0
)

func calcDelta(params map[string]float64) float64 {
	x := []float64{1, 1, 0, -1, -1, -1, 0, 1, 1, 1, 0, -1, -1, -1, 0, 1, 1, 1, 0, -1, -1, -1, 0, 1, 2, -2, 0, 0, 0, 0, 0, 0}
	y := []float64{0, 1, 1, 1, 0, -1, -1, -1, 0, 1, 1, 1, 0, -1, -1, -1, 0, 1, 1, 1, 0, -1, -1, -1, 0, 0, 2, -2, 0, 0, 0, 0}
	z := []float64{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, 0, 0, 0, 0, 2, -2, 1, -1}
	sum1X := 0.0
	sum1Y := 0.0
	sum1Z := 0.0
	sum1 := 0.0
	sum2 := 0.0
	vX := 0.0
	vY := 0.0
	vZ := 0.0
	dist := 0.0
	particleRadius := math.Pow(params["mass"]/params["rho0"], 1.0/3.0)
	hR2 := 0.0
	for i := 0; i < len(x); i++ {
		vX = x[i] * 0.8 * particleRadius
		vY = y[i] * 0.8 * particleRadius
		vZ = z[i] * 0.8 * particleRadius
		dist = math.Sqrt(vX*vX + vY*vY + vZ*vZ)
		if dist <= params["h"]*params["simulation_scale"] {
			hR2 = math.Pow((params["h"]*params["simulation_scale"] - dist), 2)
			sum1X += hR2 * vX / dist
			sum1Y += hR2 * vY / dist
			sum1Z += hR2 * vZ / dist

			sum2 += hR2 * hR2
		}
	}

	sum1 = sum1X*sum1X + sum1Y*sum1Y + sum1Z*sum1Z
	return 1.0 / (params["beta"] * params["grad_wspiky_coefficient"] * params["grad_wspiky_coefficient"] * (sum1 + sum2))
}

func genParamsSibernetic(h, xDim, yDim, zDim, mass float64) map[string]float64 {
	param := map[string]float64{}

	param["particles"] = 30
	param["h"] = h
	param["x_max"] = h * xDim
	param["x_min"] = 0
	param["y_max"] = h * yDim
	param["y_min"] = 0
	param["z_max"] = h * zDim
	param["z_min"] = 0
	param["mass"] = mass
	param["rho0"] = 1000.0
	param["time_step"] = 4.0 * 5.0e-06
	param["simulation_scale"] = 0.0037 * math.Pow(mass, 1.0/3.0) / math.Pow(0.00025, 1.0/3.0)

	param["beta"] = math.Pow(param["time_step"], 2) * math.Pow(param["mass"], 2) * 2 / math.Pow(param["rho0"], 2)
	param["wpoly6_coefficient"] = 315.0 / (64.0 * math.Pi * math.Pow(h*param["simulation_scale"], 9.0))
	param["grad_wspiky_coefficient"] = -45.0 / (math.Pi * math.Pow((h*param["simulation_scale"]), 6.0))
	param["divgrad_wviscosity_coefficient"] = -param["grad_wspiky_coefficient"]
	param["mass_mult_wpoly6_coefficient"] = param["mass"] * param["wpoly6_coefficient"]
	param["mass_mult_grad_wspiky_coefficient"] = param["mass"] * param["grad_wspiky_coefficient"]
	param["mass_mult_divgrad_viscosity_coefficient"] = param["mass"] * param["divgrad_wviscosity_coefficient"]
	param["surf_tens_coeff"] = param["mass_mult_wpoly6_coefficient"] * param["simulation_scale"]
	param["delta"] = calcDelta(param)

	param["gravity_x"] = 0.0
	param["gravity_y"] = -9.8
	param["gravity_z"] = 0.0
	param["mu"] = 0.1 * 0.00004
	return param
}

func genModel(xDim, yDim, zDim float64, fileName string) {
	pCount := 0
	h := 3.34
	r0 := h * 0.5
	param := genParamsSibernetic(h, xDim, yDim, zDim, MASS)
	model := drawBounds(xDim, yDim, zDim, h, r0)
	for x := 4 * r0; x < h*xDim*5/7-4*r0; x += r0 {
		for y := 4 * r0; y < h*yDim-4*r0; y += r0 {
			for z := 4 * r0; z < h*zDim-4*r0; z += r0 {
				model = append(model, []float64{x, y, z, float64(LIQUID), 0.0, 0.0, 0.0, 1.0})
				pCount++
			}
		}
	}
	for len(model)%8 != 0 {
		model = model[:len(model)-1]
	}
	fmt.Printf("model has - %d particles", len(model))
	if fileName == "" {
		fileName = fmt.Sprintf("%dP.txt", len(model))
	}
	putToFile(param, model, fileName)

}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func putToFile(param map[string]float64, model [][]float64, fileName string) {
	f, err := os.Create(fileName)
	check(err)
	defer f.Close()
	_, err = f.Write([]byte("parameters[\n"))
	check(err)
	for k, v := range param {
		_, err = f.Write([]byte(fmt.Sprintf("%s: %.20e\n", k, v)))
		check(err)
	}
	_, err = f.Write([]byte("]\n"))
	check(err)
	_, err = f.Write([]byte("model[\n"))
	check(err)
	_, err = f.Write([]byte("position[\n"))
	check(err)
	for _, p := range model {
		_, err = f.Write([]byte(fmt.Sprintf("%f\t%f\t%f\t%f\n", p[0], p[1], p[2], p[3]+0.1)))
		check(err)
	}
	_, err = f.Write([]byte("]\n"))
	check(err)
	_, err = f.Write([]byte("velocity[\n"))
	check(err)
	for _, p := range model {
		_, err = f.Write([]byte(fmt.Sprintf("%f\t%f\t%f\t%f\n", p[4], p[5], p[6], p[3]+0.1)))
		check(err)
	}
	_, err = f.Write([]byte("]\n"))
	check(err)
	_, err = f.Write([]byte("connection[\n"))
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("]\n"))
	check(err)
	_, err = f.Write([]byte("membranes[\n"))
	check(err)
	_, err = f.Write([]byte("]\n"))
	check(err)
	_, err = f.Write([]byte("particleMemIndex[\n"))
	check(err)
	_, err = f.Write([]byte("]\n]\n"))
	check(err)
}

func main() {
	genModel(8, 8, 8, "")
	//genModel(192, 120, 120, "")
	//genModel(384, 192, 192, "")
}

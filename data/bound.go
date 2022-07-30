package main

import "math"

func drawBounds(xDim, yDim, zDim, h, r0 float64) [][]float64 {
	nx := int(xDim * h / r0) // Numbers of boundary particles on X-axis
	ny := int(yDim * h / r0) // Numbers of boundary particles on Y-axis
	nz := int(zDim * h / r0) // Numbers of boundary particles on Z-axis
	pp := [][]float64{}
	// 1 - top and bottom
	for ix := 0; ix < nx; ix++ {
		for iy := 0; iy < ny; iy++ {
			if ((ix == 0) || (ix == nx-1)) || ((iy == 0) || (iy == ny-1)) {
				if ((ix == 0) || (ix == nx-1)) && ((iy == 0) || (iy == ny-1)) { // corners
					x := float64(ix)*r0 + r0/2.0
					y := float64(iy)*r0 + r0/2.0
					z := 0.0*r0 + r0/2.0
					coefX1, coefX2 := 0.0, 0.0
					if ix == 0 {
						coefX1 = 1.0
					}
					if ix == nx-1 {
						coefX2 = 1.0
					}
					coefY1, coefY2 := 0.0, 0.0
					if iy == 0 {
						coefY1 = 1.0
					}
					if iy == ny-1 {
						coefY2 = 1.0
					}
					vel_x := (1.0*coefX1 - 1.0*coefX2) / math.Sqrt(3.0)

					vel_y := (1.0*coefY1 - 1.0*coefY2) / math.Sqrt(3.0)
					vel_z := 1.0 / math.Sqrt(3.0)
					pp = append(pp, []float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.0})

					x = float64(ix)*r0 + r0/2.0
					y = float64(iy)*r0 + r0/2.0
					z = (float64(nz)-1.0)*r0 + r0/2.0
					vel_x = (1.0*coefX1 - 1.0*coefX2) / math.Sqrt(3.0)
					vel_y = (1.0*coefY1 - 1.0*coefY2) / math.Sqrt(3.0)
					vel_z = -1.0 / math.Sqrt(3.0)
					pp = append(pp, []float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.0})

				} else { // edges
					x := float64(ix)*r0 + r0/2.0
					y := float64(iy)*r0 + r0/2.0
					z := 0.0*r0 + r0/2.0
					coefX1, coefX2 := 0.0, 0.0
					if ix == 0 {
						coefX1 = 1.0
					}
					if ix == nx-1 {
						coefX2 = 1.0
					}
					coefY1, coefY2 := 0.0, 0.0
					if iy == 0 {
						coefY1 = 1.0
					}
					if iy == ny-1 {
						coefY2 = 1.0
					}
					vel_x := (1.0*coefX1 -
						coefX2) / math.Sqrt(2.0)
					vel_y := (1.0*coefY1 -
						coefY2) / math.Sqrt(2.0)
					vel_z := 1.0 / math.Sqrt(2.0)
					pp = append(
						pp,
						[]float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.},
					)
					x = float64(ix)*r0 + r0/2.0
					y = float64(iy)*r0 + r0/2.0
					z = (float64(nz)-1.0)*r0 + r0/2.0
					vel_x = (1.0*coefX1 -
						coefX2) / math.Sqrt(2.0)
					vel_y = (1.0*coefY1 - coefY2) / math.Sqrt(2.0)
					vel_z = -1.0 / math.Sqrt(2.0)
					pp = append(
						pp,
						[]float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.0},
					)
				}
			} else { // planes
				x := float64(ix)*r0 + r0/2.0
				y := float64(iy)*r0 + r0/2.0
				z := 0.0*r0 + r0/2.0
				vel_x := 0.0
				vel_y := 0.0
				vel_z := 1.0
				pp = append(
					pp,
					[]float64{x, y, z, 1.0, float64(BOUND), vel_x, vel_y, vel_z, 1.0},
				)

				z = (float64(nz)-1.0)*r0 + r0/2.0
				vel_x = 0.0
				vel_y = 0.0
				vel_z = -1.0
				pp = append(
					pp,
					[]float64{x, y, z, 1.0, float64(BOUND), vel_x, vel_y, vel_z, 1.0},
				)
			}
		}
	}
	// 2 - side walls OX-OZ and opposite
	for ix := 0; ix < nx; ix++ {
		for iz := 1; iz < nz-1; iz++ {
			if (ix == 0) && (ix == nx-1) {
				x := float64(ix)*r0 + r0/2.0
				y := 0.0*r0 + r0/2.0
				z := float64(iz)*r0 + r0/2.0
				var coefX1, coefX2, coefZ1, coefZ2 float64
				if ix == 0 {
					coefX1 = 1.0
				} else if ix == nx-1 {
					coefX2 = 1.0
				}
				if iz == 0 {
					coefZ1 = 1.0
				} else if iz == nz-1 {
					coefZ2 = 1.0
				}

				vel_x := (1.0 * (coefX1 - coefX2)) / math.Sqrt(2.0)
				vel_y := 1.0 / math.Sqrt(2.0)
				vel_z := 1.0 * (coefZ1 - coefZ2) / math.Sqrt(2.0)
				pp = append(
					pp,
					[]float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.0},
				)

				y = float64(ny-1)*r0 + r0/2.0
				vel_x = (1.0 * (coefX1 - coefX2)) / math.Sqrt(2.0)
				vel_y = -1.0 / math.Sqrt(2.0)
				vel_z = 1.0 * (coefZ1 - coefZ2) / math.Sqrt(2.0)

				pp = append(
					pp,
					[]float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.0},
				)
			} else {
				x := float64(ix)*r0 + r0/2.0
				y := 0.0*r0 + r0/2.0
				z := float64(iz)*r0 + r0/2.0
				vel_x := 0.0
				vel_y := 1.0
				vel_z := 0.0
				pp = append(
					pp,
					[]float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.0},
				)
				x = float64(ix)*r0 + r0/2.0
				y = (float64(ny)-1)*r0 + r0/2.0
				z = float64(iz)*r0 + r0/2.0
				vel_x = 0.0
				vel_y = -1.0
				vel_z = 0.0
				pp = append(
					pp,
					[]float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.0},
				)
			}
		}
	}

	// 3 - side walls OY-OZ and opposite
	for iy := 1; iy < ny-1; iy++ {
		for iz := 1; iz < nz-1; iz++ {
			x := 0.0*r0 + r0/2.0
			y := float64(iy)*r0 + r0/2.0
			z := float64(iz)*r0 + r0/2.0
			vel_x := 1.0
			vel_y := 0.0
			vel_z := 0.0
			pp = append(
				pp,
				[]float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.0},
			)
			x = (float64(nx)-1)*r0 + r0/2.0
			y = float64(iy)*r0 + r0/2.0
			z = float64(iz)*r0 + r0/2.0
			vel_x = -1.0
			vel_y = 0.0
			vel_z = 0.0

			pp = append(
				pp,
				[]float64{x, y, z, float64(BOUND), vel_x, vel_y, vel_z, 1.0},
			)
		}
	}
	return pp
}

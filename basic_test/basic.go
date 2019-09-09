package main

import "fmt"
import "math"
import "strconv"

type Point struct {
	x float64
	y float64
}

type Detail interface {
	PrintInfo()
}

func (point *Point) PrintInfo() {
	fmt.Printf("x:%4.2f y:%4.2f\n", point.x, point.y)
}

type TmpInfo struct {
	length float64
	sin    float64
	cos    float64
}

func distance(start *Point, end *Point) (length float64, sin float64, cos float64) {
	total := math.Pow((end.x-start.x), 2.0) + math.Pow((end.y-start.y), 2.0)
	length = math.Sqrt(total)
	sin = (end.y - start.y) / length
	cos = (end.x - start.x) / length
	return length, sin, cos
}

func main() {

	var total float64 = 0.0
	var avg float64 = 0.0
	var need float64 = 0.0
	avg_num := 5.0
	points := []Point{}
	result := []Point{}
	point_dict := map[int]TmpInfo{}

	point1 := Point{x: 0.0, y: 0.0}
	point2 := Point{x: 0.0, y: 1.0}
	point3 := Point{x: 1.0, y: 1.0}
	point4 := Point{x: 1.0, y: 0.0}

	points = append(points, point1)
	points = append(points, point2)
	points = append(points, point3)
	points = append(points, point4)

	points_length := len(points)

	for index := 0; index < points_length; index++ {
		var start Point
		var end Point
		if (index + 1) == points_length {
			start = points[index]
			end = points[0]

		} else {
			start = points[index]
			end = points[index+1]
		}
		length, sin, cos := distance(&start, &end)
		total = total + length
		point_dict[index] = TmpInfo{length: length, sin: sin, cos: cos}
	}

	avg = total / avg_num
	fmt.Printf("%f\n", avg)
	fmt.Printf("%v\n", point_dict)
	need = avg
	index := 0
	start := points[index]
	var plength, psin, pcos float64
	tmpinfo := point_dict[index]
	plength, psin, pcos = tmpinfo.length, tmpinfo.sin, tmpinfo.cos
	// result = append(result, start)
	fmt.Printf("%v\n", tmpinfo)
	fmt.Printf("length:%f, need:%f\n", plength, need)

	for {
		if plength >= need {
			var p_x, p_y float64
			if pcos != 0 {
				p_x = start.x + pcos*need
			} else {
				p_x = start.x
			}
			if psin != 0 {
				p_y = start.y + psin*need
			} else {
				p_y = start.y
			}
			r := Point{x: p_x, y: p_y}
			result = append(result, r)
			fmt.Println("#####")
			start.PrintInfo()
			r.PrintInfo()
			fmt.Println("#####")
			plength -= need
			need = avg
			p_str := fmt.Sprintf("%4.4f", plength)
			plength, _ = strconv.ParseFloat(p_str, 64)
			start = r

		} else {
			need = need - plength
			index = index + 1
			if index > (points_length - 1) {
				break
			}
			tmpinfo = point_dict[index]
			start = points[index]
			plength, psin, pcos = tmpinfo.length, tmpinfo.sin, tmpinfo.cos
			fmt.Printf("%v\n", tmpinfo)
		}
	}

	for _, p := range result {
		p.PrintInfo()
	}
}

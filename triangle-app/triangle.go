package trianglen

//Comes to use only if triangle possible, direct invocation utilizes base test
func CheckTriangleType(a float64, b float64, c float64) (string) {

	if !IsTriangle(a, b, c) {
		return "Impossible"
	}

	if IsEquilateral(a, b, c) {
		return "Equilateral"
	}

	if IsIsoceles(a, b, c) {
		return "Isoceles"
	}

	return "Scalene"
}

//Base Test to determine even if there is triangle possible

func IsTriangle(a float64, b float64, c float64) bool {

	if a <= 0 || b <=0 || c <= 0 {
		return false
	}

	if a+b <= c || a+c <= b || b+c <= a {
		return false
	}

	return true
}


func IsEquilateral(a float64, b float64, c float64) bool {

	if !IsTriangle(a,b,c) {
		return false
	}

	if a==b && b==c {
		return true
	}

	return false
}

func IsIsoceles(a float64, b float64, c float64) bool {

	if !IsTriangle(a,b,c) {
		return false
	}

	if a==b || b==c || c==a {
		return true
	}

	return false
}



package closure

func main() {
	kf := getKeyFunc()
	kf2 := kf
	println(kf2())
	println(kf())
}

type keyFunc func() int

func getKeyFunc() keyFunc {
	k := 0
	return func() int {
		k++
		return k
	}
}

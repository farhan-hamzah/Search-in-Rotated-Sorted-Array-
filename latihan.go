package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var n, i, target int
	var b bool
	b = false
	fmt.Scan(&n, &target)

	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	i = 0
	for i < n{
		if A[i] == target{
			b = true
			break
		}
		i++
	}
	if b == true{
		fmt.Print(i)
	}else{
		fmt.Print(-1)
	}
}
# perm

Permutations without recursion.


	func main() {
		p := perm.New([]int{11, 22, 33})

		for p.Next() {
			fmt.Println(p.Get())
		}
	}

package secret

import "fmt"

func Secret2() string {
	return Secret() + "s3cr31_2"
}

func _init() {
	fmt.Println("[init]-2", Secret2())
}

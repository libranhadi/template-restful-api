package helper


import(
	"fmt"
)
func PanicIfError(err error)  {
	if err != nil {
		fmt.Println("ERROR ", err)
		panic(err)
	}
}
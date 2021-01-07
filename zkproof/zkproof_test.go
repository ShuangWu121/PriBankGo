package zkproof

import(
	"fmt"
	"testing"

	//"../fields"
	"math/big"
   // "bytes"
    //"github.com/ShuangWu121/secp256k1"
)



func TestPedersenVectorComit(t *testing.T) {
	fmt.Println("Test Pedersen Vector Commitments")

	//define the public vector 
	b1 := big.NewInt(int64(3))
	b2:= big.NewInt(int64(144444))
	b3:= big.NewInt(int64(333))
	b4:= big.NewInt(int64(653))
	b5:= big.NewInt(int64(23))
	b6:= big.NewInt(int64(198))
	b7:= big.NewInt(int64(3))

	publicInputs := []*big.Int{b1,b2,b3,b4,b5,b6,b7}

	c:=PedersenVectorComit(publicInputs)
	fmt.Println(c.x,c.y)
}
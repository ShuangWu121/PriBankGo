package zkproof

import(
	"fmt"
	"testing"

	//"../fields"
	"math/big"
   // "bytes"
    "github.com/ShuangWu121/secp256k1"
    "github.com/stretchr/testify/assert"
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

	
	fmt.Println(publicInputs)
}

func TestZkPds_PubVec(t *testing.T) {
	

	fmt.Println("Test Zero-knowledge proof for product of Pedersen Commitments and public vectors ")

    //define the private vector
    a1 := big.NewInt(int64(3423))
	a2:= big.NewInt(int64(444))
	a3:= big.NewInt(int64(372))
	a4:= big.NewInt(int64(2))
	a5:= big.NewInt(int64(78643))
	a6:= big.NewInt(int64(34567898))
	a7:= big.NewInt(int64(13))

	privateInputs := []*big.Int{a1,a2,a3,a4,a5,a6,a7}

	//define the public vector 
	b1 := big.NewInt(int64(3))
	b2:= big.NewInt(int64(144444))
	b3:= big.NewInt(int64(333))
	b4:= big.NewInt(int64(653))
	b5:= big.NewInt(int64(23))
	b6:= big.NewInt(int64(198))
	b7:= big.NewInt(int64(3))

	publicInputs := []*big.Int{b1,b2,b3,b4,b5,b6,b7}

    //Generator G
	cv:=secp256k1.SECP256K1()
	G:=CurvePoint{cv.Params().Gx,cv.Params().Gy}
	fmt.Println("the base point is:",G)

 
    //Generator vector Hi
    hi:=Generators(len(privateInputs))
    fmt.Println("\nHi are:",hi)

    //Generate  Pedersen commitments for each element in a vector, r=234
    r:=big.NewInt(int64(234))
	pdsComits,err:=PedersenComitsForVector(privateInputs,hi,r,G)
	assert.Nil(t, err)


	fmt.Println("\nPedersen Commitments for private inputs:",pdsComits)

    H:=CurvePoint{cv.Params().Gx,cv.Params().Gy}

    gamma := big.NewInt(int64(6))
	alpha := big.NewInt(int64(1))
	beta := big.NewInt(int64(30))
	random_t := big.NewInt(int64(4))
    
	pf:=ZKproofPdsComits_PubVec(hi,publicInputs,gamma,alpha,beta,random_t,H)
	fmt.Println("\nc0.x",pf.c0.X.Text(16))
	fmt.Println("\nc0.y",pf.c0.Y.Text(16))

	fmt.Println("\nomega.x",pf.omega.X)
	fmt.Println("\nomega.y",pf.omega.Y)
	fmt.Println("\nd1.x",pf.d1.X.Text(16))
	fmt.Println("\nd1.y",pf.d1.Y)
	fmt.Println("\nd2.x",pf.d2.X)
	fmt.Println("\nd2.y",pf.d2.Y)
	fmt.Println("theta1",pf.theta1)
	fmt.Println("theta2",pf.theta2)
    
    check:=ZKverifyPdsComits_PubVec(hi,publicInputs,pf,H)

    fmt.Println("check:",check)

}
package main

import(
	"fmt"
    "github.com/ShuangWu121/PriBankGo/zkproof"
	"github.com/arnaucube/go-snark/fields"
	"math/big"
   // "bytes"
    "github.com/ShuangWu121/secp256k1"
    "github.com/ShuangWu121/PriBankGo/r1csqap"
)

func main(){
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
	G:=zkproof.CurvePoint{cv.Params().Gx,cv.Params().Gy}
	commitkey:= big.NewInt(int64(834323243))
	H,_:=zkproof.CurveScalarMult(G,commitkey)
	//fmt.Println("the base point is:",G)

 
    //Generator vector Hi
    hi:=zkproof.Generators(len(privateInputs))
    //fmt.Println("\nHi are:",hi)

    //Generate  Pedersen commitments for each element in a vector, r=234
    //r:=big.NewInt(int64(234))
	//pdsComits,_:=zkproof.PedersenComitsForVector(privateInputs,hi,r,G)



	//fmt.Println("\nPedersen Commitments for private inputs:",pdsComits)

    gamma := big.NewInt(int64(6))
	random_t := big.NewInt(int64(4))
    
	pf:=zkproof.ZKproofPdsComits_PubVec(hi,publicInputs,gamma,random_t,H)
    check:=zkproof.ZKverifyPdsComits_PubVec(hi,publicInputs,pf,H)


   // zkproof.ZKproofPdsVec_PubVec(hi,privateInputs,publicInputs)
    fmt.Println("Padding result:",zkproof.Padding([]*big.Int{b1,b2,b3,b4,b5}))

    fmt.Println("check:",check)


    //test for bulletproof

    N, _ := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
    f := fields.NewFq(N)
    polyf := r1csqap.NewPolynomialField(f)

    a := []*big.Int{a1,a2,a3,a4,a5,a6,a7}
    b := []*big.Int{b1,b2,b3,b4,b5,b6,b7}
    a=zkproof.Padding(a)
    b=zkproof.Padding(b)

    ab:=big.NewInt(int64(0))
    for i:=0;i<len(a);i++{
       temp:=polyf.F.Mul(a[i],b[i])
       ab=polyf.F.Add(temp,ab)
    }

    gi:=zkproof.Generators(len(a))
 
    ra:=big.NewInt(int64(234))
    ca:=zkproof.PedersenVectorComit(a,gi,H,ra)
    rab:=big.NewInt(int64(9484735))
    cab:=zkproof.PedersenComit(ab,rab,G,H)

    

    pf_bulletproof:=zkproof.ZKproofPdsVec_PubVec(gi,G,H,ca,cab,a,b,ra,rab,polyf)

    fmt.Println("bulletproof check : ",zkproof.ZKverifyPdsVec_PubVec(gi,G,H,ca,cab,b, polyf,pf_bulletproof))
    


}
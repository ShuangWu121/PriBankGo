package main

import(
	"fmt"
	"math/big"
	//"strings"
	//"encoding/binary"
	"os"
	"bufio"
    "bytes"
	"github.com/ShuangWu121/PriBankGo/circuitcompiler"
	"github.com/arnaucube/go-snark/fields"
	"github.com/ShuangWu121/PriBankGo/r1csqap"
	"encoding/gob"
	"github.com/ShuangWu121/PriBankGo/zkproof"
	"github.com/ShuangWu121/secp256k1"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/rand"
	//"unsafe"
    //"errors"
    //"github.com/ethereum/go-ethereum/crypto"
)


func EvalPolys(polyf r1csqap.PolynomialField,ux [][]*big.Int,x *big.Int)([]*big.Int){
	U:=[]*big.Int{}
	for i := 0; i < len(ux); i++ {
    	ui:=polyf.Eval(ux[i],x)
	    U=append(U,ui)
	}
	return U
}


func main(){
	fmt.Println("Generating QAP for Circuit")


	// parse the code
	circuitFile, _ := os.Open("cir.txt")


	// parse circuit code
	parser := circuitcompiler.NewParser(bufio.NewReader(circuitFile))
	circuit, err := parser.Parse()
	if err!=nil{fmt.Println("circuit parse wrong")}

	
	// code to R1CS
	fmt.Println("\nGenerating R1CS from code ...")
	u, v, w := circuit.GenerateR1CS()
	



	// Set Finite Field, speck256k1
	N, _ := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)

	f := fields.NewFq(N)
	// new Polynomial Field
	polyf := r1csqap.NewPolynomialField(f)
    

    // R1CS to QAP, compute the polynomial from the matrics

    fmt.Println("\ncompute QAP ...")

	ux, vx, wx, zx := polyf.R1CSToQAP(u, v, w)

    fmt.Println("\nGenerate private inputs")
    //these are the private witness
	b1 := big.NewInt(int64(100))
	b2 := big.NewInt(int64(34))
	b3 := big.NewInt(int64(2000))
	b4 := big.NewInt(int64(5))
	b1new := big.NewInt(int64(93))
	b2new := big.NewInt(int64(31))
	b3new := big.NewInt(int64(2010))
	b4new := big.NewInt(int64(5))

	v12 := big.NewInt(int64(7))
	v121:= big.NewInt(int64(0))
	v122:= big.NewInt(int64(0))
	v123:= big.NewInt(int64(0))
	v124:= big.NewInt(int64(0))
	v125:= big.NewInt(int64(0))
	v126:= big.NewInt(int64(1))
	v127:= big.NewInt(int64(1))
	v128:= big.NewInt(int64(1))

	v23:= big.NewInt(int64(10))
	v231:= big.NewInt(int64(0))
	v232:= big.NewInt(int64(1))
	v233:= big.NewInt(int64(0))
	v234:= big.NewInt(int64(1))
	v235:= big.NewInt(int64(0))

	v43 := big.NewInt(int64(0))


	privateInputs := []*big.Int{b1,b2,b3,b4,b1new,b2new,b3new,b4new,v12,v121,v122,v123,v124,v125,v126,v127,v128,v23,v231,v232,v233,v234,v235,v43}

	//public witness
	total := big.NewInt(int64(2139))


	publicSignals := []*big.Int{total}

    // wittness
	wires, _ := circuit.CalculateWitness(privateInputs, publicSignals)
	

	fmt.Println("\nThe number of wires is:",len(wires))
	fmt.Println("wires values:",wires)
	fmt.Println("\nsignals are :",circuit.Signals)

	//for i := 0; i < len(wires); i++ {
   // 	fmt.Println(circuit.Signals[i]+" is ",wires[i])
	//}	
    


    fmt.Println("\nR1CS is correct? (result is valid when the inputs are valid) ",r1csqap.Check_r1cs(wires,u,v,w,polyf))
    fmt.Println("\nQAP is correct? (result is valid when the inputs are valid) ",r1csqap.Check_QAP(wires,ux,vx,wx,u,v,w,polyf))


    
    

    //compute Ax=sum{a_iu_i(x)}, Bx=sum{a_iv_i(x)}, Cx=sum{a_iw_i(x)}
    // and P(x)=Ax*Bx-Cx  (which should be equal to hx*zx)
    // this end up with three polynomials
    
    Ax, Bx, Cx, px := polyf.CombinePolynomials(wires, ux, vx, wx)

    

    
    
   // fmt.Println("px:",px)

	fmt.Println("\ntest the correctness of the witnesses")
	hx := polyf.DivisorPolynomial(px, zx)
    //fmt.Println("\nhx is:",hx)

    // hx==px/zx so px==hx*zx
	buf1:=&bytes.Buffer{}
	gob.NewEncoder(buf1).Encode(px)

	buf2:=&bytes.Buffer{}
	gob.NewEncoder(buf2).Encode(polyf.Mul(hx,zx))

	fmt.Println("if px==hx*zx",bytes.Equal(buf1.Bytes(),buf2.Bytes()))

    
    //Prover commmit to all wires and coefficients of hx
    fmt.Println("\nProver commmits to all wires")


    //Generator G H
	cv:=secp256k1.SECP256K1()
	G:=zkproof.CurvePoint{cv.Params().Gx,cv.Params().Gy}
	H:=zkproof.CurvePoint{cv.Params().Gx,cv.Params().Gy}

    hi:=zkproof.Generators(len(wires))
    hi_hx:=zkproof.Generators(len(hx))

    gamma,_:=rand.Int(rand.Reader,N)
    c,_:=zkproof.PedersenComitsForVector(wires,hi,gamma,G)//commitments to wires
    ch,_:=zkproof.PedersenComitsForVector(hx,hi_hx,gamma,G)//commitments to coefficients of hx

    //compute challenge x
    buf:=&bytes.Buffer{}
	gob.NewEncoder(buf).Encode(append(append(ch,H),append(hi_hx,append(hi,append(c,G)...)...)...))


    x := new(big.Int).SetBytes(crypto.Keccak256(buf.Bytes()))
    x=f.Affine(x)
    
    fmt.Println("challenge is",x)

    //those are public information
    U:=EvalPolys(polyf,ux,x)//U=[u1,u2,...] u1=u1(x) evaluated at x
    V:=EvalPolys(polyf,vx,x)
    W:=EvalPolys(polyf,wx,x)
    z:=polyf.Eval(zx,x)


    //Compute X=[1*z(x),x*z(x),x^2*z(x),....] for computing h(x)z(x)

    X:=[]*big.Int{}
	for i := 0; i < len(hx); i++ {
		xi:=polyf.F.Mul(polyf.F.Exp(x,big.NewInt(int64(i))),z)
	    X=append(X,xi)
	}


 
    //Prover compute A=a_iu_i(x) B=a_iv_i(x) C=a_iw_i(x), A,B,C are single number


    A:=polyf.Eval(Ax,x)
    B:=polyf.Eval(Bx,x)
    C:=polyf.Eval(Cx,x)
    HZ:=polyf.F.Mul(polyf.Eval(hx,x),polyf.Eval(zx,x))

   // fmt.Println("A is", A)
    //fmt.Println("B is", B)
    //fmt.Println("C is", C)

    
    //Prover generates the proof
   
    //pfA is to prove ca=g^sum{a_iU[i]}
	random_tA,_ := rand.Int(rand.Reader,N)
    pfA:=zkproof.ZKproofPdsComits_PubVec(hi,U,gamma,random_tA,H)
    ca_p:=zkproof.PedersenComit(A,polyf.F.Neg(random_tA),G,H)
    
 
	random_tB,_ := rand.Int(rand.Reader,N)

    pfB:=zkproof.ZKproofPdsComits_PubVec(hi,V,gamma,random_tB,H)
    cb_p:=zkproof.PedersenComit(B,polyf.F.Neg(random_tB),G,H)
    

    
	random_tW,_ := rand.Int(rand.Reader,N)

    pfW:=zkproof.ZKproofPdsComits_PubVec(hi,W,gamma,random_tW,H)
    

    
	random_tH,_ := rand.Int(rand.Reader,N)

    pfH:=zkproof.ZKproofPdsComits_PubVec(hi_hx,X,gamma,random_tH,H)
    
    //pfA,pfB,pfW,pfH allow verifier compute ca,cb,cw,h(x)*z(x)


    
    rt:=polyf.F.Neg(polyf.F.Add(random_tW,random_tH))
    right:=zkproof.PedersenComit(polyf.F.Add(HZ,C),rt,G,H)


    pfProduct:=zkproof.ZkproofPdsProduct(ca_p,cb_p,right,G,H,A,B,polyf.F.Neg(random_tA),polyf.F.Neg(random_tB),rt,polyf)



    
    

    fmt.Println("\ncommitment size:",len(c)*64,"bytes")
    fmt.Println("proof size:704 bytes")

    /////////////////verifer check

    fmt.Println("\nVerification:")

    //Verifier computes ca
    fmt.Println("validation A:",zkproof.ZKverifyPdsComits_PubVec(hi,U,pfA,H))
    ca:=zkproof.CurvePointVecMult(c,U)
	ca,_=zkproof.CurveSub(ca,pfA.Omega)
	

    //Verifier computes cb
    fmt.Println("validation B:",zkproof.ZKverifyPdsComits_PubVec(hi,V,pfB,H))
    
	cb:=zkproof.CurvePointVecMult(c,V)
	cb,_=zkproof.CurveSub(cb,pfB.Omega)
	


    //Verifier computes cw
    fmt.Println("validation C:",zkproof.ZKverifyPdsComits_PubVec(hi,W,pfW,H))
	cw:=zkproof.CurvePointVecMult(c,W)
	cw,_=zkproof.CurveSub(cw,pfW.Omega)

    //Verifier computes commitment for hx*zx

    fmt.Println("validation HZ:",zkproof.ZKverifyPdsComits_PubVec(hi_hx,X,pfH,H))
	chz:=zkproof.CurvePointVecMult(ch,X)
	chz,_=zkproof.CurveSub(chz,pfH.Omega)
	

    //chcek is chz*cw is the product of ca cb

    c_right,_:=zkproof.CurveAdd(chz,cw)
    fmt.Println("product check: com(A*B)==com(C+HZ)",zkproof.ZkverifyPdsProduct(ca,cb,c_right,G,H,pfProduct,polyf))



    
    



	


}
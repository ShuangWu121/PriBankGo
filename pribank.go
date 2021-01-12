package main

import(
	"fmt"
	"math/big"
	"strings"
	//"os"
    "bytes"
	"github.com/arnaucube/go-snark/circuitcompiler"
	"github.com/arnaucube/go-snark/fields"
	"github.com/arnaucube/go-snark/r1csqap"
	"encoding/gob"
	"github.com/ShuangWu121/PriBankGo/zkproof"
	"github.com/ShuangWu121/secp256k1"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/rand"
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


	/* build the circuit for PriBank (simple version)
	 b1new=b1-v12
	 b2new=b2+v12
	 total=b1new+b2new
	 v12=v121*2+v122
	 v12*(v12-1)=0
	*/
	
	code := `
	func main(private b1, private b2,`+
		`private b1new, private b2new, private b3new,`+
		`private v12,private v121,private v122, private v123,`+
		`public total):
		s0 = b2+v12
		equals(s0,b2new)
		s1 = b1new +v12
		equals(s1, b1)`+
		// check v12, use variable z
		`z1=v121*2
		z2=z1+v122
		z3=z2+v123
		equals(z3,v12)`+
        //check range of v12... use variable zz
		`zz1=v121-1
		zz2=zz1*v121
		zz3=0+0
		equals(zz2,zz3)

		s2=b1new+b2new 
		s8=s2+b3new
		equals(s8,total)
		out = 1 * 1
	`
	fmt.Print("\nBuild the circuit:",code)

	// parse the code
	parser := circuitcompiler.NewParser(strings.NewReader(code))
	circuit, err := parser.Parse()
	if err!=nil{fmt.Println("circuit parse wrong")}

	
	// code to R1CS
	fmt.Println("\nGenerating R1CS from code, the matrics of u, v, w are: ...")
	u, v, w := circuit.GenerateR1CS()
	fmt.Println(u,v,w)


	// Set Finite Field, speck256k1
	N, _ := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)

	f := fields.NewFq(N)
	// new Polynomial Field
	polyf := r1csqap.NewPolynomialField(f)
    

    // R1CS to QAP, compute the polynomial from the matrics

	ux, vx, wx, zx := polyf.R1CSToQAP(u, v, w)

	fmt.Println("\nThe QAP for the circuit is u_i(x),v_i(x),w_i(x),z_i(x)...")

 
    //these are the private witness
	b1 := big.NewInt(int64(5))
	b2 := big.NewInt(int64(2))
	b1new := big.NewInt(int64(2))
	b2new := big.NewInt(int64(5))
	b3new := big.NewInt(int64(4))
	v12 := big.NewInt(int64(3))
	v121:= big.NewInt(int64(1))
	v122:= big.NewInt(int64(1))
	v123:= big.NewInt(int64(0))


	privateInputs := []*big.Int{b1,b2,b1new,b2new,b3new,v12,v121,v122,v123}

	//public witness
	total := big.NewInt(int64(11))
	publicSignals := []*big.Int{total}

    // wittness
	wires, err := circuit.CalculateWitness(privateInputs, publicSignals)
	if(err!=nil){fmt.Println("circuit inputs wrong")}

	fmt.Println("The whole wires of the circuit are:",wires)

    fmt.Println("Compute the polynomials with the witnesses, compute Ax=sum{a_iu_i(x)}, Bx=sum{a_iv_i(x)}, Cx=sum{a_iw_i(x)},P(x)=Ax*Bx-Cx")

    //compute Ax=sum{a_iu_i(x)}, Bx=sum{a_iv_i(x)}, Cx=sum{a_iw_i(x)}
    // and P(x)=Ax*Bx-Cx  (which should be equal to hx*zx)
    // this end up with three polynomials
    
    Ax, Bx, Cx, px := polyf.CombinePolynomials(wires, ux, vx, wx)
    

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
    fmt.Println("Prover commmits to all wires")


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

    fmt.Println("Prover compute A=a_iu_i(x) B=a_iv_i(x) C=a_iw_i(x), HZ=h(x)*z(x):")

    A:=polyf.Eval(Ax,x)
    B:=polyf.Eval(Bx,x)
    C:=polyf.Eval(Cx,x)
    HZ:=polyf.F.Mul(polyf.Eval(hx,x),polyf.Eval(zx,x))

    fmt.Println("A is", A)
    fmt.Println("B is", B)
    fmt.Println("C is", C)

    
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



    /////////////////verifer check

    //Verifier computes ca
    fmt.Println("validation A:",zkproof.ZKverifyPdsComits_PubVec(hi,U,pfA,H))
    ca:=zkproof.CurvePointVecMult(c,U)
	ca,_=zkproof.CurveSub(ca,pfA.Omega)
	

    //Verifier computes cb
    fmt.Println("validation B:",zkproof.ZKverifyPdsComits_PubVec(hi,V,pfB,H))
    
	cb:=zkproof.CurvePointVecMult(c,V)
	cb,_=zkproof.CurveSub(cb,pfB.Omega)
	


    //Verifier computes cw
    fmt.Println("validation W:",zkproof.ZKverifyPdsComits_PubVec(hi,W,pfW,H))
	cw:=zkproof.CurvePointVecMult(c,W)
	cw,_=zkproof.CurveSub(cw,pfW.Omega)

    //Verifier computes commitment for hx*zx

    fmt.Println("validation HZ:",zkproof.ZKverifyPdsComits_PubVec(hi_hx,X,pfH,H))
	chz:=zkproof.CurvePointVecMult(ch,X)
	chz,_=zkproof.CurveSub(chz,pfH.Omega)
	

    //chcek is chz*cw is the product of ca cb

    c_right,_:=zkproof.CurveAdd(chz,cw)
    fmt.Println("inner product check",zkproof.ZkverifyPdsProduct(ca,cb,c_right,G,H,pfProduct,polyf))



    
    



	


}
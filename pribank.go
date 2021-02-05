package main

import(
	"fmt"
	"math/big"
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
	"time"
	"runtime"
	"sync"

)


func EvalPolys(polyf r1csqap.PolynomialField,ux [][]*big.Int,x *big.Int)([]*big.Int){
	U:=[]*big.Int{}
	for i := 0; i < len(ux); i++ {
    	ui:=polyf.Eval(ux[i],x)
	    U=append(U,ui)
	}
	return U
}

func AddTxValueBits(txs []*big.Int,length int64) []*big.Int{
    bitslength:=fmt.Sprintf("0%db",length)
    var bits []*big.Int
    for j:=0;j<len(txs);j++{
    	input:=txs[j]
    	s := fmt.Sprintf("%"+bitslength, input) 

    	for i := 0; i < int(length); i++{

       		bits=append(bits,big.NewInt(int64(s[i]-'0')))

    	}
    }
   
    return bits
}

func circuit_init()(*circuitcompiler.Circuit){

	// parse the circuit 
	circuitFile, _ := os.Open("circuit.txt")



	parser := circuitcompiler.NewParser(bufio.NewReader(circuitFile))
	circuit, err := parser.Parse()
	if err!=nil{fmt.Println("circuit parse wrong")}
	return circuit

}

func seprate_witness(privateSignals,privateInputs,publicSignals []*big.Int)(){

}

func InputsGenerator(f fields.Fq)([]*big.Int,[]*big.Int,[]*big.Int){

	var privateSignals []*big.Int //include the private signals for range proof
    var privateInputs []*big.Int

    const users=2
    const balanceRange=6
    const TransactionsRange=6
    
    //generate the original balance
    //maxReceiveV:=f.Mul(f.Exp(big.NewInt(int64(2)),big.NewInt(int64(TransactionsRange))),big.NewInt(int64(users-1)))
    maxBalance:=f.Exp(big.NewInt(int64(2)),big.NewInt(int64(balanceRange-2)))
    total:=big.NewInt(int64(0))
    for i:=1;i<users+1;i++{
    	balance,_:=rand.Int(rand.Reader,maxBalance)
    	total=f.Add(total,balance)
    	privateSignals=append(privateSignals,balance)
    }

    fmt.Println("old balances are:",privateSignals)

    //generate transactions
    var tx [users+1][users+1]*big.Int
    for i:=0;i<users+1;i++{
    	for j:=0;j<users+1;j++{
    		tx[i][j]=big.NewInt(int64(0))
    	}
    }
    

    //user i make transaction to j, that stores in tx[i][j]
    var TxsArray []*big.Int
    var subResult []*big.Int
    for i:=1;i<users+1;i++{
    	sum:=big.NewInt(int64(0))
    	for j:=1;j<users+1;j++{
          if i!=j {
          	 var trans *big.Int
             if f.Sub(privateSignals[i-1],sum).Cmp(f.Sub(f.Exp(big.NewInt(int64(2)),big.NewInt(int64(TransactionsRange))),big.NewInt(int64(1))))<0 {
             	if f.Sub(privateSignals[i-1],sum).Cmp(big.NewInt(int64(0)))<0{
             		trans=big.NewInt(int64(0))}else{trans,_=rand.Int(rand.Reader,f.Sub(privateSignals[i-1],sum))}
             }else {
             	trans,_=rand.Int(rand.Reader,f.Sub(f.Exp(big.NewInt(int64(2)),big.NewInt(int64(TransactionsRange-3))),big.NewInt(int64(1))))
             }
             if new(big.Int).Mod(trans,big.NewInt(int64(2)))==big.NewInt(int64(1)){
             	tx[i][j]=big.NewInt(int64(0)) 
                trans=big.NewInt(int64(0))}else{tx[i][j]=trans}
             sum=f.Add(sum,trans)
             TxsArray=append(TxsArray,trans)
          }
        }
        subResult=append(subResult,f.Sub(privateSignals[i-1],sum))
    } 

    fmt.Println("transactions are:",tx)

    //compute new balances
    for j:=1;j<users+1;j++{
    	for i:=1;i<users+1;i++{
    		subResult[j-1]=f.Add(subResult[j-1],tx[i][j])
    	}
    }
    fmt.Println("new balances are:",subResult)
    privateSignals=append(privateSignals,subResult...)

    //generate randomness t
    var d []*big.Int
    var t []*big.Int
    N, _ := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
    for i:=1;i<users+1;i++{
    	temp,_:=rand.Int(rand.Reader,N)
    	t=append(t,temp)
    	d=append(d,f.Add(subResult[i-1],temp))

    }
    privateSignals=append(privateSignals,t...)

    privateSignals=append(privateSignals,TxsArray...)

    privateInputs=append(privateInputs,privateSignals...)



    privateSignals=append(privateSignals,AddTxValueBits(subResult,balanceRange)...)

    //fmt.Println("add new balances bits",privateSignals)
    privateSignals=append(privateSignals,AddTxValueBits(TxsArray,TransactionsRange)...)
    //fmt.Println("private signals length:",len(privateSignals))

    //public inputs

    
    var publicSignals []*big.Int
    publicSignals=append(publicSignals,total)
    publicSignals=append(publicSignals,d...)
   // fmt.Println("publicSignals length",t)

     

    
   /*
	b1 := big.NewInt(int64(100))
	b2 := big.NewInt(int64(34))
	b3 := big.NewInt(int64(200))
	b4 := big.NewInt(int64(5))
	b1new := big.NewInt(int64(93))
	b2new := big.NewInt(int64(31))
	b3new := big.NewInt(int64(210))
	b4new := big.NewInt(int64(5))
	t1 := f.Affine(big.NewInt(int64(388695)))
	t2 := f.Affine(big.NewInt(int64(3433335)))
	t3 := f.Affine(big.NewInt(int64(344565)))
	t4 := f.Affine(big.NewInt(int64(323455665)))
    privateSignals=append(privateSignals,[]*big.Int{b1,b2,b3,b4,b1new,b2new,b3new,b4new,t1,t2,t3,t4}...)
    

	v12 := big.NewInt(int64(7))
    v13 := big.NewInt(int64(0))
    v14 := big.NewInt(int64(0))
    
    v21 := big.NewInt(int64(0))
    v23:= big.NewInt(int64(10))
    v24:= big.NewInt(int64(0))

    v31:= big.NewInt(int64(0))
    v32:= big.NewInt(int64(0))
    v34:= big.NewInt(int64(0))

    v41:= big.NewInt(int64(0))
    v42:= big.NewInt(int64(0))
    v43 := big.NewInt(int64(0))
    Txs:=[]*big.Int{v12,v13,v14,v21,v23,v24,v31,v32,v34,v41,v42,v43}

    privateSignals=append(privateSignals,Txs...)
    privateInputs=append(privateInputs,privateSignals...)
    privateSignals=append(privateSignals,AddTxValueBits([]*big.Int{b1new,b2new,b3new,b4new})...)
    privateSignals=append(privateSignals,AddTxValueBits(Txs)...)

    fmt.Println("private signals length:",len(privateSignals))

    //public inputs
	total := big.NewInt(int64(339))
	d1:=f.Add(t1,b1new)
	d2:=f.Add(t2,b2new)
	d3:=f.Add(t3,b3new)
	d4:=f.Add(t4,b4new)
    publicSignals := []*big.Int{total,d1,d2,d3,d4}*/

    return privateInputs,privateSignals,publicSignals

}


func main(){


	fmt.Println("Generating QAP for Circuit")


	circuit:=circuit_init()

	
	// code to R1CS
	fmt.Println("\nGenerating R1CS from code ...")
	u, v, w := circuit.GenerateR1CS()
	



	// Set Finite Field, order of speck256k1, and Polynomial Field
	N, _ := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
	f := fields.NewFq(N)
	polyf := r1csqap.NewPolynomialField(f)


    // generate private inputs
    fmt.Println("\nGenerate private inputs")

	//privateSignals include range proof bits.
	privateInputs,privateSignals,publicSignals:=InputsGenerator(f)
    
	fmt.Println("privateInputs are:",privateInputs)
    fmt.Println("\nPrivate inputs length:",len(privateInputs))
	
	
    // all wires in the circuit
	wires, _ := circuit.CalculateWitness(privateSignals, publicSignals)


	

	fmt.Println("\nThe number of wires is:",len(wires))
	fmt.Println("wires values:",wires)

	reader1 := bufio.NewReader(os.Stdin)
    fmt.Print("see wires values ")
    press, _ := reader1.ReadString('\n')
    fmt.Print("You press " + press)

	fmt.Println("\nsignals are :",circuit.Signals)

    

    // R1CS to QAP, compute the polynomials 
    fmt.Println("\ncompute QAP ...")

	ux, vx, wx, zx := polyf.R1CSToQAP(u, v, w)




    

    
    


   // fmt.Println("\nR1CS is correct? (result is valid when the inputs are valid) ",r1csqap.Check_r1cs(wires,u,v,w,polyf))
    //fmt.Println("\nQAP is correct? (result is valid when the inputs are valid) ",r1csqap.Check_QAP(wires,ux,vx,wx,u,v,w,polyf))


    
    

    //compute Ax=sum{a_iu_i(x)}, Bx=sum{a_iv_i(x)}, Cx=sum{a_iw_i(x)}
    // and P(x)=Ax*Bx-Cx  (which should be equal to hx*zx)
    // this end up with three polynomials

    proveTime:=time.Now()
    
    var wg sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())
    
    Ax, Bx, Cx, px := polyf.CombinePolynomials(wires, ux, vx, wx)

  

	fmt.Println("\ntest the correctness of the witnesses")
	hx := polyf.DivisorPolynomial(px, zx)

    // hx==px/zx so px==hx*zx
	buf1:=&bytes.Buffer{}
	gob.NewEncoder(buf1).Encode(px)

	buf2:=&bytes.Buffer{}
	gob.NewEncoder(buf2).Encode(polyf.Mul(hx,zx))

	fmt.Println("if px==hx*zx",bytes.Equal(buf1.Bytes(),buf2.Bytes()))

    
    //Prover commmit to all wires
    fmt.Println("\nProver commmits ")


    //Initiate Elliptic curve
	cv:=secp256k1.SECP256K1()
	G:=zkproof.CurvePoint{cv.Params().Gx,cv.Params().Gy}
	H:=zkproof.CurvePoint{cv.Params().Gx,cv.Params().Gy}

    

    //Prover commits
    //seprate the inner wires and input wires
    

    wires_public:=wires[0:len(publicSignals)+1]
    fmt.Println("public wires:",wires_public)
    
    wires_input:=wires[len(publicSignals)+1:(len(privateInputs)+len(publicSignals)+1)]  //only include balances transactions and t
    fmt.Println("privateInputs are :", wires_input)
    
    //commit to the inputs
    gamma,_:=rand.Int(rand.Reader,N)
    hi:=zkproof.Generators(len(wires_input))
    c_inputs,_:=zkproof.PedersenComitsForVector(wires_input,hi,gamma,G)
    
    //commit to the inner wires and hx
   

    pos_inner:=len(publicSignals)+len(privateInputs)+1
    wires_inner_len:=len(wires[pos_inner:])
    wires_inner:=zkproof.Padding(wires[pos_inner:])
   
    

    r_inner,_:=rand.Int(rand.Reader,N)
    gi:=zkproof.Generators(len(wires_inner))

    commit_inner:=zkproof.PedersenVectorComit(wires_inner,gi,H,r_inner)
   
    
    r_hx,_:=rand.Int(rand.Reader,N)
    hx=zkproof.Padding(hx)
    hi_hx:=zkproof.Generators(len(hx))
    ch:=zkproof.PedersenVectorComit(hx,hi_hx,H,r_hx)
   


    

    //compute challenge x
    buf:=&bytes.Buffer{}
	gob.NewEncoder(buf).Encode(append(append(append(hi_hx,append(hi,append(c_inputs,G)...)...),ch),H))


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

    var random_tA *big.Int
    var r_w_inner *big.Int
    var r_u_inner *big.Int
    var r_v_inner *big.Int
    var random_tW *big.Int
    var random_tB *big.Int
    var r_hz *big.Int
    var ca_p zkproof.CurvePoint
    var cb_p zkproof.CurvePoint
    var pf_bulletproof_u zkproof.Pf_PdsVec_PubVec
    var pf_bulletproof_v zkproof.Pf_PdsVec_PubVec
    var pf_bulletproof_w zkproof.Pf_PdsVec_PubVec
    var pf_bulletproof_hxzx zkproof.Pf_PdsVec_PubVec
    var pfA zkproof.Pf_PdsComits_PubVec
    var pfB zkproof.Pf_PdsComits_PubVec
    var pfW zkproof.Pf_PdsComits_PubVec
    var c_u_inner zkproof.CurvePoint
    var c_v_inner zkproof.CurvePoint
    var c_w_inner zkproof.CurvePoint
    var c_hz zkproof.CurvePoint


    
     
    //Prover generates the proof
   
    //pfA is to prove ca=g^sum{a_iU[i]}  i range is (len(publicSignals)+1:(len(privateInputs)+2))
   // wg.Add(1)
    //go func(){
	random_tA,_ = rand.Int(rand.Reader,N)
    pfA=zkproof.ZKproofPdsComits_PubVec(hi,U[len(publicSignals)+1:pos_inner],gamma,random_tA,H)

    
    u_inner:=big.NewInt(int64(0))
    for i:=0;i<wires_inner_len;i++{
       temp:=polyf.F.Mul(wires_inner[i],U[pos_inner+i])
       u_inner=polyf.F.Add(temp,u_inner)
    }

       
    
    r_u_inner,_=rand.Int(rand.Reader,N)
    c_u_inner=zkproof.PedersenComit(u_inner,r_u_inner,G,H)


    pf_bulletproof_u=zkproof.ZKproofPdsVec_PubVec(gi,G,H,commit_inner,c_u_inner,wires_inner,zkproof.Padding(U[pos_inner:]),r_inner,r_u_inner,polyf)

    ca_p=zkproof.PedersenComit(A,polyf.F.Add(polyf.F.Neg(random_tA),r_u_inner),G,H)
    //wg.Done()
  //  }()

    //wg.Add(1)
    //go func(){
	random_tB,_ = rand.Int(rand.Reader,N)
    pfB=zkproof.ZKproofPdsComits_PubVec(hi,V[len(publicSignals)+1:pos_inner],gamma,random_tB,H)

    v_inner:=big.NewInt(int64(0))
    for i:=0;i<wires_inner_len;i++{
       temp:=polyf.F.Mul(wires_inner[i],V[pos_inner+i])
       v_inner=polyf.F.Add(temp,v_inner)
    }
    r_v_inner,_=rand.Int(rand.Reader,N)
    c_v_inner=zkproof.PedersenComit(v_inner,r_v_inner,G,H)
    pf_bulletproof_v=zkproof.ZKproofPdsVec_PubVec(gi,G,H,commit_inner,c_v_inner,wires_inner,zkproof.Padding(V[pos_inner:]),r_inner,r_v_inner,polyf)
    cb_p=zkproof.PedersenComit(B,polyf.F.Add(polyf.F.Neg(random_tB),r_v_inner),G,H)
    
    //wg.Done()
    //}()
 
    //wg.Add(1)
   // go func(){
    
	random_tW,_ = rand.Int(rand.Reader,N)
    pfW=zkproof.ZKproofPdsComits_PubVec(hi,W[len(publicSignals)+1:pos_inner],gamma,random_tW,H)
    w_inner:=big.NewInt(int64(0))
    for i:=0;i<wires_inner_len;i++{
       temp:=polyf.F.Mul(wires_inner[i],W[pos_inner+i])
       w_inner=polyf.F.Add(temp,w_inner)
    }
    r_w_inner,_=rand.Int(rand.Reader,N)
    c_w_inner=zkproof.PedersenComit(w_inner,r_w_inner,G,H)
    pf_bulletproof_w=zkproof.ZKproofPdsVec_PubVec(gi,G,H,commit_inner,c_w_inner,wires_inner,zkproof.Padding(W[pos_inner:]),r_inner,r_w_inner,polyf)

    

    fmt.Println("hx:",len(hx))
    
    hz:=big.NewInt(int64(0))
    for i:=0;i<len(hx);i++{
       temp:=polyf.F.Mul(hx[i],X[i])
       hz=polyf.F.Add(temp,hz)
    }
    r_hz,_=rand.Int(rand.Reader,N)
    c_hz=zkproof.PedersenComit(hz,r_hz,G,H)
    pf_bulletproof_hxzx=zkproof.ZKproofPdsVec_PubVec(hi_hx,G,H,ch,c_hz,hx,zkproof.Padding(X),r_hx,r_hz,polyf)
    //wg.Done()
    //}()
    //wg.Wait()
    
    
    //pfA,pfB,pfW,pfH allow verifier compute ca,cb,cw,h(x)*z(x)


  
    rt:=polyf.F.Add(r_w_inner,polyf.F.Add(polyf.F.Neg(random_tW),r_hz))



    right:=zkproof.PedersenComit(polyf.F.Add(HZ,C),rt,G,H)


    pfProduct:=zkproof.ZkproofPdsProduct(ca_p,cb_p,right,G,H,A,B,polyf.F.Add(polyf.F.Neg(random_tA),r_u_inner),polyf.F.Add(polyf.F.Neg(random_tB),r_v_inner),rt,polyf)

    elapsedProof:= time.Since(proveTime)
    fmt.Println("Prove done,used ",elapsedProof)

    
    

    fmt.Println("\ncommitment size:",len(c_inputs)*33,"bytes")
    fmt.Println("proof without counting commitments size:",(33*3+32*2)*3+292+len(pf_bulletproof_u.LR)*33*3+97*4+len(pf_bulletproof_hxzx.LR)*33,"bytes")

    fmt.Println("length of u v w pf LR",len(pf_bulletproof_u.LR))
    fmt.Println("length of hx pf LR",len(pf_bulletproof_hxzx.LR))


    /////////////////verifer check

    start := time.Now()
    
    

    fmt.Println("\nVerification:")

    //Verifier computes ca

    wg.Add(1)
    var ca zkproof.CurvePoint
    go func(){

    	fmt.Println("validation A:",zkproof.ZKverifyPdsComits_PubVec(hi,U[len(publicSignals)+1:pos_inner],pfA,H))
    	fmt.Println("bulletproof check A: ",zkproof.ZKverifyPdsVec_PubVec(gi,G,H,commit_inner,c_u_inner,zkproof.Padding(U[pos_inner:]),polyf,pf_bulletproof_u))
   
    	ca=zkproof.CurvePointVecMult(c_inputs,U[len(publicSignals)+1:pos_inner])
		ca,_=zkproof.CurveSub(ca,pfA.Omega)
		ca,_=zkproof.CurveAdd(ca,c_u_inner)
	
   

	for i:=0;i<len(publicSignals)+1;i++{
		var temp zkproof.CurvePoint
		if(i==0){temp,_=zkproof.CurveScalarMult(G,U[0])
	   }else{ temp,_=zkproof.CurveScalarMult(G,polyf.F.Mul(U[i],publicSignals[i-1]))}
		ca,_=zkproof.CurveAdd(ca,temp)


	}
	wg.Done()
    }()
	

    //Verifier computes cb
    wg.Add(1)
    var cb zkproof.CurvePoint
    go func(){

    fmt.Println("validation B:",zkproof.ZKverifyPdsComits_PubVec(hi,V[len(publicSignals)+1:pos_inner],pfB,H))
    fmt.Println("bulletproof check B: ",zkproof.ZKverifyPdsVec_PubVec(gi,G,H,commit_inner,c_v_inner,zkproof.Padding(V[pos_inner:]),polyf,pf_bulletproof_v))
    cb=zkproof.CurvePointVecMult(c_inputs,V[len(publicSignals)+1:pos_inner])
	cb,_=zkproof.CurveSub(cb,pfB.Omega)
	cb,_=zkproof.CurveAdd(cb,c_v_inner)

	for i:=0;i<len(publicSignals)+1;i++{
		var temp zkproof.CurvePoint
		if(i==0){temp,_=zkproof.CurveScalarMult(G,V[0])
	   }else{ temp,_=zkproof.CurveScalarMult(G,polyf.F.Mul(V[i],publicSignals[i-1]))}
		cb,_=zkproof.CurveAdd(cb,temp)
	}
	wg.Done()
	}()
   


    //Verifier computes cw
    wg.Add(1)
    var cw zkproof.CurvePoint
    go func(){
    fmt.Println("validation C:",zkproof.ZKverifyPdsComits_PubVec(hi,W[len(publicSignals)+1:pos_inner],pfW,H))
	fmt.Println("bulletproof check W: ",zkproof.ZKverifyPdsVec_PubVec(gi,G,H,commit_inner,c_w_inner,zkproof.Padding(W[pos_inner:]),polyf,pf_bulletproof_w))
	
    cw=zkproof.CurvePointVecMult(c_inputs,W[len(publicSignals)+1:pos_inner])
	cw,_=zkproof.CurveSub(cw,pfW.Omega)
	cw,_=zkproof.CurveAdd(cw,c_w_inner)

	for i:=0;i<len(publicSignals)+1;i++{
		var temp zkproof.CurvePoint
		if(i==0){temp,_=zkproof.CurveScalarMult(G,W[0])
	   }else{ temp,_=zkproof.CurveScalarMult(G,polyf.F.Mul(W[i],publicSignals[i-1]))}
		cw,_=zkproof.CurveAdd(cw,temp)
	}
	wg.Done()
    
    }()
    //Verifier computes commitment for hx*zx

    wg.Add(1)

    go func(){
                                                            
    fmt.Println("validation HZ:",zkproof.ZKverifyPdsVec_PubVec(hi_hx,G,H,ch,c_hz,zkproof.Padding(X),polyf,pf_bulletproof_hxzx))
    wg.Done()
	
    }()
    //chcek is chz*cw is the product of ca cb

    wg.Wait()

    c_right,_:=zkproof.CurveAdd(c_hz,cw)
    fmt.Println("product check: com(A*B)==com(C+HZ)",zkproof.ZkverifyPdsProduct(ca,cb,c_right,G,H,pfProduct,polyf))
    elapsed := time.Since(start)
    fmt.Println("Verification done,used ",elapsed)
    fmt.Println("\nnumbers of gates:",len(u))

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("The process is done, press any key and enter to stop: ")
    city, _ := reader.ReadString('\n')
    fmt.Print("You press " + city)

    




   

}
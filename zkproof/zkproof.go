package zkproof

import(
	"fmt"
	"math/big"
	"errors"
    "bytes"
    "github.com/arnaucube/go-snark/fields"
    "github.com/ShuangWu121/secp256k1"
    "github.com/ethereum/go-ethereum/crypto"
    "crypto/rand"
    "github.com/ShuangWu121/PriBankGo/r1csqap"
    "encoding/gob"
    "sync"
    "runtime"
   // "encoding/gob"

)

type CurvePoint struct {
    X *big.Int
    Y *big.Int
}



func CurveScalarMult(G CurvePoint, scalar *big.Int)(CurvePoint,error){
    c:=CurvePoint{}
    r:=secp256k1.SECP256K1()
	if !r.IsOnCurve(G.X,G.Y) {
        fmt.Println("\n Curve Scaler Mult: Not on curve")
		return c, errors.New("\nPedersen Commitment: Not on curve")}
	x,y:=r.ScalarMult(G.X,G.Y,scalar.Bytes())
	c=CurvePoint{x,y}
	return c,nil
}

func CurveScalarDiv(G CurvePoint, scalar *big.Int)(CurvePoint){
    
   minus1, err := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494336", 10)
      if err!=true{
			fmt.Println("Curve scalar div: Field wrong")
		}
   c,_:=CurveScalarMult(G,minus1)
   c,_=CurveScalarMult(c,scalar)
	
   return c
}

func CurveAdd(G,H CurvePoint)(CurvePoint,error){
	c:=CurvePoint{}
    r:=secp256k1.SECP256K1()
	if (!r.IsOnCurve(G.X,G.Y)||(!r.IsOnCurve(H.X,H.Y))) {
        fmt.Println("\nCurve Add: Not on curve")
		return c, errors.New("\nCurve Add: Not on curve")}
	x,y:=r.Add(G.X,G.Y,H.X,H.Y)
	c=CurvePoint{x,y}
	return c,nil
}

func CurveSub(G,H CurvePoint)(CurvePoint,error){
	c:=CurvePoint{}
    r:=secp256k1.SECP256K1()
	if (!r.IsOnCurve(G.X,G.Y)||(!r.IsOnCurve(H.X,H.Y))) {
        fmt.Println("\nCurve Add: Not on curve")
		return c, errors.New("\nCurve Add: Not on curve")}
    minus1, err := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494336", 10)
      if err!=true{
			fmt.Println("Curve scalar div: Field wrong")
		}
    H,_=CurveScalarMult(H,minus1)
	x,y:=r.Add(G.X,G.Y,H.X,H.Y)
	c=CurvePoint{x,y}
	return c,nil
}

func CurvePointVecMult(G []CurvePoint,scaler []*big.Int)(CurvePoint){
    var c CurvePoint
    for i := 0; i < len(G); i++ {
		tempPoint,_:=CurveScalarMult(G[i],scaler[i])
	    if i==0{
	    	c=tempPoint
	    }else {c,_=CurveAdd(c,tempPoint)}

	}
	return c
}


//compute Pedersen Commitment value*G+blind*H
func PedersenComit(value,blind *big.Int,G,H CurvePoint)(CurvePoint){
	c:=CurvePoint{}
	
    c1,_:=CurveScalarMult(G,value)
    c2,_:=CurveScalarMult(H,blind)
    c,_=CurveAdd(c1,c2)

    return c
}


//compute a Perdersen vector commitment c= a1*G1+a2*G2+...+r*H
func PedersenVectorComit(a []*big.Int,G []CurvePoint, H CurvePoint,r *big.Int)(CurvePoint){
    
    c:=CurvePoint{}
    temp:=CurvePoint{}
    var wg sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(1)
	c1,_:=CurveScalarMult(G[0], a[0])
	go func(){
    
	for i := 1; i < len(a)/2; i++ {
		temp,_ := CurveScalarMult(G[i], a[i])
        c1,_=CurveAdd(c1,temp)

	}
	wg.Done()
    }()
    wg.Add(1)
    c2,_:=CurveScalarMult(G[len(a)/2], a[len(a)/2])
	go func(){
	for i := len(a)/2+1; i < len(a); i++ {
		temp,_ := CurveScalarMult(G[i], a[i])
        c2,_=CurveAdd(c2,temp)

	}
	wg.Done()
    }()

    

    wg.Wait()
    c,_=CurveAdd(c2,c1)

	temp,_=CurveScalarMult(H,r)
	c,_=CurveAdd(temp,c)
  
	return c 
}

//create a Pedersen commitment for each element of the vector, using different h over same blnding
//c_i=g^a_ih_i^r
func PedersenComitsForVector(a []*big.Int,hi []CurvePoint,r *big.Int,G CurvePoint)([]CurvePoint,error){
    c:=[]CurvePoint{}
    if len(a)!=len(hi){
    	fmt.Println("\nGenerator hi not match the number of elements in a")
        return c,errors.New("\nGenerator hi not match the number of elements in a")
    }

    temp:=CurvePoint{}
    for i := 0; i < len(a); i++ {
    	temp=PedersenComit(a[i],r,G,hi[i])
		if i==0{
			c=append(c,temp)
		}else{c=append(c,temp)}
	}	
    return c,nil
}

//generate a generators vector as an double array, lenth a
func Generators(len int)([]CurvePoint){
	c:=[]CurvePoint{}
	cv:=secp256k1.SECP256K1()
	G:=CurvePoint{cv.Params().Gx,cv.Params().Gy}
	for i := 0; i < len; i++ {
		temp,_:= CurveScalarMult(G,big.NewInt(int64(i+1)))
		if i==0{	
			c=append(c,temp)
		}else{c=append(c,temp)}
	}
	return c	
}

//zero-knowledge proof for inner product of Pedersen Commitments and a public vector

type Pf_PdsComits_PubVec struct {

	c0 CurvePoint
	Omega CurvePoint
	d1 CurvePoint
	d2 CurvePoint
	theta1 *big.Int
	theta2 *big.Int

}

func ZKproofPdsComits_PubVec(hi []CurvePoint,pubv []*big.Int,gamma,t *big.Int,H CurvePoint)(Pf_PdsComits_PubVec){
    
    //compute c0
    c0,_:=CurveScalarMult(H,gamma)

    max,_:=new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)

    alpha,_:=rand.Int(rand.Reader,max)
    beta,_:=rand.Int(rand.Reader,max)

    

    //compute tau=\prod h_i^pubv_i
    var tau CurvePoint
    for i := 0; i < len(pubv); i++ {
		temp,_:= CurveScalarMult(hi[i],pubv[i])
		if i>0{tau,_=CurveAdd(tau,temp)
	    }else {tau=temp}
	}
  

	//compute OMEGA
	omega:= PedersenComit(gamma,t,tau,H)
		
   //compute d1 d2
   d1,_:=CurveScalarMult(H,alpha)
   d2:= PedersenComit(alpha,beta,tau,H)
		
   //compute challenge x
   x := new(big.Int).SetBytes(crypto.Keccak256([]byte(tau.X.String()+tau.Y.String()+
                                                      omega.X.String()+omega.Y.String()+
                                                      d1.X.String()+d1.Y.String()+
                                                      d2.X.String()+d2.Y.String()+
                                                      H.X.String()+H.Y.String()+
                                                      c0.X.String()+c0.Y.String())))


   //compute theta1 theta2
   
   r, err := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
      if err!=true{
			fmt.Println("ZKproofPdsComits_PubVec Field wrong")
		}
   f := fields.NewFq(r)
   theta1:=f.Sub(alpha,f.Mul(x,gamma))
   theta2:=f.Sub(beta,f.Mul(x,t))
   
   pf:=Pf_PdsComits_PubVec{c0,omega,d1,d2,theta1,theta2}

   return pf
}

func ZKverifyPdsComits_PubVec(hi []CurvePoint,pubv []*big.Int,pf Pf_PdsComits_PubVec,H CurvePoint) (bool){
   
  
   //compute tau=\prod h_i^pubv_i
    var tau CurvePoint
    for i := 0; i < len(pubv); i++ {
		temp,_:= CurveScalarMult(hi[i],pubv[i])
		if i>0{tau,_=CurveAdd(tau,temp)
	    }else {tau=temp}
	}

   x := new(big.Int).SetBytes(crypto.Keccak256([]byte(tau.X.String()+tau.Y.String()+
                                                      pf.Omega.X.String()+pf.Omega.Y.String()+
                                                      pf.d1.X.String()+pf.d1.Y.String()+
                                                      pf.d2.X.String()+pf.d2.Y.String()+
                                                      H.X.String()+H.Y.String()+
                                                      pf.c0.X.String()+pf.c0.Y.String())))
   

   //verify d1==c_0^xh^theta1
   
   temp1,_:=CurveScalarMult(pf.c0,x)
   temp2,_:=CurveScalarMult(H,pf.theta1)
   temp,_:=CurveAdd(temp1,temp2)
   if  (pf.d1.X.Cmp(temp.X)!=0)||(pf.d1.Y.Cmp(temp.Y)!=0){return false}

   //verify d2==tau^theta1 * omega^x * h^theta2
   temp1,_=CurveScalarMult(tau,pf.theta1)
   temp2,_=CurveScalarMult(H,pf.theta2)
   temp,_=CurveAdd(temp1,temp2)
   temp1,_=CurveScalarMult(pf.Omega,x)
   temp,_=CurveAdd(temp,temp1)

   if  (pf.d2.X.Cmp(temp.X)!=0)||(pf.d2.Y.Cmp(temp.Y)!=0){return false}

   
   


  return true
}



//zero-knowledge proof for product of pedersen commitments 
type pf_PdsProduct struct{
	d1 CurvePoint
	d2 CurvePoint
	c0 CurvePoint
	c1 CurvePoint

	theta_a *big.Int
	theta_b *big.Int
	theta1  *big.Int
	theta2 *big.Int
	theta_ab *big.Int

}
func ZkproofPdsProduct(ca,cb,c,G,H CurvePoint,a,b,ra,rb,t *big.Int,polyf r1csqap.PolynomialField)(pf_PdsProduct){
     
     max,_:=new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)

     alpha,_:=rand.Int(rand.Reader,max)
     beta,_:=rand.Int(rand.Reader,max)
     r1,_:=rand.Int(rand.Reader,max)
     r2,_:=rand.Int(rand.Reader,max)
     s0,_:=rand.Int(rand.Reader,max)
     s1,_:=rand.Int(rand.Reader,max)

     d1:=PedersenComit(alpha,r1,G,H)
     d2:=PedersenComit(beta,r2,G,H)
     c0:=PedersenComit(polyf.F.Add(polyf.F.Mul(alpha,b),polyf.F.Mul(beta,a)),s0,G,H)
     c1:=PedersenComit(polyf.F.Mul(alpha,beta),s1,G,H)

    
    
     x := new(big.Int).SetBytes(crypto.Keccak256([]byte(ca.X.String()+ca.Y.String()+
                                                      cb.X.String()+cb.Y.String()+
                                                      c.X.String()+c.Y.String()+
                                                      d1.X.String()+d1.Y.String()+
                                                      d2.X.String()+d2.Y.String()+
                                                      c0.X.String()+c0.Y.String()+
                                                      c1.X.String()+c1.Y.String()+
                                                      H.X.String()+H.Y.String()+
                                                      G.X.String()+G.Y.String())))
     x=polyf.F.Affine(x)

   

     theta_a:=polyf.F.Sub(alpha,polyf.F.Mul(a,x))
     theta_b:=polyf.F.Sub(beta,polyf.F.Mul(b,x))
     theta1:=polyf.F.Sub(r1,polyf.F.Mul(ra,x))
     theta2:=polyf.F.Sub(r2,polyf.F.Mul(rb,x))
     theta_ab:=polyf.F.Add(polyf.F.Sub(polyf.F.Mul(polyf.F.Mul(x,x),t),polyf.F.Mul(x,s0)),s1)

     pf:=pf_PdsProduct{d1,d2,c0,c1,theta_a,theta_b,theta1,theta2,theta_ab}
     return pf


}

func ZkverifyPdsProduct(ca,cb,c,G,H CurvePoint,pf pf_PdsProduct,polyf r1csqap.PolynomialField)(bool){
   
    x := new(big.Int).SetBytes(crypto.Keccak256([]byte(ca.X.String()+ca.Y.String()+
                                                      cb.X.String()+cb.Y.String()+
                                                      c.X.String()+c.Y.String()+
                                                      pf.d1.X.String()+pf.d1.Y.String()+
                                                      pf.d2.X.String()+pf.d2.Y.String()+
                                                      pf.c0.X.String()+pf.c0.Y.String()+
                                                      pf.c1.X.String()+pf.c1.Y.String()+
                                                      H.X.String()+H.Y.String()+
                                                      G.X.String()+G.Y.String())))
    x=polyf.F.Affine(x)

    var wg sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag:=1

	wg.Add(1)
    go func(){
    temp1,_:=CurveScalarMult(G,pf.theta_a)
    temp2,_:=CurveScalarMult(H,pf.theta1)
	temp1,_=CurveAdd(temp1,temp2)
	temp,_:=CurveScalarMult(ca,x)
	temp1,_=CurveAdd(temp1,temp)
    
    if  (pf.d1.X.Cmp(temp1.X)!=0)||(pf.d1.Y.Cmp(temp1.Y)!=0){flag=0}
    wg.Done()
    }()

    wg.Add(1)
    go func(){
    ttemp1,_:=CurveScalarMult(G,pf.theta_b)
    ttemp2,_:=CurveScalarMult(H,pf.theta2)
	ttemp1,_=CurveAdd(ttemp1,ttemp2)
	ttemp,_:=CurveScalarMult(cb,x)
	ttemp1,_=CurveAdd(ttemp1,ttemp)

	if  (pf.d2.X.Cmp(ttemp1.X)!=0)||(pf.d2.Y.Cmp(ttemp1.Y)!=0){flag=0}
	wg.Done()
    }()


    wg.Add(1)
    go func(){
	tempa1,_:=CurveScalarMult(G,polyf.F.Mul(pf.theta_b,pf.theta_a))
    tempa2,_:=CurveScalarMult(H,pf.theta_ab)
	tempa1,_=CurveAdd(tempa1,tempa2)
	tempa2,_=CurveScalarMult(pf.c0,x)
	tempa1,_=CurveAdd(tempa1,tempa2)

	tempa2,_=CurveScalarMult(c,polyf.F.Mul(x,x))
	tempa,_:=CurveAdd(pf.c1,tempa2)


    
    if  (tempa.X.Cmp(tempa1.X)!=0)||(tempa.Y.Cmp(tempa1.Y)!=0){flag=0}
    wg.Done()
    }()
    wg.Wait()

    
	if flag==1{return true}else{return false}
}


//zero-knowledge proof for addition of pedersen commitments



func ComputeL(gi []CurvePoint,a,b []*big.Int,G,H CurvePoint,r *big.Int,polyf r1csqap.PolynomialField)(CurvePoint){

    n:=len(gi)
    half:=n/2
    
    var l1 CurvePoint
    for i := 0; i < half; i++ {
		temp,_:= CurveScalarMult(gi[i+half],a[i])
		if i>0{l1,_=CurveAdd(l1,temp)
	    }else {l1=temp}
	}
    
    var ab *big.Int
    for i := 0; i < half; i++ {
    	tempab:=polyf.F.Mul(a[i],b[i+half])
		if i>0{ab=polyf.F.Add(ab,tempab)
	    }else {ab=tempab}
	}

	l2,_:= CurveScalarMult(G,ab)
	l3,_:= CurveScalarMult(H,r)

    L,_:=CurveAdd(l1,l2)
    L,_=CurveAdd(L,l3)
    return L

}

func ComputeR(gi []CurvePoint,a,b []*big.Int,G,H CurvePoint,r *big.Int,polyf r1csqap.PolynomialField)(CurvePoint){

	n:=len(gi)
    half:=n/2
    
    var R1 CurvePoint
    for i := 0; i < half; i++ {
		temp,_:= CurveScalarMult(gi[i],a[i+half])
		if i>0{R1,_=CurveAdd(R1,temp)
	    }else {R1=temp}
	}
    
    var ab *big.Int
    for i := 0; i < half; i++ {
    	tempab:=polyf.F.Mul(a[i+half],b[i])
		if i>0{ab=polyf.F.Add(ab,tempab)
	    }else {ab=tempab}
	}

	R2,_:= CurveScalarMult(G,ab)
	R3,_:= CurveScalarMult(H,r)

    R,_:=CurveAdd(R1,R2)
    R,_=CurveAdd(R,R3)
    return R

}

func ComputeG_prime(x *big.Int,gi []CurvePoint,polyf r1csqap.PolynomialField)([]CurvePoint){
    n:=len(gi)
    half:=n/2
    
    var g_prime []CurvePoint
	for i:=0;i<half;i++{
       temp1,_:=CurveScalarMult(gi[i],polyf.F.Inverse(x))
       temp2,_:=CurveScalarMult(gi[i+half],x)
       temp,_:=CurveAdd(temp1,temp2)
       g_prime=append(g_prime,temp)
	}

	return g_prime

}

func ComputeC_prime(x *big.Int,L,R,c CurvePoint,polyf r1csqap.PolynomialField)(CurvePoint){
	temp1,_:=CurveScalarMult(L,polyf.F.Mul(x,x))
	temp2,_:=CurveScalarMult(R,polyf.F.Mul(polyf.F.Inverse(x),polyf.F.Inverse(x)))
	c,_=CurveAdd(c,temp1)
	c,_=CurveAdd(c,temp2)
	return c
}

func ComputeFold(x *big.Int,a []*big.Int,polyf r1csqap.PolynomialField)([]*big.Int){
	n:=len(a)
    half:=n/2
    
    var a_prime []*big.Int
	for i:=0;i<half;i++{
       temp1:=polyf.F.Mul(a[i],polyf.F.Inverse(x))
       temp2:=polyf.F.Mul(a[i+half],x)
       temp:=polyf.F.Add(temp1,temp2)
       a_prime=append(a_prime,temp)
	}

	return a_prime

}
func Padding(input []*big.Int)([]*big.Int){
	length:=len(input)
	//fmt.Println("length",length)
	power:=0
	paddingto:=0
	for i := 1; i < length; i=i*2 {
    	power=power+1
    	paddingto=i*2
	}
	//fmt.Println("paddingto",paddingto)
	//fmt.Println("power",power)

	for i := 0; i < (paddingto-length); i++{
    	input=append(input,big.NewInt(int64(0)))
	}
	return input
}

type Pf_PdsVec_PubVec struct{
	LR []CurvePoint
	d  CurvePoint
	theta1 *big.Int
	theta2 *big.Int
}
func ZKproofPdsVec_PubVec(gi []CurvePoint,g,h,ca,cab CurvePoint,a,b []*big.Int,ra,rab *big.Int, polyf r1csqap.PolynomialField)(Pf_PdsVec_PubVec){//,c,h CurvePoint,polyf r1csqap.PolynomialField
   
   if(len(gi)!=len(a) && len(a)!=len(b)){
   	fmt.Println("zkproof for Pedersen vector commitments: vector length wrong")
   }
   r:=polyf.F.Add(ra,rab)
   c,_:=CurveAdd(ca,cab)
   max,_:=new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
   var LR [] CurvePoint

   for i:=0;len(gi)>1;i++{

   		

   		r1,_:=rand.Int(rand.Reader,max)
   		L:=ComputeL(gi,a,b,g,h,r1,polyf)

   		r2,_:=rand.Int(rand.Reader,max)
  		R:=ComputeR(gi,a,b,g,h,r2,polyf)

  		LR=append(append(LR,L),R)

 	    //compute challenge x
   	    buf:=&bytes.Buffer{}
		gob.NewEncoder(buf).Encode(append(append(append(gi,c),g),h))

		


    	x := new(big.Int).SetBytes(crypto.Keccak256(buf.Bytes()))
    	x=polyf.F.Affine(x)  

    	gi=ComputeG_prime(x,gi,polyf)
    	c=ComputeC_prime(x,L,R,c,polyf)



    	a=ComputeFold(polyf.F.Inverse(x),a,polyf)
    	b=ComputeFold(x,b,polyf)
    	r=polyf.F.Add(r,polyf.F.Add(polyf.F.Mul(r1,polyf.F.Mul(x,x)),polyf.F.Mul(r2,polyf.F.Mul(polyf.F.Inverse(x),polyf.F.Inverse(x)))))

        /*if i==0 {
    		fmt.Println("c prime is :",c)
    		fmt.Println("gi lenth:",len(gi))
    		fmt.Println("a lenth:",len(a))
    		fmt.Println("b lenth:",len(b))

            ab:=big.NewInt(int64(0))
            for i:=0;i<len(a);i++{
       			temp:=polyf.F.Mul(a[i],b[i])
       			ab=polyf.F.Add(temp,ab)
            }

            c_a_ab1,_:=CurveScalarMult(gi[0],a[0])
            c_a_ab2,_:=CurveScalarMult(gi[1],a[1])
            c_a_ab,_:=CurveAdd(c_a_ab1,c_a_ab2)

            c_ab,_:=CurveScalarMult(g,polyf.F.Add(polyf.F.Mul(a[0],b[0]),polyf.F.Mul(a[1],b[1])))

            c_a_ab,_=CurveAdd(c_a_ab,c_ab)

            h_r,_:=CurveScalarMult(h,r)
            c_a_ab,_=CurveAdd(h_r,c_a_ab)

            fmt.Println("c prime from raw compute:",c_a_ab)
            fmt.Println("c prime from verifier:",c)
    	}

    	if i==1 {
    		fmt.Println("c prime is :",c)
    		fmt.Println("gi lenth:",len(gi))
    		fmt.Println("a lenth:",len(a))
    		fmt.Println("b lenth:",len(b))

            ab:=big.NewInt(int64(0))
            for i:=0;i<len(a);i++{
       			temp:=polyf.F.Mul(a[i],b[i])
       			ab=polyf.F.Add(temp,ab)
            }

            c_a_ab,_:=CurveScalarMult(gi[0],a[0])

            c_ab,_:=CurveScalarMult(g,polyf.F.Mul(a[0],b[0]))

            c_a_ab,_=CurveAdd(c_a_ab,c_ab)

            h_r,_:=CurveScalarMult(h,r)
            c_a_ab,_=CurveAdd(h_r,c_a_ab)

            fmt.Println("c prime from raw compute:",c_a_ab)
            fmt.Println("c prime from verifier:",c)
    	}
        */



   }

   	
   	alpha1,_:=rand.Int(rand.Reader,max)
   	alpha2,_:=rand.Int(rand.Reader,max)
 
   	
    temp1,_:=CurveScalarMult(gi[0],alpha1)
    temp2,_:=CurveScalarMult(g,polyf.F.Mul(alpha1,b[0]))
    temp1,_=CurveAdd(temp1,temp2)
    temp2,_=CurveScalarMult(h,alpha2)
    d,_:=CurveAdd(temp1,temp2)
    
    //compute challenge x
   	buf:=&bytes.Buffer{}
	gob.NewEncoder(buf).Encode(append(append(append(gi,c),g),h))


    x := new(big.Int).SetBytes(crypto.Keccak256(buf.Bytes()))
    x=polyf.F.Affine(x)  

   	
   	theta1:=polyf.F.Sub(alpha1,polyf.F.Mul(x,a[0]))
   	theta2:=polyf.F.Sub(alpha2,polyf.F.Mul(x,r))
    pf:=Pf_PdsVec_PubVec{LR,d,theta1,theta2}
    return pf

    

}

func ZKverifyPdsVec_PubVec(gi []CurvePoint,g,h,ca,cab CurvePoint,b []*big.Int, polyf r1csqap.PolynomialField,pf Pf_PdsVec_PubVec)(bool){

	c,_:=CurveAdd(ca,cab)
	for i:=0;len(gi)>1;i++{

 	    //compute challenge x
   	    buf:=&bytes.Buffer{}
		gob.NewEncoder(buf).Encode(append(append(append(gi,c),g),h))
    	x := new(big.Int).SetBytes(crypto.Keccak256(buf.Bytes()))
    	x=polyf.F.Affine(x)  

    	gi=ComputeG_prime(x,gi,polyf)
    	c=ComputeC_prime(x,pf.LR[2*i],pf.LR[2*i+1],c,polyf)

    	//if i==0{fmt.Println("c prime in verification:",c)}
    	//if i==1{fmt.Println("c prime in verification,i==1:",c)}

    	b=ComputeFold(x,b,polyf)

   }

    //compute challenge x
   	buf:=&bytes.Buffer{}
	gob.NewEncoder(buf).Encode(append(append(append(gi,c),g),h))


    x := new(big.Int).SetBytes(crypto.Keccak256(buf.Bytes()))
    x=polyf.F.Affine(x)  

    temp1,_:=CurveScalarMult(c,x)
    temp2,_:=CurveScalarMult(gi[0],pf.theta1)
    temp1,_=CurveAdd(temp1,temp2)

    temp2,_=CurveScalarMult(g,polyf.F.Mul(b[0],pf.theta1))
    temp1,_=CurveAdd(temp1,temp2)
    temp2,_=CurveScalarMult(h,pf.theta2)
    temp,_:=CurveAdd(temp1,temp2)



   if  (temp.X.Cmp(pf.d.X)!=0)||(temp.Y.Cmp(pf.d.Y)!=0){return false}

   return true
}
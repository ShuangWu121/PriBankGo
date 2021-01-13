package zkproof

import(
	"fmt"
	"math/big"
	"errors"
   // "bytes"
    "github.com/arnaucube/go-snark/fields"
    "github.com/ShuangWu121/secp256k1"
    "github.com/ethereum/go-ethereum/crypto"
    "crypto/rand"
    "github.com/ShuangWu121/PriBankGo/r1csqap"
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
	for i := 0; i < len(a); i++ {
		temp,_ := CurveScalarMult(G[i], a[i])
        if i==0{
        	c=temp
        }else{c,_=CurveAdd(c,temp)}

	}
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

type pf_PdsComits_PubVec struct {

	c0 CurvePoint
	Omega CurvePoint
	d1 CurvePoint
	d2 CurvePoint
	theta1 *big.Int
	theta2 *big.Int

}

func ZKproofPdsComits_PubVec(hi []CurvePoint,pubv []*big.Int,gamma,t *big.Int,H CurvePoint)(pf_PdsComits_PubVec){
    
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
   
   pf:=pf_PdsComits_PubVec{c0,omega,d1,d2,theta1,theta2}

   return pf
}

func ZKverifyPdsComits_PubVec(hi []CurvePoint,pubv []*big.Int,pf pf_PdsComits_PubVec,H CurvePoint) (bool){
   
  
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

    temp1,_:=CurveScalarMult(G,pf.theta_a)
    temp2,_:=CurveScalarMult(H,pf.theta1)
	temp1,_=CurveAdd(temp1,temp2)
	temp,_:=CurveScalarMult(ca,x)
	temp1,_=CurveAdd(temp1,temp)
    
    if  (pf.d1.X.Cmp(temp1.X)!=0)||(pf.d1.Y.Cmp(temp1.Y)!=0){return false}

  
    temp1,_=CurveScalarMult(G,pf.theta_b)
    temp2,_=CurveScalarMult(H,pf.theta2)
	temp1,_=CurveAdd(temp1,temp2)
	temp,_=CurveScalarMult(cb,x)
	temp1,_=CurveAdd(temp1,temp)

	if  (pf.d2.X.Cmp(temp1.X)!=0)||(pf.d2.Y.Cmp(temp1.Y)!=0){return false}

	temp1,_=CurveScalarMult(G,polyf.F.Mul(pf.theta_b,pf.theta_a))
    temp2,_=CurveScalarMult(H,pf.theta_ab)
	temp1,_=CurveAdd(temp1,temp2)
	temp2,_=CurveScalarMult(pf.c0,x)
	temp1,_=CurveAdd(temp1,temp2)

	temp2,_=CurveScalarMult(c,polyf.F.Mul(x,x))
	temp,_=CurveAdd(pf.c1,temp2)


    
    if  (temp.X.Cmp(temp1.X)!=0)||(temp.Y.Cmp(temp1.Y)!=0){return false}


	return true
}


//zero-knowledge proof for addition of pedersen commitments

type pf_PdsAdd struct{

}

func Zk_proofPdsAdd(ca,cb,c,G,H CurvePoint,a,b,ra,rb,t *big.Int,polyf r1csqap.PolynomialField)(){


}


func ComputeL(gi []CurvePoint,a,b []*big.Int,H,G,c CurvePoint,r *big.Int)(CurvePoint){

    n:=len(gi)
    half:=n/2
    
    var l1 CurvePoint
    for i := 0; i < half; i++ {
		temp,_:= CurveScalarMult(gi[i+half],a[i])
		if i>0{l1,_=CurveAdd(l1,temp)
	    }else {l1=temp}
	}

	N, _ := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
    f := fields.NewFq(N)
    
    var ab *big.Int
    for i := 0; i < half; i++ {
    	tempab:=f.Mul(a[i],b[i+half])
		if i>0{ab=f.Add(ab,tempab)
	    }else {ab=tempab}
	}

	l2,_:= CurveScalarMult(G,ab)
	l3,_:= CurveScalarMult(H,r)

    L,_:=CurveAdd(l1,l2)
    L,_=CurveAdd(L,l3)
    return L

}


func ZKproofPdsVec_PubVec(gi []CurvePoint,a,b []*big.Int,c,h CurvePoint)(){

}

func ZKverifyPdsVec_PubVec()(){

}
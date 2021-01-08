package zkproof

import(
	"fmt"
	"math/big"
	"errors"
   // "bytes"
    "github.com/arnaucube/go-snark/fields"
    "github.com/ShuangWu121/secp256k1"
    "github.com/ethereum/go-ethereum/crypto"
)

type CurvePoint struct {
    x *big.Int
    y *big.Int
}



func CurveScalarMult(G CurvePoint, scalar *big.Int)(CurvePoint,error){
    c:=CurvePoint{}
    r:=secp256k1.SECP256K1()
	if !r.IsOnCurve(G.x,G.y) {
        fmt.Println("\n Curve Scaler Mult: Not on curve")
		return c, errors.New("\nPedersen Commitment: Not on curve")}
	x,y:=r.ScalarMult(G.x,G.y,scalar.Bytes())
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
	if !r.IsOnCurve(G.x,G.y)||(!r.IsOnCurve(H.x,H.y)) {
        fmt.Println("\nCurve Add: Not on curve")
		return c, errors.New("\nPedersen Commitment: Not on curve")}
	x,y:=r.Add(G.x,G.y,H.x,H.y)
	c=CurvePoint{x,y}
	return c,nil
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
	omega CurvePoint
	d1 CurvePoint
	d2 CurvePoint
	theta1 *big.Int
	theta2 *big.Int

}

func ZKproofPdsComits_PubVec(hi []CurvePoint,pubv []*big.Int,gamma,alpha,beta,t *big.Int,H CurvePoint)(pf_PdsComits_PubVec){
    
    //compute c0
    c0,_:=CurveScalarMult(H,gamma)

    

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
   x := new(big.Int).SetBytes(crypto.Keccak256([]byte(tau.x.String()+tau.y.String()+
                                                      omega.x.String()+omega.y.String()+
                                                      d1.x.String()+d1.y.String()+
                                                      d2.x.String()+d2.y.String()+
                                                      H.x.String()+H.y.String()+
                                                      c0.x.String()+c0.y.String())))


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

   x := new(big.Int).SetBytes(crypto.Keccak256([]byte(tau.x.String()+tau.y.String()+
                                                      pf.omega.x.String()+pf.omega.y.String()+
                                                      pf.d1.x.String()+pf.d1.y.String()+
                                                      pf.d2.x.String()+pf.d2.y.String()+
                                                      H.x.String()+H.y.String()+
                                                      pf.c0.x.String()+pf.c0.y.String())))

   //verify d1==c_0^xh^theta1
   
   temp1,_:=CurveScalarMult(pf.c0,x)
   temp2,_:=CurveScalarMult(H,pf.theta1)
   temp,_:=CurveAdd(temp1,temp2)
   if  (pf.d1.x.Cmp(temp.x)!=0)||(pf.d1.y.Cmp(temp.y)!=0){return false}

   //verify d2==tau^theta1 * omega^x * h^theta2
   temp1,_=CurveScalarMult(tau,pf.theta1)
   temp2,_=CurveScalarMult(H,pf.theta2)
   temp,_=CurveAdd(temp1,temp2)
   temp1,_=CurveScalarMult(pf.omega,x)
   temp,_=CurveAdd(temp,temp1)

   if  (pf.d2.x.Cmp(temp.x)!=0)||(pf.d2.y.Cmp(temp.y)!=0){return false}

   
   


  return true
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
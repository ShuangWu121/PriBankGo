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

func PedersenVectorComit(a []*big.Int)(CurvePoint){
	r:=secp256k1.SECP256K1()
	x:=r.Params().Gx
	y:=r.Params().Gy
	for i := 0; i < len(a); i++ {
		xtemp,ytemp := r.ScalarMult(r.Params().Gx,r.Params().Gy, a[i].Bytes())
        x,y=r.Add(x,y,xtemp,ytemp)

	}
    c:=CurvePoint{x,y}
	return c 
}

//compute Pedersen Commitment value*G+blind*H
func PedersenComit(value,blind,Gx,Gy,Hx,Hy *big.Int)(*big.Int,*big.Int,error){
	r:=secp256k1.SECP256K1()
	if !r.IsOnCurve(Gx,Gy)||(!r.IsOnCurve(Hx,Hy)) {
        fmt.Println("\nPedersen Commitment: Not on curve")
		return big.NewInt(int64(0)),big.NewInt(int64(0)), errors.New("\nPedersen Commitment: Not on curve")}
	
    x1,y1:=r.ScalarMult(Gx,Gy,value.Bytes())
    x2,y2:=r.ScalarMult(Hx,Hy,blind.Bytes())
    x,y:=r.Add(x1,y1,x2,y2)
    return x,y,nil 
}

//create a Pedersen commitment for each element of the vector, using different h over same blnding
//c_i=g^a_ih_i^r
func PedersenComitsForVector(a []*big.Int,hi [][]*big.Int,blnding,Gx,Gy *big.Int)([][]*big.Int,error){
    var c [][]*big.Int
    if len(a)!=len(hi){
    	fmt.Println("\nGenerator hi not match the number of elements in a")
        return c,errors.New("\nGenerator hi not match the number of elements in a")
    }
    for i := 0; i < len(a); i++ {
		tempx,tempy,err := PedersenComit(a[i],blnding,Gx,Gy,hi[i][0],hi[i][1])
		if err!=nil{
			fmt.Println("PedersenComitForVector wrong")
			return c,errors.New("PedersenComitForVector wrong")
		}
		c=append(c,[]*big.Int{tempx,tempy})
	}	
    return c,nil
}

//generate a generators as an double array
func Generators(a int)([][]*big.Int){
	var c [][]*big.Int
	cv:=secp256k1.SECP256K1()
	for i := 0; i < a; i++ {
		tempx,tempy:= cv.ScalarMult(cv.Params().Gx,cv.Params().Gy,big.NewInt(int64(i+1)).Bytes())
		c=append(c,[]*big.Int{tempx,tempy})
	}
	return c	
}

//zero-knowledge proof for inner product of Pedersen Commitments and a public vector

func ZKproofPdsComits_PubVec(ci,hi [][]*big.Int,pubv []*big.Int,gamma,alpha,beta,t *big.Int,H CurvePoint)(CurvePoint,CurvePoint,CurvePoint,CurvePoint,*big.Int,*big.Int){
    //compute c0
    cv:=secp256k1.SECP256K1()
    c0x,c0y:=cv.ScalarMult(H.x,H.y,gamma.Bytes())

    var taux *big.Int
    var tauy *big.Int

    //compute tau=\prod h_i^pubv_i
    for i := 0; i < len(pubv); i++ {
		tempx,tempy:= cv.ScalarMult(hi[i][0],hi[i][1],pubv[i].Bytes())
		if i>0{taux,tauy=cv.Add(taux,tauy,tempx,tempy)
	    }else {taux=tempx;tauy=tempy}
	}

	//compute OMEGA
	omegax,omegay,err1:= PedersenComit(gamma,t,taux,tauy,H.x,H.y)
		if err1!=nil{
			fmt.Println("ZKproofPdsComits_PubVec wrong")
		}
   //compute d1 d2
   d1x,d1y:=cv.ScalarMult(H.x,H.y,alpha.Bytes())
   d2x,d2y,err2:= PedersenComit(alpha,beta,taux,tauy,H.x,H.y)
		if err2!=nil{
			fmt.Println("ZKproofPdsComits_PubVec wrong")
		}
   //compute challenge x
   x := new(big.Int).SetBytes(crypto.Keccak256([]byte(taux.String()+tauy.String()+
                                                      omegax.String()+omegay.String()+
                                                      d1x.String()+d1y.String()+
                                                      d2x.String()+d2y.String()+
                                                      H.x.String()+H.y.String()+
                                                      c0x.String()+c0y.String())))


   //compute theta1 theta2
   
   r, err3 := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
      if err3!=true{
			fmt.Println("ZKproofPdsComits_PubVec Field wrong")
		}
   f := fields.NewFq(r)
   theta1:=f.Sub(alpha,f.Mul(x,gamma))
   theta2:=f.Sub(beta,f.Mul(x,t))
   

   c0:=CurvePoint{c0x,c0y}
   omega:=CurvePoint{omegax,omegay}
   d1:=CurvePoint{d1x,d1y}
   d2:=CurvePoint{d2x,d2y}

   return c0,omega,d1,d2,theta1,theta2
}

func ZKverifyPdsComits_PubVec(ci,hi [][]*big.Int,pubv []*big.Int,H,c0,omega,d1,d2 CurvePoint,theta1,theta2 *big.Int) (CurvePoint,bool){
   
   Cab:=CurvePoint{}
  
   
   
   cv:=secp256k1.SECP256K1()

   var taux *big.Int
   var tauy *big.Int
  
   //compute tau=\prod h_i^pubv_i
    for i := 0; i < len(pubv); i++ {
		tempx,tempy:= cv.ScalarMult(hi[i][0],hi[i][1],pubv[i].Bytes())
		if i>0{taux,tauy=cv.Add(taux,tauy,tempx,tempy)
	    }else {taux=tempx;tauy=tempy}
	}

   x := new(big.Int).SetBytes(crypto.Keccak256([]byte(taux.String()+tauy.String()+
                                                      omega.x.String()+omega.y.String()+
                                                      d1.x.String()+d1.y.String()+
                                                      d2.x.String()+d2.y.String()+
                                                      H.x.String()+H.y.String()+
                                                      c0.x.String()+c0.y.String())))

   //verify d1==c_0^xh^theta1
   
   tempx1,tempy1:=cv.ScalarMult(c0.x,c0.y,x.Bytes())
   tempx2,tempy2:=cv.ScalarMult(H.x,H.y,theta1.Bytes())
   tempx,tempy:=cv.Add(tempx1,tempy1,tempx2,tempy2)
   if  (d1.x.Cmp(tempx)!=0)||(d1.y.Cmp(tempy)!=0){return Cab,false}

   //verify d2==c_0^xh^theta1 
   tempx1,tempy1=cv.ScalarMult(taux,tauy,theta1.Bytes())
   tempx2,tempy2=cv.ScalarMult(H.x,H.y,theta2.Bytes())
   tempx1,tempy1=cv.Add(tempx1,tempy1,tempx2,tempy2)
   tempx2,tempy2=cv.ScalarMult(omega.x,omega.y,x.Bytes())
   tempx,tempy=cv.Add(tempx1,tempy1,tempx2,tempy2)

   if  (d2.x.Cmp(tempx)!=0)||(d2.y.Cmp(tempy)!=0){return Cab,false}

   
   //compute Cab=\prod c_i^{b_i}/omega
   var Cabx *big.Int
   var Caby *big.Int
   for i := 0; i < len(pubv); i++ {
		tempx,tempy:= cv.ScalarMult(ci[i][0],ci[i][1],pubv[i].Bytes())
		if i>0{Cabx,Caby=cv.Add(Cabx,Caby,tempx,tempy)
	    }else {Cabx=tempx;Caby=tempy}
	}
   
   minus1, err := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494336", 10)
      if err!=true{
			fmt.Println("ZKverifyPdsComits_PubVec Field wrong")
		}
   omega.x,omega.y=cv.ScalarMult(omega.x,omega.y,minus1.Bytes())
   Cabx,Caby=cv.Add(Cabx,Caby,omega.x,omega.y)
   Cab=CurvePoint{Cabx,Caby}



  return Cab,true
}
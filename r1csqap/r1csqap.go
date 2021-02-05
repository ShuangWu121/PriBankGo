package r1csqap

import (
	"bytes"
	"math/big"
	"fmt"
    "time"
	"github.com/arnaucube/go-snark/fields"
	"runtime"
	"sync"
	"log"
    "os"
	
)

// Transpose transposes the *big.Int matrix
func Transpose(matrix [][]*big.Int) [][]*big.Int {
	var r [][]*big.Int
	for i := 0; i < len(matrix[0]); i++ {
		var row []*big.Int
		for j := 0; j < len(matrix); j++ {
			row = append(row, matrix[j][i])
		}
		r = append(r, row)
	}
	return r
}

// ArrayOfBigZeros creates a *big.Int array with n elements to zero
func ArrayOfBigZeros(num int) []*big.Int {
	bigZero := big.NewInt(int64(0))
	var r []*big.Int
	for i := 0; i < num; i++ {
		r = append(r, bigZero)
	}
	return r
}
func BigArraysEqual(a, b []*big.Int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !bytes.Equal(a[i].Bytes(), b[i].Bytes()) {
			return false
		}
	}
	return true
}

// PolynomialField is the Polynomial over a Finite Field where the polynomial operations are performed
type PolynomialField struct {
	F fields.Fq
}

// NewPolynomialField creates a new PolynomialField with the given FiniteField
func NewPolynomialField(f fields.Fq) PolynomialField {
	return PolynomialField{
		f,
	}
}

// Mul multiplies two polinomials over the Finite Field
func (pf PolynomialField) Mul(a, b []*big.Int) []*big.Int {
	r := ArrayOfBigZeros(len(a) + len(b) - 1)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			r[i+j] = pf.F.Add(
				r[i+j],
				pf.F.Mul(a[i], b[j]))
		}
	}
	return r
}

// Div divides two polinomials over the Finite Field, returning the result and the remainder
func (pf PolynomialField) Div(a, b []*big.Int) ([]*big.Int, []*big.Int) {
	// https://en.wikipedia.org/wiki/Division_algorithm
	r := ArrayOfBigZeros(len(a) - len(b) + 1)
	rem := a
	for len(rem) >= len(b) {
		l := pf.F.Div(rem[len(rem)-1], b[len(b)-1])
		pos := len(rem) - len(b)
		r[pos] = l
		aux := ArrayOfBigZeros(pos)
		aux1 := append(aux, l)
		aux2 := pf.Sub(rem, pf.Mul(b, aux1))
		rem = aux2[:len(aux2)-1]
	}
	return r, rem
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Add adds two polinomials over the Finite Field
func (pf PolynomialField) Add(a, b []*big.Int) []*big.Int {
	r := ArrayOfBigZeros(max(len(a), len(b)))
	for i := 0; i < len(a); i++ {
		r[i] = pf.F.Add(r[i], a[i])
	}
	for i := 0; i < len(b); i++ {
		r[i] = pf.F.Add(r[i], b[i])
	}
	return r
}

// Sub subtracts two polinomials over the Finite Field
func (pf PolynomialField) Sub(a, b []*big.Int) []*big.Int {
	r := ArrayOfBigZeros(max(len(a), len(b)))
	for i := 0; i < len(a); i++ {
		r[i] = pf.F.Add(r[i], a[i])
	}
	for i := 0; i < len(b); i++ {
		r[i] = pf.F.Sub(r[i], b[i])
	}
	return r
}

// Eval evaluates the polinomial over the Finite Field at the given value x
func (pf PolynomialField) Eval(v []*big.Int, x *big.Int) *big.Int {
	r := big.NewInt(int64(0))
	for i := 0; i < len(v); i++ {
		xi := pf.F.Exp(x, big.NewInt(int64(i)))
		elem := pf.F.Mul(v[i], xi)
		r = pf.F.Add(r, elem)
	}
	return r
}

// NewPolZeroAt generates a new polynomial that has value zero at the given value
func (pf PolynomialField) NewPolZeroAt(pointPos, totalPoints int, height *big.Int) []*big.Int {
	fac := big.NewInt(int64(1))
	for i := 1; i < totalPoints+1; i++ {
		if i != pointPos {
			fac = pf.F.Mul(fac, big.NewInt(int64(pointPos - i)))
		}
	}
	facBig :=fac// big.NewInt(int64(fac))
	hf := pf.F.Div(height, facBig)
	r := []*big.Int{hf}
	for i := 1; i < totalPoints+1; i++ {
		if i != pointPos {
			ineg := big.NewInt(int64(-i))
			b1 := big.NewInt(int64(1))
			r = pf.Mul(r, []*big.Int{ineg, b1})
		}
	}
	return r
}

// LagrangeInterpolation performs the Lagrange Interpolation / Lagrange Polynomials operation
func (pf PolynomialField) LagrangeInterpolation(v []*big.Int) []*big.Int {
	// https://en.wikipedia.org/wiki/Lagrange_polynomial
	var r []*big.Int
	for i := 0; i < len(v); i++ {
		r = pf.Add(r, pf.NewPolZeroAt(i+1, len(v), v[i]))
	}
	//
	return r
}

func ParalleUVW(a [][]*big.Int,X uint8,pf PolynomialField)([][]*big.Int){
    var alphas [][]*big.Int
    for i := 0; i < len(a); i++ {
		alphas = append(alphas, pf.LagrangeInterpolation(a[i]))
	}
	fmt.Println("Compute ",string(X)," done")
	return alphas
}



// R1CSToQAP converts the R1CS values to the QAP values
func (pf PolynomialField) R1CSToQAP(a, b, c [][]*big.Int) ([][]*big.Int, [][]*big.Int, [][]*big.Int, []*big.Int) {
	start := time.Now()
	aT := Transpose(a)
	bT := Transpose(b)
	cT := Transpose(c)

	var wg sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("cpu numbers",runtime.NumCPU())

	var alphas [][]*big.Int

	var alphasTemp [][]*big.Int
   
 
    for i := 0; i < len(aT); i++ {
			alphasTemp=append(alphasTemp,[]*big.Int{big.NewInt(int64(0))})
	}
	
    for i := 0; i < len(aT); i++ {
    	wg.Add(1)
    	a:=i
		go func(){
			
			//fmt.Println("index",a,"start")
			alphasTemp[a]=pf.LagrangeInterpolation(aT[a])
			defer wg.Done()
		}()
	}
	
    

    /*
    go func(){
    	
    	alphas=ParalleUVW(aT,'u',pf)
        wg.Done()  
    }()*/

	var betas [][]*big.Int

	var betasTemp [][]*big.Int
   
 
    for i := 0; i < len(bT); i++ {
			betasTemp=append(betasTemp,[]*big.Int{big.NewInt(int64(0))})
	}
	
    for i := 0; i < len(bT); i++ {
    	wg.Add(1)
    	a:=i
		go func(){
			
			//fmt.Println("index",a,"start")
			betasTemp[a]=pf.LagrangeInterpolation(bT[a])
			defer wg.Done()
		}()
	}
    /*wg.Add(1)
	 go func(){
	 	
  	  	betas=ParalleUVW(bT,'v',pf)
    	wg.Done()   
    }()

    */

    
	var gammas [][]*big.Int

	var gammasTemp [][]*big.Int
   
 
    for i := 0; i < len(cT); i++ {
			gammasTemp=append(gammasTemp,[]*big.Int{big.NewInt(int64(0))})
	}
	
    for i := 0; i < len(cT); i++ {
    	wg.Add(1)
    	a:=i
		go func(){
			
			//fmt.Println("index",a,"start")
			gammasTemp[a]=pf.LagrangeInterpolation(cT[a])
			defer wg.Done()
		}()
	}
	
   /* wg.Add(1)
	go func(){
   	 	gammas=ParalleUVW(cT,'w',pf)
    	wg.Done()
    
    }()*/


    z1 := []*big.Int{big.NewInt(int64(1))}

    wg.Add(1)
	go func(){
		for i := 0; i < len(a)/2; i++ {
			z1 = pf.Mul(
				z1,
				[]*big.Int{
					pf.F.Neg(
						big.NewInt(int64(i+1))),
						big.NewInt(int64(1)),
				})

		}
	 wg.Done()
    }()

    z2 := []*big.Int{big.NewInt(int64(1))}

    wg.Add(1)
	go func(){
		for i := len(a)/2; i < len(a); i++ {
			z2 = pf.Mul(
				z2,
				[]*big.Int{
					pf.F.Neg(
						big.NewInt(int64(i+1))),
						big.NewInt(int64(1)),
				})

		}
	 wg.Done()
    }()



	/*
	z := []*big.Int{big.NewInt(int64(1))}

    wg.Add(1)
	go func(){
		for i := 0; i < len(a); i++ {
			z = pf.Mul(
				z,
				[]*big.Int{
					pf.F.Neg(
						big.NewInt(int64(i+1))),
						big.NewInt(int64(1)),
				})

		}
	 wg.Done()
    }()*/

    wg.Wait()

    z:=pf.Mul(z1,z2)

    for i := 0; i < len(aT); i++ {
    	alphas = append(alphas, alphasTemp[i])
	}


	for i := 0; i < len(bT); i++ {
    	betas = append(betas, betasTemp[i])
	}

	for i := 0; i < len(cT); i++ {
    	gammas = append(gammas, gammasTemp[i])
	}

	elapsed := time.Since(start)
	fmt.Println("Compute QAP done,used ",elapsed)

	f, err := os.Create("qap_time.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    msg := fmt.Sprintf("Compute QAP done,used ",elapsed)
        _, _= f.WriteString(msg)
	
	
	return alphas, betas, gammas, z
}

// CombinePolynomials combine the given polynomials arrays into one, also returns the P(x)
func (pf PolynomialField) CombinePolynomials(r []*big.Int, ap, bp, cp [][]*big.Int) ([]*big.Int, []*big.Int, []*big.Int, []*big.Int) {
	var ax []*big.Int
	for i := 0; i < len(r); i++ {
		m := pf.Mul([]*big.Int{r[i]}, ap[i])
		ax = pf.Add(ax, m)
	}
	var bx []*big.Int
	for i := 0; i < len(r); i++ {
		m := pf.Mul([]*big.Int{r[i]}, bp[i])
		bx = pf.Add(bx, m)
	}
	var cx []*big.Int
	for i := 0; i < len(r); i++ {
		m := pf.Mul([]*big.Int{r[i]}, cp[i])
		cx = pf.Add(cx, m)
	}

	px := pf.Sub(pf.Mul(ax, bx), cx)
	return ax, bx, cx, px
}

// DivisorPolynomial returns the divisor polynomial given two polynomials
func (pf PolynomialField) DivisorPolynomial(px, z []*big.Int) []*big.Int {
	quo, _ := pf.Div(px, z)
	return quo
}

func Check_r1cs(wires []*big.Int,u,v,w [][]*big.Int,polyf PolynomialField )(bool){

	correct:=true
	fmt.Println("\nnumbers of gates:",len(u))

	for j :=0;  j<len(u);j++{
    	    sum_u:=big.NewInt(int64(0))
    	    sum_v:=big.NewInt(int64(0))
    	    sum_w:=big.NewInt(int64(0))
            for i := 0; i < len(wires); i++ {
    	        sum_u=polyf.F.Add(polyf.F.Mul(wires[i],u[j][i]),sum_u)
    	        sum_v=polyf.F.Add(polyf.F.Mul(wires[i],v[j][i]),sum_v)
    	        sum_w=polyf.F.Add(polyf.F.Mul(wires[i],w[j][i]),sum_w)
    	        
	        }
	        temp:=polyf.F.Mul(sum_u,sum_v)	
	        if temp.Cmp(sum_w)!=0{

			fmt.Println("R1CS not correct with gates ",j)
			fmt.Println("R1CS not correct:u",u[j])
			fmt.Println("R1CS not correct:v",v[j])
			fmt.Println("R1CS not correct:w",w[j])
            correct=false
			
		}
    }
    return correct

}

func Check_QAP(wires []*big.Int,ux,vx,wx,u,v,w [][]*big.Int,polyf PolynomialField )(bool){

	correct:=true

	for j :=1;  j<len(u);j++{
    	    v_u:=big.NewInt(int64(0))
    	    v_v:=big.NewInt(int64(0))
    	    v_w:=big.NewInt(int64(0))
            for i := 0; i < len(wires); i++ {
    	        v_u=polyf.Eval(ux[i],big.NewInt(int64(j)))
    	        	if v_u.Cmp(polyf.F.Affine(u[j-1][i]))!=0{
			        fmt.Println("not correct ux",i,"the value is",v_u,"but should be:",u[j-1][i],"evaluated at",j)
                    ut:=Transpose(u)
                    fmt.Println("lagelange build on:",ut[i])
                    fmt.Println("check the lagelange:",polyf.Eval(polyf.LagrangeInterpolation(ut[i]),big.NewInt(int64(j))))
                    correct=false 
		            }
    	        v_v=polyf.Eval(vx[i],big.NewInt(int64(j)))
    	        	if v_v.Cmp(polyf.F.Affine(v[j-1][i]))!=0{
			        fmt.Println("not correct vx",i)
			        correct=false
		            }
		        v_w=polyf.Eval(wx[i],big.NewInt(int64(j)))
    	        	if v_w.Cmp(polyf.F.Affine(w[j-1][i]))!=0{
			        fmt.Println("not correct wx",i)
			        correct=false
		            }
    	        
	        }
	        	
    }
   
    return correct

}

package main

import(
	"fmt"
	"log"
    "os"
)


func main(){

	/* build the circuit for PriBank (simple version)
	 b1new=b1-v12u
	 b2new=b2+v12u
	 total=b1new+b2new
	 v12u=v12u1*2+v12u2
	 v12u*(v12u-1)=0
	*/

	
	code:=
	    `func main(private b1, private b2,private b3,private b4,`+
		`private b1new, private b2new, private b3new,private b4new,`+
		`private t1,private t2,private t3,private t4,`+
		`private v12u,private v13u,private v14u,`+
		`private v21u,private v23u,private v24u,`+
		`private v31u,private v32u, private v34u,`+
		`private v41u,private v42u, private v43u,`+
		//transaction of user1
		`private b1new1,private b1new2,private b1new3,private b1new4,private b1new5,private b1new6,private b1new7,private b1new8,`+
		`private b2new1,private b2new2,private b2new3,private b2new4,private b2new5,private b2new6,private b2new7,private b2new8,`+
		`private b3new1,private b3new2,private b3new3,private b3new4,private b3new5,private b3new6,private b3new7,private b3new8,`+
		`private b4new1,private b4new2,private b4new3,private b4new4,private b4new5,private b4new6,private b4new7,private b4new8,`+


		`private v12u1,private v12u2,private v12u3,private v12u4,private v12u5,private v12u6,private v12u7,private v12u8,`+
		`private v13u1,private v13u2,private v13u3,private v13u4,private v13u5,private v13u6,private v13u7,private v13u8,`+
		`private v14u1,private v14u2,private v14u3,private v14u4,private v14u5,private v14u6,private v14u7,private v14u8,`+

		`private v21u1,private v21u2,private v21u3,private v21u4,private v21u5,private v21u6,private v21u7,private v21u8,`+
		`private v23u1,private v23u2,private v23u3,private v23u4,private v23u5,private v23u6,private v23u7,private v23u8,`+
		`private v24u1,private v24u2,private v24u3,private v24u4,private v24u5,private v24u6,private v24u7,private v24u8,`+

		`private v31u1,private v31u2,private v31u3,private v31u4,private v31u5,private v31u6,private v31u7,private v31u8,`+
        `private v32u1,private v32u2,private v32u3,private v32u4,private v32u5,private v32u6,private v32u7,private v32u8,`+
        `private v34u1,private v34u2,private v34u3,private v34u4,private v34u5,private v34u6,private v34u7,private v34u8,`+

        
        `private v41u1,private v41u2,private v41u3,private v41u4,private v41u5,private v41u6,private v41u7,private v41u8,`+
        `private v42u1,private v42u2,private v42u3,private v42u4,private v42u5,private v42u6,private v42u7,private v42u8,`+
		`private v43u1,private v43u2,private v43u3,private v43u4,private v43u5,private v43u6,private v43u7,private v43u8,`+
		`public total,public d1,public d2,public d3,public d4):`+
		
		// check v12u, use variable z12b...
		`
		z12b1=v12u1*2
		z12b2=z12b1+v12u2
		z12b3=z12b2*2
		z12b4=z12b3+v12u3
		z12b5=z12b4*2
		z12b6=z12b5+v12u4
		z12b7= z12b6*2
		z12b8=z12b7+v12u5
		z12b9=z12b8*2
		z12b10=z12b9+v12u6
		z12b11=z12b10*2
		z12b12=z12b11+v12u7
		z12b13=z12b12*2
		z12b14=z12b13+v12u8
		equals(z12b14,v12u)`+
        //check range of v12u.bits are 1 or 0, use variable z12o..a z12o..b
		`
		value0=0+0
		z12o1a=1-v12u1
		z12o1b=z12o1a*v12u1
		equals(z12o1b,value0)

		z12o2a=1-v12u2
		z12o2b=z12o2a*v12u2
		equals(z12o2b,value0)

		z12o3a=1-v12u3
		z12o3b=z12o3a*v12u3
		equals(z12o3b,value0)

		z12o4a=1-v12u4
		z12o4b=z12o4a*v12u4
		equals(z12o4b,value0)

		z12o5a=1-v12u5
		z12o5b=z12o5a*v12u5
		equals(z12o5b,value0)

		z12o6a=1-v12u6
		z12o6b=z12o6a*v12u6
		equals(z12o6b,value0)

		z12o7a=1-v12u7
		z12o7b=z12o7a*v12u7
		equals(z12o7b,value0)

		z12o8a=1-v12u8
		z12o8b=z12o8a*v12u8
		equals(z12o8b,value0)`+
		//check v24u
		`
		z24b1=v24u1*2
		z24b2=z24b1+v24u2
		z24b3=z24b2*2
		z24b4=z24b3+v24u3
		z24b5=z24b4*2
		z24b6=z24b5+v24u4
		z24b7= z24b6*2
		z24b8=z24b7+v24u5
		z24b9=z24b8*2
		z24b10=z24b9+v24u6
		z24b11=z24b10*2
		z24b12=z24b11+v24u7
		z24b13=z24b12*2
		z24b14=z24b13+v24u8
		equals(z24b14,v24u)`+

		//check range of v24u.bits are 1 or 0, use variable z24o..a z24o..b
		`
		z24o1a=1-v24u1
		z24o1b=z24o1a*v24u1
		equals(z24o1b,value0)

		z24o2a=1-v24u2
		z24o2b=z24o2a*v24u2
		equals(z24o2b,value0)

		z24o3a=1-v24u3
		z24o3b=z24o3a*v24u3
		equals(z24o3b,value0)

		z24o4a=1-v24u4
		z24o4b=z24o4a*v24u4
		equals(z24o4b,value0)

		z24o5a=1-v24u5
		z24o5b=z24o5a*v24u5
		equals(z24o5b,value0)

		z24o6a=1-v24u6
		z24o6b=z24o6a*v24u6
		equals(z24o6b,value0)

		z24o7a=1-v24u7
		z24o7b=z24o7a*v24u7
		equals(z24o7b,value0)

		z24o8a=1-v24u8
		z24o8b=z24o8a*v24u8
		equals(z24o8b,value0)`+

		//check v13u
		`
		z13b1=v13u1*2
		z13b2=z13b1+v13u2
		z13b3=z13b2*2
		z13b4=z13b3+v13u3
		z13b5=z13b4*2
		z13b6=z13b5+v13u4
		z13b7= z13b6*2
		z13b8=z13b7+v13u5
		z13b9=z13b8*2
		z13b10=z13b9+v13u6
		z13b11=z13b10*2
		z13b12=z13b11+v13u7
		z13b13=z13b12*2
		z13b14=z13b13+v13u8
		equals(z13b14,v13u)`+
		//check range of v13u.bits are 1 or 0, use variable z13o..a z13o..b
		`
		z13o1a=1-v13u1
		z13o1b=z13o1a*v13u1
		equals(z13o1b,value0)

		z13o2a=1-v13u2
		z13o2b=z13o2a*v13u2
		equals(z13o2b,value0)

		z13o3a=1-v13u3
		z13o3b=z13o3a*v13u3
		equals(z13o3b,value0)

		z13o4a=1-v13u4
		z13o4b=z13o4a*v13u4
		equals(z13o4b,value0)

		z13o5a=1-v13u5
		z13o5b=z13o5a*v13u5
		equals(z13o5b,value0)

		z13o6a=1-v13u6
		z13o6b=z13o6a*v13u6
		equals(z13o6b,value0)

		z13o7a=1-v13u7
		z13o7b=z13o7a*v13u7
		equals(z13o7b,value0)

		z13o8a=1-v13u8
		z13o8b=z13o8a*v13u8
		equals(z13o8b,value0)`+
		//check b1new
		`
		zb1newb1=b1new1*2
		zb1newb2=zb1newb1+b1new2
		zb1newb3=zb1newb2*2
		zb1newb4=zb1newb3+b1new3
		zb1newb5=zb1newb4*2
		zb1newb6=zb1newb5+b1new4
		zb1newb7= zb1newb6*2
		zb1newb8=zb1newb7+b1new5
		zb1newb9=zb1newb8*2
		zb1newb10=zb1newb9+b1new6
		zb1newb11=zb1newb10*2
		zb1newb12=zb1newb11+b1new7
		zb1newb13=zb1newb12*2
		zb1newb14=zb1newb13+b1new8
		equals(zb1newb14,b1new)`+
		//check range of b1new.bits are 1 or 0, use variable zb1newo..a zb1newo..b
		`
		zb1newo1a=1-b1new1
		zb1newo1b=zb1newo1a*b1new1
		equals(zb1newo1b,value0)

		zb1newo2a=1-b1new2
		zb1newo2b=zb1newo2a*b1new2
		equals(zb1newo2b,value0)

		zb1newo3a=1-b1new3
		zb1newo3b=zb1newo3a*b1new3
		equals(zb1newo3b,value0)

		zb1newo4a=1-b1new4
		zb1newo4b=zb1newo4a*b1new4
		equals(zb1newo4b,value0)

		zb1newo5a=1-b1new5
		zb1newo5b=zb1newo5a*b1new5
		equals(zb1newo5b,value0)

		zb1newo6a=1-b1new6
		zb1newo6b=zb1newo6a*b1new6
		equals(zb1newo6b,value0)

		zb1newo7a=1-b1new7
		zb1newo7b=zb1newo7a*b1new7
		equals(zb1newo7b,value0)

		zb1newo8a=1-b1new8
		zb1newo8b=zb1newo8a*b1new8
		equals(zb1newo8b,value0)`+

		//check b2new
		`
		zb2newb1=b2new1*2
		zb2newb2=zb2newb1+b2new2
		zb2newb3=zb2newb2*2
		zb2newb4=zb2newb3+b2new3
		zb2newb5=zb2newb4*2
		zb2newb6=zb2newb5+b2new4
		zb2newb7= zb2newb6*2
		zb2newb8=zb2newb7+b2new5
		zb2newb9=zb2newb8*2
		zb2newb10=zb2newb9+b2new6
		zb2newb11=zb2newb10*2
		zb2newb12=zb2newb11+b2new7
		zb2newb13=zb2newb12*2
		zb2newb14=zb2newb13+b2new8
		equals(zb2newb14,b2new)`+
		//check range of b2new.bits are 1 or 0, use variable zb2newo..a zb2newo..b
		`
		zb2newo1a=1-b2new1
		zb2newo1b=zb2newo1a*b2new1
		equals(zb2newo1b,value0)

		zb2newo2a=1-b2new2
		zb2newo2b=zb2newo2a*b2new2
		equals(zb2newo2b,value0)

		zb2newo3a=1-b2new3
		zb2newo3b=zb2newo3a*b2new3
		equals(zb2newo3b,value0)

		zb2newo4a=1-b2new4
		zb2newo4b=zb2newo4a*b2new4
		equals(zb2newo4b,value0)

		zb2newo5a=1-b2new5
		zb2newo5b=zb2newo5a*b2new5
		equals(zb2newo5b,value0)

		zb2newo6a=1-b2new6
		zb2newo6b=zb2newo6a*b2new6
		equals(zb2newo6b,value0)

		zb2newo7a=1-b2new7
		zb2newo7b=zb2newo7a*b2new7
		equals(zb2newo7b,value0)

		zb2newo8a=1-b2new8
		zb2newo8b=zb2newo8a*b2new8
		equals(zb2newo8b,value0)`+

		//check b3new
		`
		zb3newb1=b3new1*2
		zb3newb2=zb3newb1+b3new2
		zb3newb3=zb3newb2*2
		zb3newb4=zb3newb3+b3new3
		zb3newb5=zb3newb4*2
		zb3newb6=zb3newb5+b3new4
		zb3newb7= zb3newb6*2
		zb3newb8=zb3newb7+b3new5
		zb3newb9=zb3newb8*2
		zb3newb10=zb3newb9+b3new6
		zb3newb11=zb3newb10*2
		zb3newb12=zb3newb11+b3new7
		zb3newb13=zb3newb12*2
		zb3newb14=zb3newb13+b3new8
		equals(zb3newb14,b3new)`+
		//check range of b3new.bits are 1 or 0, use variable zb3newo..a zb3newo..b
		`
		zb3newo1a=1-b3new1
		zb3newo1b=zb3newo1a*b3new1
		equals(zb3newo1b,value0)

		zb3newo2a=1-b3new2
		zb3newo2b=zb3newo2a*b3new2
		equals(zb3newo2b,value0)

		zb3newo3a=1-b3new3
		zb3newo3b=zb3newo3a*b3new3
		equals(zb3newo3b,value0)

		zb3newo4a=1-b3new4
		zb3newo4b=zb3newo4a*b3new4
		equals(zb3newo4b,value0)

		zb3newo5a=1-b3new5
		zb3newo5b=zb3newo5a*b3new5
		equals(zb3newo5b,value0)

		zb3newo6a=1-b3new6
		zb3newo6b=zb3newo6a*b3new6
		equals(zb3newo6b,value0)

		zb3newo7a=1-b3new7
		zb3newo7b=zb3newo7a*b3new7
		equals(zb3newo7b,value0)

		zb3newo8a=1-b3new8
		zb3newo8b=zb3newo8a*b3new8
		equals(zb3newo8b,value0)`+

		//check b4new
		`
		zb4newb1=b4new1*2
		zb4newb2=zb4newb1+b4new2
		zb4newb3=zb4newb2*2
		zb4newb4=zb4newb3+b4new3
		zb4newb5=zb4newb4*2
		zb4newb6=zb4newb5+b4new4
		zb4newb7= zb4newb6*2
		zb4newb8=zb4newb7+b4new5
		zb4newb9=zb4newb8*2
		zb4newb10=zb4newb9+b4new6
		zb4newb11=zb4newb10*2
		zb4newb12=zb4newb11+b4new7
		zb4newb13=zb4newb12*2
		zb4newb14=zb4newb13+b4new8
		equals(zb4newb14,b4new)`+
		//check range of b4new.bits are 1 or 0, use variable zb4newo..a zb4newo..b
		`
		zb4newo1a=1-b4new1
		zb4newo1b=zb4newo1a*b4new1
		equals(zb4newo1b,value0)

		zb4newo2a=1-b4new2
		zb4newo2b=zb4newo2a*b4new2
		equals(zb4newo2b,value0)

		zb4newo3a=1-b4new3
		zb4newo3b=zb4newo3a*b4new3
		equals(zb4newo3b,value0)

		zb4newo4a=1-b4new4
		zb4newo4b=zb4newo4a*b4new4
		equals(zb4newo4b,value0)

		zb4newo5a=1-b4new5
		zb4newo5b=zb4newo5a*b4new5
		equals(zb4newo5b,value0)

		zb4newo6a=1-b4new6
		zb4newo6b=zb4newo6a*b4new6
		equals(zb4newo6b,value0)

		zb4newo7a=1-b4new7
		zb4newo7b=zb4newo7a*b4new7
		equals(zb4newo7b,value0)

		zb4newo8a=1-b4new8
		zb4newo8b=zb4newo8a*b4new8
		equals(zb4newo8b,value0)`+
		// check v14u, use variable z14b...
		`
		z14b1=v14u1*2
		z14b2=z14b1+v14u2
		z14b3=z14b2*2
		z14b4=z14b3+v14u3
		z14b5=z14b4*2
		z14b6=z14b5+v14u4
		z14b7= z14b6*2
		z14b8=z14b7+v14u5
		z14b9=z14b8*2
		z14b10=z14b9+v14u6
		z14b11=z14b10*2
		z14b12=z14b11+v14u7
		z14b13=z14b12*2
		z14b14=z14b13+v14u8
		equals(z14b14,v14u)`+
        //check range of v14u.bits are 1 or 0, use variable z14o..a z14o..b
		`
		z14o1a=1-v14u1
		z14o1b=z14o1a*v14u1
		equals(z14o1b,value0)

		z14o2a=1-v14u2
		z14o2b=z14o2a*v14u2
		equals(z14o2b,value0)

		z14o3a=1-v14u3
		z14o3b=z14o3a*v14u3
		equals(z14o3b,value0)

		z14o4a=1-v14u4
		z14o4b=z14o4a*v14u4
		equals(z14o4b,value0)

		z14o5a=1-v14u5
		z14o5b=z14o5a*v14u5
		equals(z14o5b,value0)

		z14o6a=1-v14u6
		z14o6b=z14o6a*v14u6
		equals(z14o6b,value0)

		z14o7a=1-v14u7
		z14o7b=z14o7a*v14u7
		equals(z14o7b,value0)

		z14o8a=1-v14u8
		z14o8b=z14o8a*v14u8
		equals(z14o8b,value0)`+

		//check v21u 
		`
		z21b1=v21u1*2
		z21b2=z21b1+v21u2
		z21b3=z21b2*2
		z21b4=z21b3+v21u3
		z21b5=z21b4*2
		z21b6=z21b5+v21u4
		z21b7= z21b6*2
		z21b8=z21b7+v21u5
		z21b9=z21b8*2
		z21b10=z21b9+v21u6
		z21b11=z21b10*2
		z21b12=z21b11+v21u7
		z21b13=z21b12*2
		z21b14=z21b13+v21u8
		equals(z21b14,v21u)`+
		//check range of v21u.bits are 1 or 0, use variable z21o..a z21o..b
		`
		z21o1a=1-v21u1
		z21o1b=z21o1a*v21u1
		equals(z21o1b,value0)

		z21o2a=1-v21u2
		z21o2b=z21o2a*v21u2
		equals(z21o2b,value0)

		z21o3a=1-v21u3
		z21o3b=z21o3a*v21u3
		equals(z21o3b,value0)

		z21o4a=1-v21u4
		z21o4b=z21o4a*v21u4
		equals(z21o4b,value0)

		z21o5a=1-v21u5
		z21o5b=z21o5a*v21u5
		equals(z21o5b,value0)

		z21o6a=1-v21u6
		z21o6b=z21o6a*v21u6
		equals(z21o6b,value0)

		z21o7a=1-v21u7
		z21o7b=z21o7a*v21u7
		equals(z21o7b,value0)

		z21o8a=1-v21u8
		z21o8b=z21o8a*v21u8
		equals(z21o8b,value0)`+
        
		// check v23u, use variable z23b...
		`
		z23b1=v23u1*2
		z23b2=z23b1+v23u2
		z23b3=z23b2*2
		z23b4=z23b3+v23u3
		z23b5=z23b4*2
		z23b6=z23b5+v23u4
		z23b7=z23b6*2
		z23b8=z23b7+v23u5
		z23b9=z23b8*2
		z23b10=z23b9+v23u6
		z23b11=z23b10*2
		z23b12=z23b11+v23u7
		z23b13=z23b12*2
		z23b14=z23b13+v23u8
		equals(z23b14,v23u)`+
		//check range of v23u.bits are 1 or 0, use variable z23o..a z23o..b
		`
		z23o1a=1-v23u1
		z23o1b=z23o1a*v23u1
		equals(z23o1b,value0)

		z23o2a=1-v23u2
		z23o2b=z23o2a*v23u2
		equals(z23o2b,value0)

		z23o3a=1-v23u3
		z23o3b=z23o3a*v23u3
		equals(z23o3b,value0)

		z23o4a=1-v23u4
		z23o4b=z23o4a*v23u4
		equals(z23o4b,value0)

		z23o5a=1-v23u5
		z23o5b=z23o5a*v23u5
		equals(z23o5b,value0)`+

		//check v31u
		`
		zv31ub1=v31u1*2
		zv31ub2=zv31ub1+v31u2
		zv31ub3=zv31ub2*2
		zv31ub4=zv31ub3+v31u3
		zv31ub5=zv31ub4*2
		zv31ub6=zv31ub5+v31u4
		zv31ub7= zv31ub6*2
		zv31ub8=zv31ub7+v31u5
		zv31ub9=zv31ub8*2
		zv31ub10=zv31ub9+v31u6
		zv31ub11=zv31ub10*2
		zv31ub12=zv31ub11+v31u7
		zv31ub13=zv31ub12*2
		zv31ub14=zv31ub13+v31u8
		equals(zv31ub14,v31u)`+
		//check range of v31u.bits are 1 or 0, use variable zv31uo..a zv31uo..b
		`
		zv31uo1a=1-v31u1
		zv31uo1b=zv31uo1a*v31u1
		equals(zv31uo1b,value0)

		zv31uo2a=1-v31u2
		zv31uo2b=zv31uo2a*v31u2
		equals(zv31uo2b,value0)

		zv31uo3a=1-v31u3
		zv31uo3b=zv31uo3a*v31u3
		equals(zv31uo3b,value0)

		zv31uo4a=1-v31u4
		zv31uo4b=zv31uo4a*v31u4
		equals(zv31uo4b,value0)

		zv31uo5a=1-v31u5
		zv31uo5b=zv31uo5a*v31u5
		equals(zv31uo5b,value0)

		zv31uo6a=1-v31u6
		zv31uo6b=zv31uo6a*v31u6
		equals(zv31uo6b,value0)

		zv31uo7a=1-v31u7
		zv31uo7b=zv31uo7a*v31u7
		equals(zv31uo7b,value0)

		zv31uo8a=1-v31u8
		zv31uo8b=zv31uo8a*v31u8
		equals(zv31uo8b,value0)`+

		//check v32u
		`
		zv32ub1=v32u1*2
		zv32ub2=zv32ub1+v32u2
		zv32ub3=zv32ub2*2
		zv32ub4=zv32ub3+v32u3
		zv32ub5=zv32ub4*2
		zv32ub6=zv32ub5+v32u4
		zv32ub7= zv32ub6*2
		zv32ub8=zv32ub7+v32u5
		zv32ub9=zv32ub8*2
		zv32ub10=zv32ub9+v32u6
		zv32ub11=zv32ub10*2
		zv32ub12=zv32ub11+v32u7
		zv32ub13=zv32ub12*2
		zv32ub14=zv32ub13+v32u8
		equals(zv32ub14,v32u)`+
		//check range of v32u.bits are 1 or 0, use variable zv32uo..a zv32uo..b
		`
		zv32uo1a=1-v32u1
		zv32uo1b=zv32uo1a*v32u1
		equals(zv32uo1b,value0)

		zv32uo2a=1-v32u2
		zv32uo2b=zv32uo2a*v32u2
		equals(zv32uo2b,value0)

		zv32uo3a=1-v32u3
		zv32uo3b=zv32uo3a*v32u3
		equals(zv32uo3b,value0)

		zv32uo4a=1-v32u4
		zv32uo4b=zv32uo4a*v32u4
		equals(zv32uo4b,value0)

		zv32uo5a=1-v32u5
		zv32uo5b=zv32uo5a*v32u5
		equals(zv32uo5b,value0)

		zv32uo6a=1-v32u6
		zv32uo6b=zv32uo6a*v32u6
		equals(zv32uo6b,value0)

		zv32uo7a=1-v32u7
		zv32uo7b=zv32uo7a*v32u7
		equals(zv32uo7b,value0)

		zv32uo8a=1-v32u8
		zv32uo8b=zv32uo8a*v32u8
		equals(zv32uo8b,value0)`+

		//check v34u
		`
		z34b1=v34u1*2
		z34b2=z34b1+v34u2
		z34b3=z34b2*2
		z34b4=z34b3+v34u3
		z34b5=z34b4*2
		z34b6=z34b5+v34u4
		z34b7= z34b6*2
		z34b8=z34b7+v34u5
		z34b9=z34b8*2
		z34b10=z34b9+v34u6
		z34b11=z34b10*2
		z34b12=z34b11+v34u7
		z34b13=z34b12*2
		z34b14=z34b13+v34u8
		equals(z34b14,v34u)`+

		//check range of v34u.bits are 1 or 0, use variable z34o..a z34o..b
		`
		z34o1a=1-v34u1
		z34o1b=z34o1a*v34u1
		equals(z34o1b,value0)

		z34o2a=1-v34u2
		z34o2b=z34o2a*v34u2
		equals(z34o2b,value0)

		z34o3a=1-v34u3
		z34o3b=z34o3a*v34u3
		equals(z34o3b,value0)

		z34o4a=1-v34u4
		z34o4b=z34o4a*v34u4
		equals(z34o4b,value0)

		z34o5a=1-v34u5
		z34o5b=z34o5a*v34u5
		equals(z34o5b,value0)

		z34o6a=1-v34u6
		z34o6b=z34o6a*v34u6
		equals(z34o6b,value0)

		z34o7a=1-v34u7
		z34o7b=z34o7a*v34u7
		equals(z34o7b,value0)

		z34o8a=1-v34u8
		z34o8b=z34o8a*v34u8
		equals(z34o8b,value0)`+
		// check v43u, use variable z43b...
		`
		z43b1=v43u1*2
		z43b2=z43b1+v43u2
		z43b3=z43b2*2
		z43b4=z43b3+v43u3
		z43b5=z43b4*2
		z43b6=z43b5+v43u4
		z43b7= z43b6*2
		z43b8=z43b7+v43u5
		z43b9=z43b8*2
		z43b10=z43b9+v43u6
		z43b11=z43b10*2
		z43b12=z43b11+v43u7
		z43b13=z43b12*2
		z43b14=z43b13+v43u8
		equals(z43b14,v43u)`+

		//check v41u
		`
		z41b1=v41u1*2
		z41b2=z41b1+v41u2
		z41b3=z41b2*2
		z41b4=z41b3+v41u3
		z41b5=z41b4*2
		z41b6=z41b5+v41u4
		z41b7= z41b6*2
		z41b8=z41b7+v41u5
		z41b9=z41b8*2
		z41b10=z41b9+v41u6
		z41b11=z41b10*2
		z41b12=z41b11+v41u7
		z41b13=z41b12*2
		z41b14=z41b13+v41u8
		equals(z41b14,v41u)`+

		//check range of v41u.bits are 1 or 0, use variable z41o..a z41o..b
		`
		z41o1a=1-v41u1
		z41o1b=z41o1a*v41u1
		equals(z41o1b,value0)

		z41o2a=1-v41u2
		z41o2b=z41o2a*v41u2
		equals(z41o2b,value0)

		z41o3a=1-v41u3
		z41o3b=z41o3a*v41u3
		equals(z41o3b,value0)

		z41o4a=1-v41u4
		z41o4b=z41o4a*v41u4
		equals(z41o4b,value0)

		z41o5a=1-v41u5
		z41o5b=z41o5a*v41u5
		equals(z41o5b,value0)

		z41o6a=1-v41u6
		z41o6b=z41o6a*v41u6
		equals(z41o6b,value0)

		z41o7a=1-v41u7
		z41o7b=z41o7a*v41u7
		equals(z41o7b,value0)

		z41o8a=1-v41u8
		z41o8b=z41o8a*v41u8
		equals(z41o8b,value0)`+

		//check v42u
		`
		z42b1=v42u1*2
		z42b2=z42b1+v42u2
		z42b3=z42b2*2
		z42b4=z42b3+v42u3
		z42b5=z42b4*2
		z42b6=z42b5+v42u4
		z42b7= z42b6*2
		z42b8=z42b7+v42u5
		z42b9=z42b8*2
		z42b10=z42b9+v42u6
		z42b11=z42b10*2
		z42b12=z42b11+v42u7
		z42b13=z42b12*2
		z42b14=z42b13+v42u8
		equals(z42b14,v42u)`+

		//check range of v42u.bits are 1 or 0, use variable z42o..a z42o..b
		`
		z42o1a=1-v42u1
		z42o1b=z42o1a*v42u1
		equals(z42o1b,value0)

		z42o2a=1-v42u2
		z42o2b=z42o2a*v42u2
		equals(z42o2b,value0)

		z42o3a=1-v42u3
		z42o3b=z42o3a*v42u3
		equals(z42o3b,value0)

		z42o4a=1-v42u4
		z42o4b=z42o4a*v42u4
		equals(z42o4b,value0)

		z42o5a=1-v42u5
		z42o5b=z42o5a*v42u5
		equals(z42o5b,value0)

		z42o6a=1-v42u6
		z42o6b=z42o6a*v42u6
		equals(z42o6b,value0)

		z42o7a=1-v42u7
		z42o7b=z42o7a*v42u7
		equals(z42o7b,value0)

		z42o8a=1-v42u8
		z42o8b=z42o8a*v42u8
		equals(z42o8b,value0)`+
        //check range of v43u.bits are 1 or 0, use variable z43o..a z43o..b
		`
		z43o1a=1-v43u1
		z43o1b=z43o1a*v43u1
		equals(z43o1b,value0)

		z43o2a=1-v43u2
		z43o2b=z43o2a*v43u2
		equals(z43o2b,value0)

		z43o3a=1-v43u3
		z43o3b=z43o3a*v43u3
		equals(z43o3b,value0)

		z43o4a=1-v43u4
		z43o4b=z43o4a*v43u4
		equals(z43o4b,value0)

		z43o5a=1-v43u5
		z43o5b=z43o5a*v43u5
		equals(z43o5b,value0)

		z43o6a=1-v43u6
		z43o6b=z43o6a*v43u6
		equals(z43o6b,value0)

		z43o7a=1-v43u7
		z43o7b=z43o7a*v43u7
		equals(z43o7b,value0)

		z43o8a=1-v43u8
		z43o8b=z43o8a*v43u8
		equals(z43o8b,value0)`+
		//check update u1, substitution use u1bSub... addition use u1bAdd...
		`
		u1bSub1=b1-v12u
		u1bSub2=u1bSub1-v13u
		u1bSub3=u1bSub2-v14u


		u1bAdd1=u1bSub3+v21u
		u1bAdd2=u1bAdd1+v31u
		u1bAdd3=u1bAdd2+v41u
		equals(u1bAdd3,b1new)
		`+

        //check update u2, substitution use u2bSub... addition use u2bAdd...
		`
		u2bSub1 = b2-v21u
		u2bSub2 = u2bSub1-v23u
		u2bSub3 = u2bSub2-v24u

		u2bAdd1=u2bSub3+v12u
		u2bAdd2=u2bAdd1+v32u
		u2bAdd3=u2bAdd2+v42u
		equals(u2bAdd3,b2new)`+

        //check update u3, substitution use u3bSub... addition use u3bAdd...
		`
		u3bSub1 = b3-v31u
        u3bSub2 = u3bSub1-v32u
        u3bSub3 = u2bSub2-v34u 

		u3bAdd1 = u3bSub3+v13u
		u3bAdd2=u3bAdd1+v23u
		u3bAdd3=u3bAdd2+v43u
		equals(u3bAdd3,b3new)`+
		
        `
		u4bSub1 = b3-v43u
        u4bSub2 = u4bSub1-v42u
        u4bSub3 = u2bSub2-v41u 

		u4bAdd1 = u4bSub3+v14u
		u4bAdd2=u4bAdd1+v24u
		u4bAdd3=u4bAdd2+v34u
		equals(u4bAdd3,b3new)`+

        //check total balance use variable B
		`
		B1=b1new+b2new 
		B2=B1+b3new
		B3=B2+b4new
		equals(B3,total)`+
        //mask relation
		`
		m1=d1-t1
		equals(m1,b1new)
		m2=d2-t2
		equals(m2,b2new)
		m3=d3-t3
		equals(m3,b3new)
		m4=d4-t4
		equals(m4,b4new)
		out = 1 * 1
	`
		
	fmt.Print("\nBuild the circuit:",code)

	f, err := os.Create("cir.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(code)

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("done")

    


}
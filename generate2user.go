package main

import(
	"fmt"
	"log"
    "os"
)


func main(){

	/* build the circuit for PriBank (simple version)
	 b1new=b1-v12
	 b2new=b2+v12
	 total=b1new+b2new
	 v12=v121*2+v122
	 v12*(v12-1)=0
	*/

	
	code:=
	    `func main(private b1, private b2,private b3,private b4,`+
		`private b1new, private b2new, private b3new,private b4new,`+
		`private t1,private t2,private t3,private t4,`+
		`private v12,private v13,private v14,`+
		`private v21,private v23,private v24u,`+
		`private v31u,private v32u,`+
		`private v43,`+
		//transaction of user1
		`private b1new1,private b1new2,private b1new3,private b1new4,private b1new5,private b1new6,private b1new7,private b1new8,`+
		`private b2new1,private b2new2,private b2new3,private b2new4,private b2new5,private b2new6,private b2new7,private b2new8,`+
		`private b3new1,private b3new2,private b3new3,private b3new4,private b3new5,private b3new6,private b3new7,private b3new8,`+
		`private b4new1,private b4new2,private b4new3,private b4new4,private b4new5,private b4new6,private b4new7,private b4new8,`+


		`private v121,private v122,private v123,private v124,private v125,private v126,private v127,private v128,`+
		`private v131,private v132,private v133,private v134,private v135,private v136,private v137,private v138,`+
		`private v141,private v142,private v143,private v144,private v145,private v146,private v147,private v148,`+

		//`private v211,private v212,private v213,private v214,private v215,private v216,private v217,private v218,`+
		//`private v231,private v232,private v233,private v234,private v235,private v236,private v237,private v238,`+
		//`private v24u1,private v24u2,private v24u3,private v24u4,private v24u5,private v24u6,private v24u7,private v24u8,`+

		//`private v31u1,private v31u2,private v31u3,private v31u4,private v31u5,private v31u6,private v31u7,private v31u8,`+
       // `private v32u1,private v32u2,private v32u3,private v32u4,private v32u5,private v32u6,private v32u7,private v32u8,`+

		//`private v431,private v432,private v433,private v434,private v435,private v436,private v437,private v438,`+
		`public total,public d1,public d2,public d3,public d4):`+
		
		// check v12, use variable z12b...
		`
		z12b1=v121*2
		z12b2=z12b1+v122
		z12b3=z12b2*2
		z12b4=z12b3+v123
		z12b5=z12b4*2
		z12b6=z12b5+v124
		z12b7= z12b6*2
		z12b8=z12b7+v125
		z12b9=z12b8*2
		z12b10=z12b9+v126
		z12b11=z12b10*2
		z12b12=z12b11+v127
		z12b13=z12b12*2
		z12b14=z12b13+v128
		equals(z12b14,v12)`+
        //check range of v12.bits are 1 or 0, use variable z12o..a z12o..b
		`
		value0=0+0
		z12o1a=1-v121
		z12o1b=z12o1a*v121
		equals(z12o1b,value0)

		z12o2a=1-v122
		z12o2b=z12o2a*v122
		equals(z12o2b,value0)

		z12o3a=1-v123
		z12o3b=z12o3a*v123
		equals(z12o3b,value0)

		z12o4a=1-v124
		z12o4b=z12o4a*v124
		equals(z12o4b,value0)

		z12o5a=1-v125
		z12o5b=z12o5a*v125
		equals(z12o5b,value0)

		z12o6a=1-v126
		z12o6b=z12o6a*v126
		equals(z12o6b,value0)

		z12o7a=1-v127
		z12o7b=z12o7a*v127
		equals(z12o7b,value0)

		z12o8a=1-v128
		z12o8b=z12o8a*v128
		equals(z12o8b,value0)`+/*
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
		equals(z24b14,v24u)`+/*

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
		equals(z24o8b,value0)`+*/

		//check v13
		`
		z13b1=v131*2
		z13b2=z13b1+v132
		z13b3=z13b2*2
		z13b4=z13b3+v133
		z13b5=z13b4*2
		z13b6=z13b5+v134
		z13b7= z13b6*2
		z13b8=z13b7+v135
		z13b9=z13b8*2
		z13b10=z13b9+v136
		z13b11=z13b10*2
		z13b12=z13b11+v137
		z13b13=z13b12*2
		z13b14=z13b13+v138
		equals(z13b14,v13)`+/*
		//check range of v13.bits are 1 or 0, use variable z13o..a z13o..b
		`
		z13o1a=1-v131
		z13o1b=z13o1a*v131
		equals(z13o1b,value0)

		z13o2a=1-v132
		z13o2b=z13o2a*v132
		equals(z13o2b,value0)

		z13o3a=1-v133
		z13o3b=z13o3a*v133
		equals(z13o3b,value0)

		z13o4a=1-v134
		z13o4b=z13o4a*v134
		equals(z13o4b,value0)

		z13o5a=1-v135
		z13o5b=z13o5a*v135
		equals(z13o5b,value0)

		z13o6a=1-v136
		z13o6b=z13o6a*v136
		equals(z13o6b,value0)

		z13o7a=1-v137
		z13o7b=z13o7a*v137
		equals(z13o7b,value0)

		z13o8a=1-v138
		z13o8b=z13o8a*v138
		equals(z13o8b,value0)`+*/
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
		// check v14, use variable z14b...
		`
		z14b1=v141*2
		z14b2=z14b1+v142
		z14b3=z14b2*2
		z14b4=z14b3+v143
		z14b5=z14b4*2
		z14b6=z14b5+v144
		z14b7= z14b6*2
		z14b8=z14b7+v145
		z14b9=z14b8*2
		z14b10=z14b9+v146
		z14b11=z14b10*2
		z14b12=z14b11+v147
		z14b13=z14b12*2
		z14b14=z14b13+v148
		equals(z14b14,v14)`+/*
        //check range of v14.bits are 1 or 0, use variable z14o..a z14o..b
		`
		z14o1a=1-v141
		z14o1b=z14o1a*v141
		equals(z14o1b,value0)

		z14o2a=1-v142
		z14o2b=z14o2a*v142
		equals(z14o2b,value0)

		z14o3a=1-v143
		z14o3b=z14o3a*v143
		equals(z14o3b,value0)

		z14o4a=1-v144
		z14o4b=z14o4a*v144
		equals(z14o4b,value0)

		z14o5a=1-v145
		z14o5b=z14o5a*v145
		equals(z14o5b,value0)

		z14o6a=1-v146
		z14o6b=z14o6a*v146
		equals(z14o6b,value0)

		z14o7a=1-v147
		z14o7b=z14o7a*v147
		equals(z14o7b,value0)

		z14o8a=1-v148
		z14o8b=z14o8a*v148
		equals(z14o8b,value0)`+

		//check v21 
		`
		z21b1=v211*2
		z21b2=z21b1+v212
		z21b3=z21b2*2
		z21b4=z21b3+v213
		z21b5=z21b4*2
		z21b6=z21b5+v214
		z21b7= z21b6*2
		z21b8=z21b7+v215
		z21b9=z21b8*2
		z21b10=z21b9+v216
		z21b11=z21b10*2
		z21b12=z21b11+v217
		z21b13=z21b12*2
		z21b14=z21b13+v218
		equals(z21b14,v21)`+/*
		//check range of v21.bits are 1 or 0, use variable z21o..a z21o..b
		`
		z21o1a=1-v211
		z21o1b=z21o1a*v211
		equals(z21o1b,value0)

		z21o2a=1-v212
		z21o2b=z21o2a*v212
		equals(z21o2b,value0)

		z21o3a=1-v213
		z21o3b=z21o3a*v213
		equals(z21o3b,value0)

		z21o4a=1-v214
		z21o4b=z21o4a*v214
		equals(z21o4b,value0)

		z21o5a=1-v215
		z21o5b=z21o5a*v215
		equals(z21o5b,value0)

		z21o6a=1-v216
		z21o6b=z21o6a*v216
		equals(z21o6b,value0)

		z21o7a=1-v217
		z21o7b=z21o7a*v217
		equals(z21o7b,value0)

		z21o8a=1-v218
		z21o8b=z21o8a*v218
		equals(z21o8b,value0)`+
        
		// check v23, use variable z23b...
		`
		z23b1=v231*2
		z23b2=z23b1+v232
		z23b3=z23b2*2
		z23b4=z23b3+v233
		z23b5=z23b4*2
		z23b6=z23b5+v234
		z23b7=z23b6*2
		z23b8=z23b7+v235
		z23b9=z23b8*2
		z23b10=z23b9+v236
		z23b11=z23b10*2
		z23b12=z23b11+v237
		z23b13=z23b12*2
		z23b14=z23b13+v238
		equals(z23b14,v23)`+/*
		//check range of v23.bits are 1 or 0, use variable z23o..a z23o..b
		`
		z23o1a=1-v231
		z23o1b=z23o1a*v231
		equals(z23o1b,value0)

		z23o2a=1-v232
		z23o2b=z23o2a*v232
		equals(z23o2b,value0)

		z23o3a=1-v233
		z23o3b=z23o3a*v233
		equals(z23o3b,value0)

		z23o4a=1-v234
		z23o4b=z23o4a*v234
		equals(z23o4b,value0)

		z23o5a=1-v235
		z23o5b=z23o5a*v235
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
		equals(zv31ub14,v31u)`+/*
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
		equals(zv32ub14,v32u)`+/*
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
		// check v43, use variable z43b...
		`
		z43b1=v431*2
		z43b2=z43b1+v432
		z43b3=z43b2*2
		z43b4=z43b3+v433
		z43b5=z43b4*2
		z43b6=z43b5+v434
		z43b7= z43b6*2
		z43b8=z43b7+v435
		z43b9=z43b8*2
		z43b10=z43b9+v436
		z43b11=z43b10*2
		z43b12=z43b11+v437
		z43b13=z43b12*2
		z43b14=z43b13+v438
		equals(z43b14,v43)`+/*
        //check range of v43.bits are 1 or 0, use variable z43o..a z43o..b
		`
		z43o1a=1-v431
		z43o1b=z43o1a*v431
		equals(z43o1b,value0)

		z43o2a=1-v432
		z43o2b=z43o2a*v432
		equals(z43o2b,value0)

		z43o3a=1-v433
		z43o3b=z43o3a*v433
		equals(z43o3b,value0)

		z43o4a=1-v434
		z43o4b=z43o4a*v434
		equals(z43o4b,value0)

		z43o5a=1-v435
		z43o5b=z43o5a*v435
		equals(z43o5b,value0)

		z43o6a=1-v436
		z43o6b=z43o6a*v436
		equals(z43o6b,value0)

		z43o7a=1-v437
		z43o7b=z43o7a*v437
		equals(z43o7b,value0)

		z43o8a=1-v438
		z43o8b=z43o8a*v438
		equals(z43o8b,value0)`+*/
		//check update u1, substitution use u1bSub... addition use u1bAdd...
		`
		u1bSub1=b1-v12
		u1bSub2=u1bSub1-v13
		u1bSub3=u1bSub2-v14


		u1bAdd1=u1bSub3+v21
		u1bAdd2=u1bAdd1+v31u
		equals(u1bAdd2,b1new)
		`+

        //check update u2, substitution use u2bSub... addition use u2bAdd...
		`
		u2bSub1 = b2-v21
		u2bSub2 = u2bSub1-v23
		u2bSub3 = u2bSub2-v24u

		u2bAdd1=u2bSub3+v12
		u2bAdd2=u2bAdd1+v32u
		equals(u2bAdd2,b2new)`+

        //check update u3, substitution use u3bSub... addition use u3bAdd...
		`
		u3bSub1 = b3-v31u
        u3bSub2 = u3bSub1-v32u

		u3bAdd1 = b3+v23
		u3bAdd2=u3bAdd1+v43
		equals(u3bAdd2,b3new)

		s5 =b4-v43
		equals(s5,b4new)`+

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
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
		`private v23,private v43,`+
		//transaction of user1
		`private v121,private v122,private v123,private v124,private v125,private v126,private v127,private v128,`+
		`private v131,private v132,private v133,private v134,private v135,private v136,private v137,private v138,`+
		`private v141,private v142,private v143,private v144,private v145,private v146,private v147,private v148,`+

		`private v231,private v232,private v233,private v234,private v235,private v236,private v237,private v238,`+
		`private v431,private v432,private v433,private v434,private v435,private v436,private v437,private v438,`+
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
		equals(z12o8b,value0)`+
		//check v13
		`z13b1=v131*2
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
		equals(z13b14,v13)`+
		//check range of v13.bits are 1 or 0, use variable z13o..a z13o..b
		`
		value0=0+0
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
		equals(z13o8b,value0)`+
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
		equals(z14b14,v14)`+
        //check range of v14.bits are 1 or 0, use variable z14o..a z14o..b
		`
		value0=0+0
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
		equals(z23b14,v23)`+
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
		equals(z43b14,v43)`+
        //check range of v43.bits are 1 or 0, use variable z43o..a z43o..b
		`
		value0=0+0
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
		equals(z43o8b,value0)`+
		//check update u1
		`
		u1bSub1=b1-v12
		u1bSub2=u1bSub1-v13
		u1bSub3=u1bSub2-v14
		equals(u1bSub3,b1new)
		`+

        //check total balance
		`
		s1 = b2+v12
		s2 = s1-v23
		equals(s2,b2new)
		s3 = b3+v23
		s4=s3+v43
		equals(s4,b3new)
		s5 =b4-v43
		equals(s5,b4new)`+
        //check total balance use variable B
		`B1=b1new+b2new 
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
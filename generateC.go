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
		`private v12,private v121,private v122,private v123,private v124,private v125,private v126,private v127,private v128,`+
		`private v23,private v231,private v232,private v233,private v234,private v235,`+
		`private v43,`+
		`public total):
		s0 = b1-v12
		equals(s0, b1new)
		s1 = b2+v12
		s2 = s1-v23
		equals(s2,b2new)
		s3 = b3+v23
		s4=s3+v43
		equals(s4,b3new)
		s5 =b4-v43
		equals(s5,b4new)`+
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

		`+
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
		equals(z23b8,v23)`+
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
        //check total balance use variable B
		`B1=b1new+b2new 
		B2=B1+b3new
		B3=B2+b4new
		equals(B3,total)
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
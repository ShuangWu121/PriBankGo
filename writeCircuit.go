package main

import (
    "fmt"
    "log"
    "os"
)

func main() {

    f, err := os.Create("circuit.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString("func main(")

    if err2 != nil {
        log.Fatal(err2)
    }

    users:=2
    balancesRange:=6
    transactionsRange:=4

    users=users+1
    balancesRange=balancesRange+1
    transactionsRange=transactionsRange+1
 
    //generate old balances b...
    for i:=1;i<users;i++{
        msg := fmt.Sprintf("private b%d,", i)
        _, err2 := f.WriteString(msg)

       if err2 != nil {
        log.Fatal(err2)
       }
    }
    
    //generate new balance b...new
     for i:=1;i<users;i++{
        msg := fmt.Sprintf("private b%dnew,", i)
        _, err2 := f.WriteString(msg)

       if err2 != nil {
        log.Fatal(err2)
       }
    }
    
    //generate randomness t...
     for i:=1;i<users;i++{
        msg := fmt.Sprintf("private t%d,", i)
        _, err2 := f.WriteString(msg)

       if err2 != nil {
        log.Fatal(err2)
       }
    }

    //generate transactions

    for i:=1;i<users;i++{

        for j:=1;j<users;j++{

            if j!=i{
                msg := fmt.Sprintf("private v%d%du,", i,j)
                 _, err2 := f.WriteString(msg)


                if err2 != nil {
                log.Fatal(err2)
            }
           }
        }
    }

    //generate bits for new balance b...new...
    for i:=1;i<users;i++{

        for j:=1;j<balancesRange;j++{
            msg := fmt.Sprintf("private b%dnew%d,", i,j)
            _, err2 := f.WriteString(msg)

            if err2 != nil {
            log.Fatal(err2)
            }
        }
    }
    //generate bits for trasactions v...u...
    for i:=1;i<users;i++{

        for j:=1;j<users;j++{

            if j!=i{

                for k:=1;k<transactionsRange;k++{
                    msg := fmt.Sprintf("private v%d%du%d,", i,j,k)
                    _, err2 := f.WriteString(msg)


                    if err2 != nil {
                        log.Fatal(err2)
                    }
                }
           }
        }
    }

    _, err2 = f.WriteString("public total,")

    if err2 != nil {
        log.Fatal(err2)
    }




    //generate d
    for i:=1;i<users-1;i++{
        msg := fmt.Sprintf("public d%d,", i)
        _, err2 := f.WriteString(msg)

       if err2 != nil {
        log.Fatal(err2)
       }
    }
    msg := fmt.Sprintf("public d%d):\n",users-1)
     _, err2 = f.WriteString(msg)

    if err2 != nil {
        log.Fatal(err2)
    }


    

    //check the sum of transaction bits
    
    for i:=1;i<users;i++{

        for j:=1;j<users;j++{

            if j!=i{
                s:=1
                for k:=1;k<transactionsRange;k++{
                    
                    if k==1{
                        msg := fmt.Sprintf("z%d%db%d=v%d%du%d*2 \n", i,j,s,i,j,k)
                        _, err2 := f.WriteString(msg)
                        if err2 != nil {
                        log.Fatal(err2)
                        }
                        s++
                        k++
                        msg = fmt.Sprintf("z%d%db%d=z%d%db%d+v%d%du%d \n", i,j,s,i,j,s-1,i,j,k)
                        _, err2 = f.WriteString(msg)
                    }else{
                        msg := fmt.Sprintf("z%d%db%d=z%d%db%d*2 \n", i,j,s,i,j,s-1)
                        _, err2 := f.WriteString(msg)
                        if err2 != nil {
                        log.Fatal(err2)
                        }
                        s++

                        msg = fmt.Sprintf("z%d%db%d=z%d%db%d+v%d%du%d \n", i,j,s,i,j,s-1,i,j,k)
                        _, err2 = f.WriteString(msg)

                    }
                    

                    s++
                    
                }
            msg := fmt.Sprintf("equals(z%d%db%d,v%d%du) \n \n", i,j,s-1,i,j)
            _, err2 := f.WriteString(msg)
            if err2 != nil {
                        log.Fatal(err2)
                    }
           }


        }
    }
    

     _, err2 = f.WriteString("value0=0+0\n")

    if err2 != nil {
        log.Fatal(err2)
    }

    //check transaction bits are 1 or 0
    
    for i:=1;i<users;i++{

        for j:=1;j<users;j++{

            if j!=i{

                for k:=1;k<transactionsRange;k++{
                    msg := fmt.Sprintf("z%d%do%da=1-v%d%du%d \n", i,j,k,i,j,k)
                    _, err2 := f.WriteString(msg)
                    if err2 != nil {
                        log.Fatal(err2)
                    }

                    msg = fmt.Sprintf("z%d%do%db=z%d%do%da*v%d%du%d \n", i,j,k,i,j,k,i,j,k)
                    _, err2 = f.WriteString(msg)
                    if err2 != nil {
                        log.Fatal(err2)
                    }

                    msg = fmt.Sprintf("equals(z%d%do%db,value0) \n", i,j,k)
                    _, err2 = f.WriteString(msg)
                    if err2 != nil {
                        log.Fatal(err2)
                    }


                }
              _, err2 = f.WriteString("\n")
           }
        }
    }
   
   //check sum of new balance bits

   for i:=1;i<users;i++{
        
        s:=1
        for k:=1;k<balancesRange;k++{
            if k==1{
                        msg := fmt.Sprintf("zb%dnewb%d=b%dnew%d*2 \n",i,s,i,k)
                        _, err2 := f.WriteString(msg)
                        if err2 != nil {
                        log.Fatal(err2)
                        }
                        s++
                        k++
                        msg = fmt.Sprintf("zb%dnewb%d=zb%dnewb%d+b%dnew%d \n", i,s,i,s-1,i,k)
                        _, err2 = f.WriteString(msg)
            }else{
                        msg := fmt.Sprintf("zb%dnewb%d=zb%dnewb%d*2 \n",i,s,i,s-1)
                        _, err2 := f.WriteString(msg)
                        if err2 != nil {
                        log.Fatal(err2)
                        }
                        s++

                        msg = fmt.Sprintf("zb%dnewb%d=zb%dnewb%d+b%dnew%d \n", i,s,i,s-1,i,k)
                        _, err2 = f.WriteString(msg)

                    }
                    

                    s++
        }

        msg := fmt.Sprintf("equals(zb%dnewb%d,b%dnew) \n \n", i,s-1,i)
            _, err2 := f.WriteString(msg)
            if err2 != nil {
                        log.Fatal(err2)
                    }


    }

   //check new balance bits are 0 or 1
   

   for i:=1;i<users;i++{


        for k:=1;k<balancesRange;k++{
            msg := fmt.Sprintf("zb%dnewo%da=1-b%dnew%d \n", i,k,i,k)
             _, err2 := f.WriteString(msg)
            if err2 != nil {
                log.Fatal(err2)
                 }

            msg = fmt.Sprintf("zb%dnewo%db=zb%dnewo%da*b%dnew%d \n", i,k,i,k,i,k)
             _, err2 = f.WriteString(msg)
            if err2 != nil {
                log.Fatal(err2)
            }

            msg = fmt.Sprintf("equals(zb%dnewo%db,value0) \n",i,k)
             _, err2 = f.WriteString(msg)
            if err2 != nil {
                log.Fatal(err2)
            }


            }
             _, err2 = f.WriteString("\n")
        
    }
    
    //check new balance updates
    

    for i:=1;i<users;i++{
        s:=1
        for j:=1;j<users;j++{
            if j!=i{
                
                if s==1{
                    msg := fmt.Sprintf("u%dSub%d=b%d-v%d%du \n",i,s,i,i,j)
                    _, err2 := f.WriteString(msg)
                    if err2 != nil {
                    log.Fatal(err2)
                    }
                }else{
                    msg := fmt.Sprintf("u%dSub%d=u%dSub%d-v%d%du \n",i,s,i,s-1,i,j)
                    _, err2 := f.WriteString(msg)
                    if err2 != nil {
                    log.Fatal(err2)
                    }
                }   
                s++

            }
        }
        _, err2 = f.WriteString("\n")
        
        
        t:=1
        for j:=1;j<users;j++{
            if j!=i{
                if t==1{
                    msg := fmt.Sprintf("u%dAdd1=u%dSub%d+v%d%du \n",i,i,s-1,j,i)
                    _, err2 = f.WriteString(msg)
                    if err2 != nil {
                    log.Fatal(err2)
                    }
                }else{
                    msg := fmt.Sprintf("u%dAdd%d=u%dAdd%d+v%d%du \n",i,t,i,t-1,j,i)
                    _, err2 = f.WriteString(msg)
                    if err2 != nil {
                    log.Fatal(err2)
                    }
                }


            t++
            }

        }
    msg := fmt.Sprintf("equals(u%dAdd%d,b%dnew) \n \n",i,t-1,i)
    _, err2 = f.WriteString(msg)
    if err2 != nil {
    log.Fatal(err2)
    }

    }
    

   //check total balance
    s:=1
    for i:=1;i<users;i++{
        if s==1{
            msg := fmt.Sprintf("B1=b1new+b2new\n")
            _, err2 = f.WriteString(msg)
            if err2 != nil {
            log.Fatal(err2)
            }
            i++
        }else{
            msg := fmt.Sprintf("B%d=B%d+b%dnew \n",s,s-1,i)
            _, err2 = f.WriteString(msg)
            if err2 != nil {
            log.Fatal(err2)
            }
        }
        s++
    }
   
    msg = fmt.Sprintf("equals(B%d,total) \n \n",s-1)
    _, err2 = f.WriteString(msg)
    if err2 != nil {
    log.Fatal(err2)
    }

  //check randomness and mask
    for i:=1;i<users;i++{
        msg := fmt.Sprintf("m%d=d%d-t%d\n",i,i,i)
        _, err2 = f.WriteString(msg)
        if err2 != nil {
        log.Fatal(err2)
        }
        msg = fmt.Sprintf("equals(m%d,b%dnew) \n",i,i)
        _, err2 = f.WriteString(msg)
        if err2 != nil {
        log.Fatal(err2)
        }


    }
   
    msg = fmt.Sprintf("out = 1 * 1 \n")
    _, err2 = f.WriteString(msg)
    if err2 != nil {
    log.Fatal(err2)
    }
        
    

    fmt.Println("done")
}
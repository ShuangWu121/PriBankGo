# PriBankGo
An implmentation for PriBank core algorithm in Go

#1 Set user number
 in file writeCircuit.go, set the number of users, users balance range and transaction values range by:
 
 `users:=2`
 
 `balancesRange:=4`
 
 `transactionsRange:=2
`

The numbers indicate the bit length of the value

The setting needs to match the setting in file pribank.go
 
Note: transaction range needs to be less than balance range, otherwise it is very easy to get overflow. 

for example, if the maximum balance value is 7, 3 users, and the maximum transaction value is 7 as well, two users send 7 to the third user, it will cause the overflow. When overflow happens, the transaction will be set to 0.

#Generate circuit
go run writeCircuit.go

the circuit description is in file circuit.txt

#Run 

go run pribank.go
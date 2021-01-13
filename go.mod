module pribank.go

go 1.15

require (
	github.com/ShuangWu121/PriBankGo/circuitcompiler v0.0.0-00010101000000-000000000000 // indirect
	github.com/ShuangWu121/PriBankGo/r1csqap v0.0.0-00010101000000-000000000000 // indirect
	github.com/ShuangWu121/PriBankGo/zkproof v0.0.0 // indirect
	github.com/ShuangWu121/secp256k1 v0.0.0-20180413221153-00116ff8c62f // indirect
	github.com/arnaucube/go-snark v0.0.4 // indirect
	github.com/ethereum/go-ethereum v1.9.25 // indirect
)

replace github.com/ShuangWu121/PriBankGo/zkproof v0.0.0 => ../PriBankGo/zkproof

replace github.com/ShuangWu121/secp256k1 v0.0.0-20180413221153-00116ff8c62f => ../PriBankGo/secp256k1

replace github.com/ShuangWu121/PriBankGo/circuitcompiler => ../PriBankGo/circuitcompiler

replace github.com/ShuangWu121/PriBankGo/r1csqap => ../PriBankGo/r1csqap

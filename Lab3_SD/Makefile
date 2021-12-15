run89:	runleia runfulcrum0

run90:	runinfo runfulcrum1

run91:	runinfo runfulcrum2

run92:	runbroker

runinfo:
	go run Informante_client/main.go
	
runbroker:
	go run Broker_Mos_Eisley/main.go

runleia:
	go run Leia_client/main.go

runfulcrum0:
	go run Fulcrum_server/main.go 0 &
	
runfulcrum1:
	go run Fulcrum_server/main.go 1 &
	
runfulcrum2:
	go run Fulcrum_server/main.go 2 &
run89:	runfulcrum0 runleia

run90:	runfulcrum1 runinfo

run91:	runfulcrum2 runinfo

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
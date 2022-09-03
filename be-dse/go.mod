module be-dse

go 1.19

require be-dse/transport v0.0.0 // indirect

require be-dse/datastruct v0.0.0 // indirect
require be-dse/logging v0.0.0 // indirect

require be-dse/service v0.0.0 // indirect

require be-dse/config v0.0.0 // indirect

require be-dse/router v0.0.0

replace be-dse/transport => ./transport

replace be-dse/datastruct => ./datastruct

replace be-dse/logging => ./logging

replace be-dse/service => ./service

replace be-dse/config => ./config

replace be-dse/router => ./router

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
)

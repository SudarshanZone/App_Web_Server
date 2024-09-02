package config

import (
	"github.com/go-ini/ini"
)

type Config struct {
	ServerAddress       string    //Webserver  
	GRPCServerAddress1  string 	 // FNO_OPen
	GRPCServerAddress2  string	 // FNO_Order_DTLS
	GRPCServerAddress3  string	 // Commodities_Open_Pos
	GRPCServerAddress4 string	// MTF_Open_Pos
	GRPCServerAddress5 string  //Equity_Main_Open 
}

func LoadConfig(fileName string) (*Config, error) {
	cfg, err := ini.Load(fileName)
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerAddress:      cfg.Section("").Key("server_address").MustString(":8080"),
		GRPCServerAddress1: cfg.Section("").Key("grpc_server_address_1").MustString("localhost:50051"),	
		GRPCServerAddress2: cfg.Section("").Key("grpc_server_address_2").MustString("localhost:50052"),
		GRPCServerAddress3: cfg.Section("").Key("grpc_server_address_3").MustString("localhost:50053"), 
		GRPCServerAddress4: cfg.Section("").Key("grpc_server_address_4").MustString("localhost:50054"),
		GRPCServerAddress5: cfg.Section("").Key("grpc_server_address_5").MustString("localhost:50055"), 
	}, nil
}

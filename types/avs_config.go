package types

type NodeConfig struct {
	// used to set the logger level (true = info, false = debug)
	Production                    bool   `yaml:"production"`
	OperatorAddress               string `yaml:"operator_address"`
	OperatorStateRetrieverAddress string `yaml:"operator_state_retriever_address"`
	AVSRegistryCoordinatorAddress string `yaml:"avs_registry_coordinator_address"`
	TokenStrategyAddr             string `yaml:"token_strategy_addr"`
	EthRpcUrl                     string `yaml:"eth_rpc_url"`
	EthWsUrl                      string `yaml:"eth_ws_url"`
	BlsPrivateKeyStorePath        string `yaml:"bls_private_key_store_path"`
	EcdsaPrivateKeyStorePath      string `yaml:"ecdsa_private_key_store_path"`
	AggregatorServerIpPortAddress string `yaml:"aggregator_server_ip_port_address"`
	RegisterOperatorOnStartup     bool   `yaml:"register_operator_on_startup"`
	EigenMetricsIpPortAddress     string `yaml:"eigen_metrics_ip_port_address"`
	EnableMetrics                 bool   `yaml:"enable_metrics"`
	NodeApiIpPortAddress          string `yaml:"node_api_ip_port_address"`
	EnableNodeApi                 bool   `yaml:"enable_node_api"`
}

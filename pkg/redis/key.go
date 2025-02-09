package redis

import "fmt"

func ChainEventListenKey(name string) string {
	return fmt.Sprintf("chain_event_listen:%v", name)
}

func MoveChainCheckInEventSeqKey(name string) string {
	return fmt.Sprintf("move_chain_check_in_event_listen:%v", name)
}

func MoveChainSignalPredictionEventSeqKey(name string) string {
	return fmt.Sprintf("move_chain_signal_prediction_event_listen:%v", name)
}

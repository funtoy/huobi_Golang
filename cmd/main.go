package main

import (
	"github.com/funtoy/huobi_golang/cmd/accountclientexample"
	"github.com/funtoy/huobi_golang/cmd/commonclientexample"
	"github.com/funtoy/huobi_golang/cmd/crossmarginclientexample"
	"github.com/funtoy/huobi_golang/cmd/etfclientexample"
	"github.com/funtoy/huobi_golang/cmd/isolatedmarginclientexample"
	"github.com/funtoy/huobi_golang/cmd/marketclientexample"
	"github.com/funtoy/huobi_golang/cmd/marketwebsocketclientexample"
	"github.com/funtoy/huobi_golang/cmd/orderclientexample"
	"github.com/funtoy/huobi_golang/cmd/walletclientexample"
	"github.com/funtoy/huobi_golang/logging/perflogger"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	//commonclientexample.RunAllExamples()
	//accountclientexample.RunAllExamples()
	//orderclientexample.RunAllExamples()
	//algoorderclientexample.RunAllExamples()
	//marketclientexample.RunAllExamples()
	//isolatedmarginclientexample.RunAllExamples()
	//crossmarginclientexample.RunAllExamples()
	//walletclientexample.RunAllExamples()
	//subuserclientexample.RunAllExamples()
	//stablecoinclientexample.RunAllExamples()
	//etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	//accountwebsocketclientexample.RunAllExamples()
	//orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}

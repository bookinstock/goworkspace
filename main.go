package main

import (
	"fmt"
	"goworkspace/app/goroutine/ants"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main run")

	// ctrlflow.Run()
	// collection.Run()
	// object.Run()

	// ---

	// foo.Run()

	// variable.Run()

	// str.Run()
	// str.RunRegex()

	// cond.Run()

	// cond.RunIf()
	// cond.RunSwitch()

	// loop.Run()

	// list.Run()
	//
	// collection.Run()

	// hashmap.Run()

	// loop.RunRange()

	// fn.Run()
	// fn.RunLambda()
	// fn.RunRecursion()

	// oo.RunStruct()
	// oo.RunInterface()
	// oo.RunEmbedding()

	// error.Run()
	// error.RunPanic()

	// concurrent.Run()
	// concurrent.RunWorker()
	// concurrent.RunRange()
	// concurrent.RunSync()
	// concurrent.RunRateLimit()

	// json.Run()

	// sort.Run()

	// dt.Run()
	// dt.RunEpoch()
	// dt.RunParse()

	// num.RunRand()
	// num.RunParse()

	// http.RunUrl()

	// str.RunSha1()
	// str.RunBase64()

	// file.Run()

	// redisx.RunExample1()
	// redisx.Run()

	ants.Example1()
	ants.Example2()
}

package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Coordinator struct {
	// Your definitions here.

}

// Your code here -- RPC handlers for the worker to call.

// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

var FileList = [...]string{"being_ernest", "dorian_gray", "frankenstein", "grimm", "huckleberry_finn", "metamorphosis", "sherlock_holmes", "tom_sawyer"}
var DoneFlag = 1

func (c *Coordinator) MyExample(args *MyArgs, reply *MyReply) (string, error) {
	// if args.Seq >= 7 {
	// 	DoneFlag = 0
	// 	return "wrong", errors.New("all txt has been get")
	// }
	reply.FileName = FileList[args.Seq]
	return reply.FileName, nil
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {
	var ret bool
	if DoneFlag == 0 {
		ret = true
	} else if DoneFlag == 1 {
		ret = false
	}
	// Your code here.

	return ret
}

// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Your code here.

	c.server()
	return &c
}

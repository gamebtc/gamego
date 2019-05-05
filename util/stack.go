package util

import (
	"context"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// 产生panic时的调用栈打印
func PrintPanicStack(extras ...interface{}) {
	if x := recover(); x != nil {
		log.Error(x)
		i := 0
		funcName, file, line, ok := runtime.Caller(i)
		for ok {
			log.Errorf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			i++
			funcName, file, line, ok = runtime.Caller(i)
		}

		for k := range extras {
			log.Errorf("EXRAS#%v DATA:%v\n", k, spew.Sdump(extras[k]))
		}
	}
}


//https://www.cnblogs.com/YYRise/p/10797794.html

func CpuProfile(ctx context.Context) {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("CPU Profile started")
	pprof.StartCPUProfile(f)
	go func(){
		select{
		case <-ctx.Done():
			pprof.StopCPUProfile()
			f.Close()
		}
	}()
}

func HeapProfile(ctx context.Context) {
	f, err := os.Create("heap.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.WriteHeapProfile(f)
	go func(){
		select{
		case <-ctx.Done():
			f.Close()
		}
	}()
}

func TraceProfile(ctx context.Context) {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Trace started")
	trace.Start(f)
	go func(){
		select{
		case <-ctx.Done():
			trace.Stop()
			f.Close()
		}
	}()
}
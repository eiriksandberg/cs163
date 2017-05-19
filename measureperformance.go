package main

import (
    "os"
    "runtime/trace"
    "time"
    "flag"
    "fmt"
    "runtime"
    "strconv"
    "sync"
)

var minDepth = 5
var n = 0

func binaryTree() {
   runtime.GOMAXPROCS(runtime.NumCPU() * 2)

   flag.Parse()
   if flag.NArg() > 0 {
      n, _ = strconv.Atoi(flag.Arg(0))
   }

   maxDepth := n
   if minDepth+2 > n {
      maxDepth = minDepth + 2
   }
   stretchDepth := maxDepth + 1

   check_l := bottomUpTree(stretchDepth).ItemCheck()
   fmt.Printf("stretch tree of depth %d\t check: %d\n", stretchDepth, check_l)

   longLivedTree := bottomUpTree(maxDepth)

   result_trees := make([]int, maxDepth+1)
   result_check := make([]int, maxDepth+1)

   var wg sync.WaitGroup
   for depth_l := minDepth; depth_l <= maxDepth; depth_l += 2 {
      wg.Add(1)
      go func(depth int) {
         iterations := 1 << uint(maxDepth-depth+minDepth)
         check := 0

         for i := 1; i <= iterations; i++ {
            check += bottomUpTree(depth).ItemCheck()
         }
         result_trees[depth] = iterations
         result_check[depth] = check

         wg.Done()
      }(depth_l)
   }
   wg.Wait()

   for depth := minDepth; depth <= maxDepth; depth += 2 {
      fmt.Printf("%d\t trees of depth %d\t check: %d\n",
         result_trees[depth], depth, result_check[depth],
      )
   }
   fmt.Printf("long lived tree of depth %d\t check: %d\n",
      maxDepth, longLivedTree.ItemCheck(),
   )
}

func bottomUpTree(depth int) *Node {
   if depth <= 0 {
      return &Node{nil, nil}
   }
   return &Node{
      bottomUpTree(depth-1),
      bottomUpTree(depth-1),
   }
}

type Node struct {
   left, right *Node
}

func (self *Node) ItemCheck() int {
   if self.left == nil {
      return 1
   }
   return 1 + self.left.ItemCheck() + self.right.ItemCheck()
}


func doWork(c chan int) {
    startTime := time.Now()
    i := 0
    for curTime := startTime; curTime.Sub(startTime) < 2; curTime = time.Now() {
        binaryTree()
        i++
    }
    c <- i
}

func main() {
    numGoRoutine := 10
    traceOutFile, _ := os.OpenFile("/tmp/Trace.out", os.O_WRONLY|os.O_CREATE, os.ModeExclusive|os.ModePerm)
    trace.Start(traceOutFile)

    // Start goroutines
    termChannel := make(chan int)
    for i := 0; i < numGoRoutine; i++ {
        go doWork(termChannel)
    }

    // Wait for completion
    for i := 0; i < numGoRoutine; i++ {
        <-termChannel
    }

    trace.Stop()
}

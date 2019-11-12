### Go G-P-M模型

----
    1. 每一个OS线程都有一个固定大小的内存块(通常为2MB)来做栈，这个栈会用来存储当前正在被调用或挂起
       的函数的内部变量. Go语言做了自己的线程
   
----
    Goroutine
    1. gotoutine的栈采用动态扩容的方式，初始值仅为2KB。由golang自己的调度器Go Schedulerl来调度
    2. GC会周期性将不再使用的内存回收，收缩栈空间
    3. grountine通过Precessor(逻辑处理器)与os线程绑定， 
       P可以看作是一个抽象资源或上下文环境， 一个P绑定一个OS线程， golang中把OS线程抽象成一个数据
       结构M（Machine）, G实际由M通过P来进行调度的
    4. 从G的成面来看，P提供了G运行的上下文环境和资源，P就是它的"CPU"
    
----
    Go调度器基本结构
    · Goroutine: 对应一个G结构体，G存储Goroutine的运行堆栈，状态以及任务函数
    · P:Processor逻辑处理器，对G来讲P相当于CPU核，G只有在P的local runq(run队列中)才能被调度
        对M来讲，P提供了执行环境Context(如内存分配状态mcache, 任务队列G)， P的数量决定系统最大并行G的数量
        （前提：物理CPU核数 >= P的数量），P的数量由用户设置的GOMAXPROCS决定，但是不论GOMAXPROCS设置为多大，P的数量最大为256
    · M：Machine(OS线程抽象)，绑定有效的P之后，进入schedule循环；
        schedule循环的机制大致是从Global队列、P的local队列以及wait队列中获取G，切换到G的执行栈上执行G的函数
        调用goexit做清理工作并回到M，如此反复。M并不保留G状态，这是G可以跨M调度的基础，M的数量是不定的，
        由Go Runtime调整，为了防止创建过多OS线程导致系统调度不过来，目前默认最大限制为10000个
        
----
    work-stealing调度算法
    * 每个P维护一个G的本地队列
    * 当一个G被创建，或变为可执行状态时，把它放到P的可执行队列中
    * 当一个G在M里执行结束后，P会从队列中把改G取出； 如果P的队列为空，M就随机选择另一个P，从其可执行的G队列中取走一半
![](../img/GPM.png)

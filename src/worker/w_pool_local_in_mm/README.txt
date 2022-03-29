Imagine you have a high load, io-heavy task and a long execution time
If you want the http response as soon as possible,
And you want the job to execute at the workerPool with high speed and use the variables in the current code in the simplest way,
This is the right tool.

It uses local in memory save data job, ringbuffer to share data mutil worker, a fast, lightweight tool, 100% golang, not another tool

It is recommended to use different worker pools for tasks that are bound by different resources, or that have different resource use patterns.
For example, tasks that use X Mb of memory may need different concurrency limits than tasks that use Y Mb of memory.
===> in short, it is recommended to create a workerPool pool for each worker specificity.


Life circle:
 1) Init list poolWorker, startWorker, worker live parallel webServer
 2) Push job to poolJob from callBack Fc
 3) PoolWorker auto get and handle job in PoolJob
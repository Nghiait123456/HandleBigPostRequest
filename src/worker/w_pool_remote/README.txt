this is remote worker pool use redis server

life circle:
 1) init worker( concurrency on routine, middleware, job, ....)
 2) worker still live and get job, listen info from system
 2) keep file main.go live for tool (supervisor)
 3) push job to poolJob, workerPool auto get job and run

note:
 if use in local memory, init redis local server
 every server webapp, you have one redis corresponding
 it's a resource-consumption worth considering, be careful when using it
 (best pratice: use local in memory, use share memory, global variable golang)

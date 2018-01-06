* locks comes up when programs run concurrently
* when two threads or processes or whatever try to access 1 piece of mem
* can run programs in parallel to increase throughput
* can also run programs in parallel to increase availability
* active vs. passive redundancy
* requests in a load balancer aren't served by all processes but LB selects one
* what if they aren't requests in a LB?
* use a lock and have the processes compete for the lock to see who can do the work
* similar to pulling jobs from a queue
* most of the time you are working on data in a database
* most databases have some mechanism for locking
* examples: pg, dynamo

# Practical Locking

Locks, mutexes, shared memory, and semaphores can make you feel bad about computers. We need them to do things efficiently.

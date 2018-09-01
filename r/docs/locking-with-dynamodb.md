# Locking With DynamoDB

September 4th 2012

What makes DynamoDB stand out is the fact that it is highly available and offered as an amaozon service --you can setup a lock service in minutes. DynamoDB offers a simple set of primitives that make locking possible:

* [Strongly Consistent Reads](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/APISummary.html#DataReadConsistency)
* [Conditional Writes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.Modifying.html#Expressions.Modifying.ConditionalWrites)

Since DynamoDB offers conditional writes, we can attempt to lock an item by creating the item unless an item already exists. Once we are done with the lock, we delete it.

### Gotchas

When using DynamoDB, we trade lock expiration robustness for ease of use. This trade-off is manifested when we consider what would happen if a processor acquires a lock, then is partitioned from DynamoDB. The process of detecting and deleting stale locks is a difficult problem. Solutions like Zookeeper use consensus algorithms to agree if a lock has been abounded before deleting. However, there are some use cases which do not require this level of sophistication.

One work around is to set a TTL on the lock item in DynamoDB. Subsequent attempts to acquire the lock will check the TTL and preempt an abandoned lock. This implies that the process that acquired the lock needs to terminate execution past the TTL. This can be achieved by wrapping your critical section in a timeout.

### How you can start using DynamoDB as a lock

The recipe for using DynamoDB as a locking service is quite simple, you do not need anything but and HTTP client to take full advantage of locking on top of DynamoDB. The basic algorithm is:

    Prune expired lock
    Create Item unless exists
    Yield to critical section inside of timeout
    Delete Item

## Conclusion

DynamoDB is not the most advanced locking service; however, if offers a lot of features at a great price. Most importantly, it's offered as a service. If you are already running in AWS it is worth a look.

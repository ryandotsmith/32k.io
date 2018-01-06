Event Buffering
===============

Eventually platforms outgrow the single-source-tree model and become
distributed systems. A common pattern in these distributed systems is
[distributed
composition](https://github.com/heroku/engineering-docs/blob/master/distributed-composition.md)
via event buffering. Here we motivate and describe this event buffering
pattern.

### The Problem

We have system that involves a client and a server. The client is an
HTTP API that accepts requests from PAAS customers. Customers will want
to scale their processes, transfer apps to other owners and so on...

The Server provides a app-group-membership API that takes an app id and
a group id and defines a relationship between the two objects. This
allows the customer to place the app into different billing groups. The
relationship is activated and the de-activated when the app is moved in
and out of groups. Apps and groups can be activated and de-activated
infinitely many times but an app can only belong to one group at a time.
The client who is accepting & allowing these movements is not
responsible for keeping history of these movements; instead it relies on
the Server for long term durability.

Let us now take a look at several approaches we can take to send data
from our client to the server.

### Example 1 - Inline (Synchronous)

**Table 1.1**

    Ephemeral Client Service

    App | Group | Time
    ----+-------+------
    A1  | G1    | 001
    A1  | G2    | 002
    A1  | G3    | 003
    A1  | G2    | 004

Table 1.1 represents app-group relationships as they exist in the real
world. These relationships were created on behalf of a customer via the
client's API. Since the client is not responsible for long-term
durability, we need to send this data to our server. Assuming there is a
function in the client to create a new relationship, we could make a
call the server in-line with this function:

    function createRelationship(app, group) {
      t = new Date()
      if app.update({"group_id": group.id}) {
        if serverWrite({"group_id": group.id, "app_id": app.id, "time": t}) {
          return true;
        } else {
          return false;
        }
      } else {
        return false
      }
    }

There are a few issues with this approach. For instance, what if server
is offline when we attempt to write the record? Should we not allow the
customer's request to complete? Also, what if another function is
calling *createRelationship()* and this function opened a database
transaction and needs to roll it back?

If this approach is taken, users should be aware that the server is now
a single point of failure for the system.

However, this approach does yield some benefits. When in-lining the
service calls, our system will produce back-pressure when components
degrade. This may be desirable if buffer bloat will put the system in a
place in which it can never keep up with load. Also, some system will
need to act on the result of the service call. Handling the service call
asynchronously is simply not feasible for these situations.

### Example 2 - Message Queue (Asynchronous)

In order to provide a more robust experience for customers of this
system, we could make the writes to our server in an external process.
This implies that if the server is down we will still be able to service
the customer's request. We can also specify retry logic in the event
that our attempts to persist the message to the server fail. Finally, we
can delay the write to the server until after all local database
transactions have been committed.

    function createRelationship(app, group) {
      t = new Date()
      if app.update({"group_id": group.id}) {
        enqueue(serverWrite({"group_id": group.id, "app_id": app.id, "time": t}));
        return true;
      } else {
        return false;
      }
    }

While this approach is a great improvement in comparison to Example 1,
there are still some issues. The first is the new operational burden we
have introduced to the *createRelationship()* function. We now depend on
at least 2 external services to complete the *createRelationship()*: our
local database and our message queue. We also depend on our message
queue being durable; if our queue begins dropping messages, we run the
risk of silently corrupting data.

**Message Queue Worker Process**

We will need some sort of worker process to work the jobs in this
message queue. Modern queues like resque and
[queue\_classic](https://github.com/ryandotsmith/queue_classic) come
with some sort of worker script out of the box.

### Example 3 - Event Buffer (Asynchronous)

The event buffer is a local, database-backed construct that removes a
message queue from the critical path. Instead of depending on a message
queue inside of our *createRelationship()* function, we write a row to
our buffer table:

    function createRelationship(app, group) {
      t = new Date();
      params = {"group_id": group.id, "app_id": app.id, "time": t};

      database.transaction.begin()
      database.relationships.write(params)
      database.buffer.write(params)
      if database.transaction.commit() {
        return true;
      } else {
        return false;
      }
    }

Notice how both of our writes are inside of a local database
transaction.

One interesting note about this approach is that we are keeping two
tables that hold the same type of information. At first this may seem
superfluous. However, since our client must only be concerned with the
current state of the system, we need a table to store historical records
until our async process can reliably deliver them to the server. To
reiterate this approach, consider the following scenario:

Say that a app oscillates between two groups. Then the following events
would have transpired in the real world.

    App | Group | Time
    ----+-------+------
    A1  | G1    | 001
    A1  | G2    | 002
    A1  | G1    | 003
    A1  | G2    | 004

At time 004 the client's relationship table will look like this:

    App | Group
    ----+-------
    A1  | G2

And the client's buffer table will look like this:

    App | Group | Time | Recorded
    ----+-------+------+---------
    A1  | G1    | 001  | 002
    A1  | G2    | 002  | 002
    A1  | G1    | 003  | 004
    A1  | G2    | 004  | 004

Notice that all of the events in the buffer are candidate for garbage
collection.

The event buffer also provides operational metrics. The implementer of
this approach can write database queries that show the number of
backlogged events. Since the backlog is stored in relational table, the
reporting mechanisms can be joined with other relations for debugging
scenarios.

**Buffer Processor**

Unlike the message queue approach, the event buffer approach requires a
little more effort with the async process. The implementer of this
approach will have to build a process that scans the buffer looking for
records to publish. Depending on the data, you can monitor the process
with Heroku scheduler or keep a process online the polls on certain
intervals.

    function processBuffer() {
      records = database.buffer.exec("SELECT * FROM buffer WHERE recorded IS NULL");
      for (var i = 0; i < records.length; i++) {
        record = records[i];
        if serverWrite({"group_id": record.group_id, "app_id": record.app_id, "time": record.time}) {
          t = new Date()
          record.update({"recorded": t})
        }
      }
    }
    setTimeout(processBuffer, 1000);

### Conclusion

We have seen several examples of how to transfer state from our client
to our server. The primary reason that we take these steps to transfer
state is to eliminate the number of services in our distributed system
that have to maintain state. Keeping a database on a service eventually
becomes and operational hazard. As tables grow in size, backups become
non-trivial, replication becomes slow, insert times can degrade
throughput, etc... For these reasons, it is desirable to isolate state
to one service in your networks of systems. Now that we have clear
patterns as to how we can reliably transfer state to durable services,
the programmers of client systems need not worry about maintaining the
state; only transferring it to a durable service.

Perhaps in another article we can discuss patterns in maintaining large
amounts of state. (Possible title: The state of state)

### Further Reading

-   [Harvest, Yield, and Scalable Tolerant
    Systems](http://radlab.cs.berkeley.edu/people/fox/static/pubs/pdf/c18.pdf)
-   [Life beyond Distributed
    Transaction](http://www.ics.uci.edu/~cs223/papers/cidr07p15.pdf)

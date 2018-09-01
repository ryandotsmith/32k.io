Tales From a Heroku User
========================

I originally presented this material in a talk I gave to the students at [Hack Reactor](http://hackreactor.com/). This may help explain the terse style used in the following paragraphs.

About
-----

-   [@ryandotsmith](https://twitter.com/ryandotsmith)
-   Engineer at Heroku for almost 3 years.
-   Programming since university.
-   Studied Math.
-   Built armature radios in junior high.

Things I have worked on
-----------------------

-   The Heroku API
-   Message queues (e.g.
    [queue\_classic](https://github.com/ryandotsmith/queue_classic))
-   Heroku's usage & billing system
-   DNS at Heroku
-   Log delivery and visibility tooling (e.g.
    [l2met](https://github.com/ryandotsmith/l2met),
    [log-shuttle](https://github.com/ryandotsmith/log-shuttle))

Technologies that I spend considerable time with
------------------------------------------------

-   Heroku
-   PostgreSQL
-   Redis
-   Ruby
-   Go
-   UNIX

Heroku
------

Things Heroku makes trivial:

-   Keeping a process alive.
-   Keeping your data alive.
-   Routing HTTP requests to more than one process.
-   Adding new processes and new apps.
-   SSL Termination.
-   Operating System Maintenance.
-   Log routing.
-   Security.

At one time in our history, everyone made their own electricity.

![](http://upload.wikimedia.org/wikipedia/commons/1/14/DynamoElectricMachinesEndViewPartlySection_USP284110.png)

A dynamo that would power a business. The business would employ people
to build and maintain the dynamo. Dyanmos were abandoned as they were
neither safe or economical. I feel that this is similar to servers and
Heroku.

How Heroku works
----------------

This is a great time for questions. Stop me at any point and we can take
a deep dive into a particular concept or system.

![](http://f.cl.ly/items/1J2h12402n0t3z0p1R1d/heroku-arch.png)

How to design your app knowing how Heroku works
-----------------------------------------------

### Stateless & Uncoordinated

-   Processes are distributed
-   Your database might be distributed
-   l2met example

![](http://f.cl.ly/items/2c3l2T3L3W2w140P3i29/l2met.png)

### Web Processes are Precious

-   Handle work in the background
-   worker pattern example

![](http://f.cl.ly/items/3L0B3Y203o131e1D1C16/worker%20pattern.png)

12 Factor
---------

[12factor.net](http://12factor.net/)

-   Good for any app
-   Forced if running on Heroku

Workflow
--------

I find that Heroku is most powerful when you learn to embrace the UNIX
toolkit.

### Config

There are no configuration files. Files are clunky and hard to manage,
the unix way of dealing with config is to use the UNIX environment
memory.

```bash
$ man env
```

Grabbing the config from your Heroku app.

```bash
$ heroku config -s > .env
```

Loading your config locally.

```bash
$ export $(cat .env)
```

Copying config from one app to another.

```bash
$ heroku config:set $(heroku config -s)
```

Log Analysis
------------

Don't think of your logs as static files. Think of the logs as live
streams of your app's health and status.

```bash
$ ruby -e '$stdout.puts("at=\"hack-reactor\"")'
```

Counting the number of HTTP requests per second.

```bash
$ heroku logs -t -p router | pv -lr >/dev/null
```

Sampling your response times.

```bash
$ heroku logs -t -p router | awk '{print $10,$11,$12}'
```

Using tools like log-shuttle and l2met we can build charts of system
metrics.

![](http://f.cl.ly/items/1q412q0C1V1Z3L093p3k/Screen%20Shot%202013-02-08%20at%203.15.17%20PM.png)

Interacting with your database
------------------------------

A database provided by the Heroku addons catalog is just a database.
With Heroku Postgresql, you can use standard UNIX tools to interact with
it.

Connecting to your database.

```bash
$ export url=`heroku config -s | grep "DATABASE_URL" | sed 's/DATABASE_URL=//'`
$ psql $url
```

Using your editor to build queries.

```bash
$ export EDITOR=vi
$ psql $url
database - \e
```

Similar feature available in bash.
```bash
$ echo 'hello world' | grep 'hello' | sed 's/hello/bye'
$ fc
```

Writing query results to your local filesystem.

```bash
$ echo 'select 1 \g out.txt' | psql $url
$ cat out.txt
```

Getting Real
------------

It is clear that Heroku is the best place to develop your application,
but it is also a great place to run and scale your application. A couple
of things to keep in mind when you grow your traffic and your uptime.

![](http://f.cl.ly/items/3k3v3N26081j0l2S0G2g/scaling.png)

-   Choose the right ratio of work between your process and your
    database.
-   Loading functions in your database is OK. Be aware that it won't
    scale horizontally.
-   Scaling horizontall can help increase throughput when your latencies
    are high.
-   Think about latency vs. throughput and how to balance the two.

[Maximizing the uptime of your Heroku
app](https://devcenter.heroku.com/articles/maximizing-availability)

![](http://f.cl.ly/items/2o3T1k293N3O0x3N1y10/Screen%20Shot%202013-02-08%20at%203.40.44%20PM.png)

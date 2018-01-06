Announcing l2met
================

Turn these: `$stdout.puts("measure#db.latency=4ms")`

Into this:

![](http://f.cl.ly/items/2R0h1x1b3V0Y0z0l1t1n/Screen%20Shot%202013-07-30%20at%209.59.52%20PM.png)

As systems grow in complexity, it becomes increasingly important to know
the precise dynamics of each system and how these systems relate to your
business. Achieving this insight is most commonly done through metrics.
However, creating a rich set of metrics that integrate both system level
information and business level information can be challenging. Today's
solutions to these problems are lacking when it comes to developer
experience. L2met is unlike other solutions in that it is optimized for
an excellent developer experience. Since l2met is painless to maintain
and easy to use, you will create more metrics and consequently have more
insight into the operations of your systems and your business.

**You can have l2met up and running in less than a minute.** Give it a
try by following these [setup
instructions.](https://github.com/ryandotsmith/l2met#getting-started)
Keep reading for more details.

**Contents:**

-   [What is L2met](#what-is-l2met)
-   [Motivation](#motivation)
-   [Structured Log Data](#structured-log-data)
-   [Log Delivery](#log-delivery)
-   [L2met and Heroku](#l2met-and-heroku)
-   [History of L2met](#history-of-l2met)
-   [Source on Github](https://github.com/ryandotsmith/l2met)

What is L2met
-------------

L2met is a network daemon that receives log lines, extracts data from
the logs using a convention, and outlets the aggregation of data to
charting services like Graphite or Librato.

Motivation
----------

L2met was developed to address several of the visibility problems we
faced at Heroku:

-   Real time metrics on specific system indicators
-   Medium term trending on those indicators
-   Ability to create actionable alerts on metrics (e.g. pagerduty)
-   Source metrics from log data

Structured Log Data
-------------------

L2met extracts data from log lines and creates metrics as a result. This
means developers can create metrics without installing an external
library. As long as you can write to STDOUT you can create metrics. Here
is an example of a structured log message:

```
measure#cache.get=4ms measure.db.get=50ms measure#cache.put=4ms
```

There are a couple things to note: key=value structure, multiple metrics
on a single log line, and units associated with the values. L2met will
parse this log line and create 3 metics: cache.get, db.get, and
cache.put. The sum, count, mean, min, median, 95/99th percentile, max,
and last values will be computed on the metrics.

Log Delivery
------------

Since there are no client libraries for the app developer to install,
the host which runs the app must deliver the logs to l2met. If your app
is running on Heroku, this is as simple as adding a log drain:

```
$ heroku drains:add https://your-l2met -a your-app
```

Once you are draining into l2met, you simply put structured logs to STDOUT.

```
puts('measure#hello=1 measure#world=1')
```

If you manage your own infrastructure, you can use
[log-shuttle(1)](http://log-shuttle.io) to format the app logs into
syslog packets and forward the logs to l2met.

```
$ export LOGS_URL=https://u:p@l2met.herokuapp.com/logs
$ echo 'measure#hello=1 measure#world=1' | log-shuttle
```

In both cases, l2met will parse 2 metrics, hello=1 and world=1.

L2met and Heroku
----------------

Most of Heroku's internal services (dyno managers, API, etc.) log using
l2met conventions. Those logs are sent to
[logplex](https://devcenter.heroku.com/articles/logplex), our
distributed log routing cluster, which then drains each app's stream to
a fleet of l2met instances. L2met computes metrics from the incoming
logs, and outlets those metrics to Librato. We use the data from
Librato's API and l2met's (experimental) HTTP outlet to alert engineers
if operating thresholds are exceeded.

The use of [logplex](https://devcenter.heroku.com/articles/logplex) and
its drains make it super easy to setup metrics for both platform (apps
that run on Heroku) and kernel services. New services can have a full
fledged metric pipeline in a matter of seconds by typing one command
(\`heroku drains:add\`), provided they are following the logging
conventions.

![](http://f.cl.ly/items/2X2S0C093Y3d0m3k462V/log-arch.png)

History of L2met
----------------

L2met was built on the shoulders of giants. It is a fresh implementation
that leverages modern technologies and solutions to a well known and
rigorously studied and solved problem. The visibility problem has
plagued many engineers for many years and consequently there have been
many great solutions. Here is a brief timeline of the problem to which I
am most familiar with:

-   2001-02-25 RRDtool
-   2008-06-02 Graphite
-   2010-12-29 Statsd
-   2011-01-25 Pulse\*
-   2012-01-05 Wcld\*
-   2012-02-13 Exprd\*
-   2012-08-07 l2met\*

\* Solutions used within Heroku.

While l2met has no direct relationship with RRD or Graphite, it does
recognize that backend storage and visualizations software is tough
business. The goal of l2met was to leverage the storage and
visualization platforms while providing the simplest input interface for
application developers. This was achieved by the fact that application
developers need only to put to STDOUT to collect metrics. This is after
all something every programmer can and has done already. If you squint
your eyes, l2met basically asks programmers to write a hello world
program.

Statsd was a great inspiration. However, there were certain
implementation details that made it unsuitable for our particular set of
engineering goals. The reliance of UDP as a transport and the lack of
proper UDP routing support on the Heroku platform meant that we were not
going to be able to run Statsd on the Heroku platform.

[Pulse](https://github.com/heroku/pulse) was the next generation
visibility solution for Heroku. It differed from Statsd in that it
parsed data from a log stream delivered by syslog-ng. Users of
[Pulse](https://github.com/heroku/pulse) would register functions in the
code base that would analyze the incoming logs. A customer web interface
was added to [Pulse](https://github.com/heroku/pulse) to report on the
numbers in real time.

[Wcld](https://github.com/ryandotsmith/wcld) varied a bit from
[Pulse](https://github.com/heroku/pulse) in that it didn't require users
to deploy code to acquire new metrics. It parsed and stored data in
hstore columns inside a PostgreSQL database. You could then write
programs that connected to the database and executed queries and the
output of the queries would be sent to a Librato account.

Exprd was the realization that there could be one program that received
and stored log data and also provided an interface to get the data out.
Users of exprd would drain their log data into it and then create
expressions that would be evaluated on the data over some interval that
was defined in the expression. The data would then be outleted to
Librato. After several production issues, we learned that a RDBMS was
not the best solution for the type of data and workload.

After dealing with several availability issues and a general lack of
multi tenancy, a new project was born. L2met was to be better than the
rest in that it would not require long term state. If a database were to
be used, it should be used in such a way that if the current database
failed, a new, fresh database could be swapped in place with little
effect. This implied that user expressions and configuration could not
be permitted. Instead, l2met defined a set of log conventions. If users
logged in a loosely defined way, they would be rewarded with metrics
with the smallest possible footprint (writing to STDOUT) across all
platforms.

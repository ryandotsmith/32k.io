The Worker Pattern
==================

Contents
--------

-   Introduction
-   Definition
-   Examples
-   Links

Introduction
------------

The modern application developer has undoubtedly dealt with the problem
of process execution locality. More precisely, when responding to HTTP
requests, developers must carefully balance what work is to be done in
and outside the process handling the request. The Worker Pattern exists
to help provide a framework for thinking about the balance.

In this article, we will define the pattern and look at several
applications of the pattern. Similar to advanced mathematical
techniques, the Worker Pattern is simple yet the application is not
alway clear. However, with sufficient exposure to examples and with
enough practice, the pattern becomes a reflex for the application
developer.

Definition
----------

**Processor:** Something that can execute instructions.

**Group of Computation:** One or many atomic instructions.

**The Worker Pattern:** To divide a group of computations amongst a set
of processors.

Example: Processing HTTP Requests
---------------------------------

A web service that requires high throughput will undoubtedly need to
ensure low latency while processing requests. In other words, the
process that is serving HTTP requests should spend the least amount of
time possible to serve the request. Subsequently if the server does not
have all of the data necessary to properly respond to the request, it
must not wait until the data is found. Instead it must let the client
know that it is working on the fulfillment of the request and that the
client should check back later. Such an arrangement will guarantee that
our web servers are always available to respond to requests with low
latency.

The application of the Worker Pattern in this case involves moving the
fulfillment of the request to another process. Leaving the server process
free to respond to other requests. Let us know explore the algorithm:

**HTTP Server**

    receive request
    look in cache for data to satisfy request
    if data in cache
      respond to request with cached data
    else
      if cache contains key=request_signature
        respond to request with no data.client should retry request
      else
        set cache key=request_signature value=NULL
        enqueue a job to fetch data
        respond to request with nothing. client should retry request
      end
    end

**Worker**

    look in queue for work
    if work
      process work
      save process result
      decode request_signature
      set cache key=request_signature, value=process_result
    end

Links
-----

-   [rack-worker](https://github.com/csquared/rack-worker)
-   [RailsConf 2010 SlideDeck](https://s3.amazonaws.com/ryandotsmith/deck.pdf)


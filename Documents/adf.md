# Software Delivery via Amazon Docker Work Flow

August 23rd 2013

## Introduction

Operations can be used vaguely in the web development space. There are
many aspects of software operations that simply can not be conveyed in a
small article or a hallway conversation. However, in this article we
will explore a sliver of operations known as software delivery.
Specifically, how to automatically create servers to receive software in
a generalized way.

For the purpose of this article, we will only consider application
servers. Servers which deal with state, namely: databases, messaging
systems, and caching layers are out of scope for this article. Using
first principles, let us explore what fundamentals we need in order to
deliver applications from developers to customers.

## The Problem

To deliver software, you need low latency access to a fleet of stand-by
servers which are well connected to the internet. These servers must
contain the app's code & configuration. The app's code need not be
homogeneous with other app servers in the network although they should
be compatible. If you can't run two unique versions of the code at the
same time, then you will need to take downtime to deploy your app. Given
user expectations and today's technology, there is no reason to assume
downtime for an app deployment. (However, if your customers are 100% ok
with downtime, always choose a simpler and less prone to error method of
software deployment)

The following scenario describes how an app should be deployed:

1. Make a release identified by a tag or the digest of the latest commit.
2. Use tool to update a config database with config changes.
3. Deliver release to build server. Possibly via Github.
4. Deliver release to a single server. Evaluates whether release can survive in production.
5. Replace servers running the previous release with servers running the latest release.
6. Cleanup config solely relating to the previous release.

Notes:

* This whole process should take 5 minutes.
* Arguably, some of these steps can be combined and automated, but these are the guiding objectives we wish to carry out.


## Requirements for the Solution

Unless you are in the business of software delivery, you want to own of
as little software delivery software as possible. Every system or module
you write to support your software delivery mechanisms will be a
liability against your organization for the entirety of its existence.
The age-old wisdom of **no code is better than no code** applies
perfectly to this situation. Therefore, we are in search of a solution
that requires the least amount of code while maximizing the efficiency
of our software delivery cycle. Any solution you evaluate for your
software delivery problem should be viewed from that perspective. So
I'll say it again, *Use solutions that require the least amount of code
while maximizing the efficiency of the software delivery cycle.*

## Amazon

Chances are good that you are relying on AWS for your servers. Thus,
Amazon's infrastructure is already a dependency for your environment. It
is now advantageous to use AWS to its fullest potential. Using
technologies like AMI, ELB, ASG, and DynamoDB are keys to success.

## Docker

Since we are hell-bent on maximizing effectiveness of software delivery
while minimizing supporting code, we must choose our next dependency
carefully. Docker is a good choice because it offers a beautiful
abstraction on top of a Linux AMI that will adapt the AMI to any type of
application we deploy.

## Amazon Docker Flow

We will build a base AMI that will serve as the foundation for any type
of application we wish to deploy. The AMI will be constructed once and
used for many deployments. The only reason to update the AMI is for OS
updates and security patches. \
 The AMI contains:

-   Ubuntu Linux 13.04
-   Docker
-   ADF-CONFIG(1) - Used to get app config from DynamoDB

Here is the overview of the components used for our server:

![](http://f.cl.ly/items/42151d3Y3c2X1K2g0X3G/adf-arch.png)

## ADF Setup Script

Here is the code required for the workflow. We have a Dockerfile, an upstart file, and a setup script that is run by upstart when the instance is booted.

<script src="https://gist.github.com/ryandotsmith/6287748.js"></script>

Demo: 0 to 'Hello World' in 108 Seconds.
----------------------------------------

<iframe width="640" height="480" src="//www.youtube.com/embed/_z_yX4pX2-M" frameborder="0" allowfullscreen=""></iframe>

## Details on the Flow

We are using Docker instead of building AMIs for each deploy because the
building of a new AMI can be quite slow. Also, the ability to include
the Dockerfile inside the projects provides a simple, clean way to
declare how the application should be built.

Each new deploy goes onto a fresh EC2 instance. We never deploy to the
same instance twice. If you want to refresh config variables, you can
restart the instance and the new config will be picked up.

In the demo, we showed Github's Releases features as the way to serve
code to our instances. This could very well be a build service like
docker.io. The thing to note here is that we set a RELEASE\_URL on our
EC2 instance to indicate to our setup script the location of the code
and Dockerfile.

We rely on the AWS console to manage our deploys. There are many
libraries for popular languages that will help build scripts for common
deployment tasks. It should be quite simple to build a campfire of
hipchat bot to automate the deployment task.

## Conclusion

Amazon and Docker make it easy to build a software delivery pipeline.
The presented flow and setup scripts demonstrate that you can build
delivery infrastructure with very few dependencies.
[Ephemeralization](http://adam.heroku.com/past/2011/4/7/ephemeralization/)
is a wonderful thing and we should always be striving to figure out how
to eliminate unnecessary components in our systems.

It should also be known that Docker is still in a pre-production state.
There are explicit warnings in the documentation that advise against
production use. While it is true that Docker is built on LXC and LXC is
a well known and trusted technology, Docker still has the capability to
terminate running LXC containers. Thus, it is possible for Docker to
take down a production site. Furthermore, there are plans to remove AUFS
in favor of BTRFS before the 1.0 release. I am advising all of my
customers to hold off on Docker for production use until the project
settles on a 1.0 release.

## Links

* [Node.js Demo App](https://github.com/ryandotsmith/adf-node-example)
* [ADF Scripts](https://gist.github.com/ryandotsmith/6287748)
* [ADF-CONFIG](https://github.com/ryandotsmith/adf-config)

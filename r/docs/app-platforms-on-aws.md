# Application Platforms on AWS

November 10th 2013

Choosing an infrastructure to run web servers can be difficult. It is not clear when IAAS (e.g. AWS) is appropriate or if PAAS (e.g. Heroku) will suffice. In the case of a startup company, it is probably best to start building on a PAAS since it provides the developer a path of least resistance. However, it is clear that there are real limitations of running a scalable business on a PAAS.

The goal of this article is to provide an architecture for building an application platform using AWS. One of the great features of PAAS is the 'S'. Servicing a platform is undoubtedly the hardest part. However, it is possible to architect AWS resources in such a way that the ongoing cost of operations is minimal. The easiest way to start using AWS is to provision an instance and issue custom commands on the instance. This is the wrong approach. Creating non-scripted infrastructure is the road to operations hell. The better way of using AWS is to design a platform for minimal operational burden, simple developer experience, and fault tolerance. The architecture described in this article has the following properties:

* AWS resource setup with 1 command in less than 5 minutes
* 1 click application deploys with versioned config and rollback capabilities
* Log aggregation and search
* System & Business metrics integration
* AWS best practices for high performance and availability

Experience shows that running a multi-tenant platform is a lot of work.  Instead, if you can focus on building a platform for a single organization which utilizes a small subset of technology, the platform can be constructed simply and with elegance. This approach also lends itself to the maximum possible performance characteristics. By definition, platforms with custom runtime components on top of AWS will be slower than platforms with 0 custom runtime components. The approach outlined in this article introduces no custom runtime components. Thus its performance characteristics are identical with the characteristics of AWS.

## The Problem

Organizations are faced with a tough optimization problem when building software. The organization must choose between fast, available, friendly infrastructure and cost. Any innovations outside of the organization's core competencies must be evaluated with extreme caution. In many cases, the use of AWS is an excellent way to outsource infrastructure while still delivering a robust infrastructure for developers to utilize.

### What Makes a Robust Infrastructure

From a product standpoint, the infrastructure must support fast iteration. Practically speaking this means that deploying new versions of the application should happen in less than 5 minutes. Deploys should be started by a web or chat interface and should be instigated by the developer. Deploys should be incrementally rolled out to prevent mass product defects. Rolling back a deploy should be as easy as the original deploy.  A robust infrastructure must also exhibit excellent operational qualities. The infrastructure must be debuggable. Easy access to business metrics, disk i/o, and network interfaces is required. The infrastructure must be reproducible. Armed with a single script and a set of IAM credentials, an operator should be able to build an environment in less than 5 minutes. A robust infrastructure can be scaled horizontally and vertically. Instances should be fully utilized and not over utilized. This requires the operator to have access to instances with 1GB of memory and instances with 64GB of memory.  Similarly the infrastructure should support horizontal scaling on demand. Moreover, the instances in the infrastructure should be abstracted into a cluster or group. If an instance under performing and consequently terminated, the group should replace the instance automatically. Instances should also be spread across availability zones to increase fault tolerance.

## A Solution

We have the technology to create our platform by simply scripting AWS primitives. The challenge we face is which primitives we choose and how we combine them. The solution described here is configured as such:

* S3 - Stores versioned release data (app & configuration)
* AMI - A single, curated base image that supports all applications
* ASG - Instances are grouped and scaled programatically by AWS
* ELB - Routes traffic to instances in the ASG
* Papertrail - Aggregates all log data
* Librato - Surfaces system and business metrics

Controlling these resources with a simple, opinionated API means that we can create our platform with a single command in less than 5 minutes.  This enables staging and test environments to be freely created and in complete parity to production.

It is assumed that infrastructure with the least amount of moving parts involved during production reconfiguration is ideal. However, there are engineering constraints that make the task of minimizing deployment parts. The primary concern in this trade-off is developer experience.  One deployment strategy is to treat the AMI as the deployment atom.  However, building an AMI, creating an ASG for the AMI, and cycling ASGs in and out of an ELB is a time consuming task. Thus the solution outlined in this article takes a slightly different strategy. The next sections will outline exactly how we diverge from the atomic AMI style of deployment.

### The Base Image

The base image includes all software necessary to receive a deploy and to run the deployed application. The base image only changes when OS level changes need to be propagated (e.g. security patches) The base image's software inventory can be grouped into 2 categories: infrastructure maintenance & application dependencies.

Infrastructure maintenance involves loading application configuration, downloading application source code, configuring upstart scripts, redirecting application logging, and cycling application processes.  Application dependencies consist of language runtime and dependency management tooling. (e.g. node.js and npm)

The base image need only provide a deployment API:

    $ /home/deploy/bin/deploy s3://releases/my-app/production/r1

### Release

A release is the combination of a bash file containing exported environment variables and a tar.gz file containing the application code.  Releases are stored in an S3 bucket with the following directory structure:

    .
    - amzn-releases.koa.la
        |-- app1
        |   |-- production
        |   |   |-- current -> r2/
        |   |   |-- env
        |   |   |-- r1
        |   |   |   |-- app.js
        |   |   |   |-- env
        |   |   |-- r2
        |   |       |-- app.js
        |   |       |-- env
        |   |       |-- lib.js
        |   |-- staging
        |-- app2

Release version numbers should correspond to tags in the Git repository.  Each app has a default env file which is the canonical source of configuration variables. When a release is made, the source is pushed to S3, the canonical env file is copied into the newly created source directory, and the version directory is linked to current. We use a current directory as a pointer so that newly created instances can programmatically grab the current release.

Source code will be pushed to S3 using git's archive command and then put to S3. This architecture lends itself to a variety of integrations such as Github post-receive hooks.

### Deployment

There are two ways to approach application deployment. The first perspective is from a single instance. Once we establish what it means for a single instance to receive a deploy we can consider a deployment to a group of instances.

Instance deployment is the process of downloading & installing a release. The deploy program, which is provided by the base image, receives a release url as an argument then downloads the release. The release is installed and the app's processes are restarted. More concretely, the deploy program executes the following steps:

1.  Download the release from the url supplied in the arguments passed to deploy
2.  Unpack the application
3.  Export the config variables
4.  Build upstart configuration based on Procfile in application source
5.  Restart process via upstart

Deploying to a group of instances is now just a matter of running the instance's deploy command on all instances in a group. The app name and environment are used to locate a group. Using the AWS SDK and an SSH library, a simple script can locate all instances in a group and issue the deploy command. The output of the deploy command can be echoed back to the operator for increased visibility. The deploy command provided by the base image ensures that failed deploys do not impact the running service.

### Operational Concerns

In addition to having a rich set of deployment commands, we must also satisfy operational requirements. The infrastructure must be capable of surfacing business and system metrics. The infrastructure must also enable system-level debuggability.

### Visibility

There are 3 key requirements to achieving infrastructure visibility:

* Log aggregation and search
* Business and system metrics
* Local instance performance tooling

Log aggregation is a great candidate for outsourcing. There are few cases in which you need to be innovative in your log operations.  Papertrail is the preferred log aggregation service and thus is recommended for this infrastructure. Thus the infrastructure needs to support delivering the log data to Papertrail. The easiest way to make the integration is via Rsyslog --which is installed on Ubuntu by default.  When the application is started via upstart, the command to start the application creates a pipe from stdout and stderr to logger. We don't use files to buffer logs since it requires a log rotation mechanism.  While the file based log buffer increases durability of the logs, we favor operational simplicity over durability. If data needs to be durable, we store the data in a database using database grade protocols.

In addition to having a log delivery pipeline, a robust infrastructure also requires business and system metrics. System metrics will originate from collectd and cloudwatch. Business metrics will be collected and aggregated within the application. Both sets of metrics are aggregated and stored in Librato. All metrics will be sourced by instance hostname which will provide the ability to debug group anomalies.

### Control Plane

All of the commands and workflows described in this article are made available by a stateless web application. Furthermore, the web application exposes its routines over HTTP to support ChatOps. There exist 3 workflows presented in the UI/COI:

-   Deploy application to environment
    -   Rollback
-   View application formation
    -   Shows ELBs and Instances behind ELB
    -   SSH link to instances
    -   Terminate instances or environments

The control plane is stateless and thus does not require a database. As such the control plane can be hosted in multiple regions connected by a cross-region DNS setup provided by Route53. Default data for launch configuration can be stored in an S3 bucket which can be operated on using AWS' Javascript Browser SDK. AWS API requests will use a control plane IAM credential. Access to the control plane will use HTTP Basic Authentication.

### Cycling Instances

Experience has shown that EC2 QoS is variable. An EC2 instance is not hardware, it is xen base virtual machine running on commodity hardware.  Over time an EC2 instance will degrade in quality. The operator will face the decision to root cause the degradation or throw away the instance and try again. Throwing away the instance is the best 1st step.  If the degradation persists across instances, then a root cause is appropriate. With this in mind, it is desirable to have an infrastructure that embraces the cycling of instances.

Since instances are created and managed by ASGs, and the ASGs are supplied with launch configuration that will execute a deploy on startup, the easiest way to cycle instances is to kill instances in the group and let the ASG manager replace them.

### Implementation
--------------

* Building the base image: [amzn-base](https://github.com/ryandotsmith/amzn-base)
* CLI Control Plane: [amzn-ship](https://github.com/ryandotsmith/amzn-ship)

### Implementation Demo

<iframe width="640" height="360" src="//www.youtube.com/embed/L7DOzsmeU2g" frameborder="0" allowfullscreen=""></iframe>

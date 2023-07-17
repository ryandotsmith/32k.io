# production checklist

*A backend perspective*

* Running 2 app servers in distinct availability zones
* Slightly over-provisioned database with hot standby
* Serve all static assets via CDN -- the app server should serve an asset exactly once
* All endpoints have been audited for auth{n,z}
    * each endpoint requires authentication
    * request data can't escalate privileges. (e.g. /users/123, /users124)
* Do a quick check against the [OWASP Top Ten][owasp]
* Enforce TLS
* Add request tracing and hook it up to Datadog
    * Global latency and endpoint specific latency views
    * Add SQL statements to database spans so that you can identify slow queries
* Rate limiting requests by ip and [by concurrency][rlc]
    * Just pick a high number so that some silly bot/script-kiddy doesn't take you down on day one
    * Don't use redis or whatever, just [do it in memory][rlim]
* Request logs with [structured data][logs]
* Exception/panic handlers via Sentry
* Runbooks describing how to
    * deploy & un-deploy
    * running migrations
* Staging environment
    * Have a client ready to smoke test basic user functionality
* Submit all billing data to your cloud provider and have all reasonable limits bumped
    * Don't have them shut off your server because you forgot to add a credit card
    * Sometimes providers have low limits on how many servers you can run and often you have to ask for an increase via support which can be slow
* Ensure that you have paid, low latency access to cloud provider tech support
* If you have time, do a load test
* Make sure more than one person has access to servers and deploy tools, etc...
* Have fun! You did it! You shipped your app to production!

February 20th, 2020

[rlc]: https://www.youtube.com/watch?v=m64SWl9bfvk
[owasp]: https://owasp.org/www-project-top-ten/
[rlim]: https://github.com/ryandotsmith/32k.io/blob/main/net/http/limit/limit.go
[logs]: https://brandur.org/logfmt

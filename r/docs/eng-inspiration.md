# Engineering Inspiration

This year marks my first decade in web services engineering. I'd be thrilled to share some of the things I've learned along the way. Let's start at the beginning.

What is a web service? Sometimes a web service is the backend for an iOS game, sometimes it's a transaction processor for a financial institution. Nevertheless, I'm talking about the type of programming where you write functions to handle an HTTP request, do business logic, read and write to a database, make an API request, and then send the client an HTTP response.

Over the years, I've developed a taste for good design in web services. And by no means have I perfected my style and I am constantly learning from people who have more or less experience than myself. And while I don't think I'll ever get it totally figured out I will say that I've found a few things that seem to work –specifically in the context of startups. Here we go:


## Simple Made Easy

[A talk by Rich Hickey](https://www.youtube.com/watch?v=rI8tNMsozo0)

Simplicity is objective. Count the number of things in a system. If the number of things is one then the system is simple otherwise it is complex. Simplicity and reliability go hand in hand.

Be careful with dependencies. Most of the time it is better to write a function to left pad a string instead of using a 3rd party dependency. The simplicity and control you get from writing code "by hand" usually dwarfs the initial time investment in writing the code.

## Crash-Only Software

[A paper by George Candea and Armando Fox](https://d.32k.io/crashonly.pdf)

Crash only is simple. Instead of worrying about how to start and stop a thing you only have to worry about starting it. This applies to http servers, requests, database operations, and other computer things.

For an HTTP API, you need to introduce idempotency keys to all your endpoints. When you get a new request, look up the ikey in your database and pick back up where you left off. Push your retry loop to the edge. Don't let people call your API without providing an SDK for them. In your SDK implement a retry loop and observe high availability!

I feel like there is a great paper out there on why it is scientifically better to push the retry driver to the edge, but I can't find it. Maybe I read it in a dream. But I feel like this is always a good idea. In fact, get rid of retry loops in your server. Just let the client retry!

## No Silver Bullet

[A paper by Fred Brooks](https://d.32k.io/no-silver-bullet.pdf)

> Einstein repeatedly argued that there must be simplified explanations of nature, because God is not capricious or arbitrary. No such faith comforts the software engineer. Much of the complexity he must master is arbitrary complexity, forced without rhyme or reason by the many human institutions and systems to which his interfaces must confirm. These differ from interface to interface, and from time to time, not because of necessity but only because they were designed by different people, rather than by God.

His note on OOP is interesting. I think we are observing that dogmatic OOP leads to poor design and consequently to a decrease in quality and maintainability. It seems like we are landing on a hybrid of OOP and functional programming. It's not uncommon to see a mix of modules with exported functions that accept primitive data types and a few classes with methods in a modern program. Which seems fine.

> The most radical possible solution for constructing software is not to construct it at all.

I love this. No code is better than no code!

> The hardest single part of building a software system is deciding precisely what to build.

This section is way ahead of its time. Time and time again I have seen that having a product manager who is responsible for the *what* of the system design has lead to cheaper, faster, and higher quality software builds.

I've heard another quote that rephrases this idea: It's a product manager's job to be responsible for building the right thing and it's an engineer's job to build the thing right. This division of thinking makes engineers and customers happy!

His ideas on growing a system instead of writing it, or building it, have really stuck with me. I've seen a lot of really great engineers take this path as well. In practice, this means starting a project with stubbed out function bodies and high concentration on function names and signature -- a focus on the data flow.

## Taking Down a Fence

[A memo by G.K. Chesterton](https://www.chesterton.org/taking-a-fence-down/)

It's easy to roll up into a new codebase with delete keys blazing. Let's face it, we love to rewrite software. I'm not saying that's bad. In fact, I think it is healthy to build an expectation of rewrites into a team's culture.

But Chesterton has reminded us all that before we delete or "refactor" something we ought to understand the thing and why the thing was built to begin with.

## Some thoughts on security after ten years of qmail 1.0

[A paper by Daniel J. Bernstein](https://d.32k.io/qmailsec-20071101.pdf)

The big thing I got from this paper is that we should build small, isolated programs for handling key material. At Chain we built HSM firmware and even took a stab at building our own transaction signing service. In both cases, we went to a lot of effort to ensure that the process that signed the transaction only had 1 or 2 APIs. We wanted the key managing process to be as simple as possible.

This is nice for several reasons:

* audit logs are focused. you don't have to sift to find security related logs
* auditing the code is easier because you don't have to to look hard for security related bits
* scaling a signing server vs. your public API is different. it's better to have fewer number of security sensitive servers.
* different servers means different access policies. it's nice to be able to limit the number of people who have access to key servers

## Notes on Programming in C (Pikestyle)

[A webpage by Rob Pike](http://doc.cat-v.org/bell_labs/pikestyle)

While this was intended for C programmers, there is a lot of wisdom for us mere mortals too. In particular, check out the section on complexity. His note on variable names is worth a read as well.

Here's another nugget:

> Data structures, not algorithms, are central to programming.

Which reminds me of another gem from Fred Brooks:

> Show me your flowchart and conceal your tables, and I shall continue to be mystified. Show me your tables, and I won't usually need your flowchart; it'll be obvious.

## Harvest, Yield, and Scalable Tolerant Systems

[A paper by Armando Fox and Eric A. Brewer](https://d.32k.io/harvest-yield-scalable-tolerant-systems.pdf)

This is a fun paper and it will certainly tickle all of your nerd bones. But, I've seen so many startups get sidetracked by ideas presented in this paper. I've found it prudent to care more about consistency than availability in the beginning. In practice, your AWS hosted Postgres instance won't go down very often. In fact, I've seen production instances go for nearly a year before seeing any downtime! I bet misconfigured deploys cost startups more downtime than any particular database technology.

<hr />

That's it! Well, that's not entirely true. The biggest influence on my career in software has been the amazing people that I've been privileged to work with. I've always tried to be a small fish in a big pond --and thank goodness for that! And thank you:  Adam Wiggins, Blake Mizerany, Daniel Farina, Jackson Owens, Keith Rarick, Mark Fine, Mark McGranaghan, Tess Rinearson, Tom Maher and every single engineer at: Heroku, Chain, and Pogo!

TODO(ryan): write a similar note for leadership

*March 30th, 2020*

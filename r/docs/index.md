<h1>Ryan R. Smith</h1>

Waddup 🤙 Im a son of God, I want to be like Jesus, I am a lovesick worshiper, I live in West Oakland, I ride harleys, I derp computers, and other things that are not listed on this webpage...

In 2019, my co-founders and I started a new project that we call Pogo. It's a secret for now. In 2013 I co-founded [Chain](https://web.archive.org/web/20180809015203/https://chain.com/) and served as CTO up until we sold it to [Stellar](https://stellar.org) in 2018. In 2010 I joined the Heroku engineering team to help build the world's greatest PAAS. Before joining Heroku I earned a B.S. in Mathematics from the University of Missouri Kansas City.

## My Journal

* [Waiting on God - Prov 28:9](/prov-28:9)
* [Infected with Goodness](/infected-with-goodness)
* [New Age Christianity](/new-age-christianity)
* [Dream Job](/dream-job)
* [Let's Give Up](/lets-give-up)
* [The Narrow Way - A Personal Story](/narrow-way)
* [The Gospel - Version 1](/gospel-v1)
* [Continuing Forgiveness](/forgiveness)
* [Rest for my soul](/rest-for-my-soul)
* [God](/God)
* [Friendships and Oak Trees](/friendships-and-oak-trees)
* [No Greater Miracle](/no-greater-miracle)
* [Remind Me of Your Name](/remind-me-of-your-name)

## Contact

* [r@32k.io](mailto:r@32k.io)
* [github.com/ryandotsmith](https://github.com/ryandotsmith)

## Talks

* [Predictable Failure](http://vimeo.com/75304752) - Monitorama 2013 [s](http://cl.ly/1o1o243O0z2A/Predictable%20Failure%20Monitorama.pdf)
* [Fewer Constraints More Concurrency](http://vimeo.com/68850147) - RubyKaigi 2013 [s](http://cl.ly/2812472J073R/Ruby%20Kaigi%202013%20-%20Fewer%20Constraints%20More%20Concurrency.pdf)

## Technical Articles

* [Production Checklist](/production-checklist)
* [Building an App Platform on AWS](/app-platforms-on-aws)
* [L2met Introduction](/l2met-introduction)

## Notes From My Career

* [Engineering Inspiration](/eng-inspiration)
* [Engineering Values](/eng-vals)
* [Personal Accountability at Chain](/personal-accountability-at-chain)
* [Farewell, Heroku](/farewell-heroku)

<hr />

[public key](/pk), [random](/random), [travel](/travel), [photos](/photos), [chopper links](/chopper-links), [gifs](/gifs)

<hr />
<img src="https://d.32k.io/earth.gif" style="margin: 0 auto; display: block; max-width: 50px;" alt="world wide web">

<p style="text-align: center">
<a href="https://hotlinewebring.club/ryandotsmith/previous">←</a>
Hotline Webring
<a href="https://hotlinewebring.club/ryandotsmith/next">→</a>
</p>

<script>
window.addEventListener('load', function() {

    function randcolor() {
        const colors = ['#bc5a45', '#618685', '#36486b', '#f18973'];
        var i = Math.floor(Math.random() * Math.floor(colors.length));
        return colors[i];
    }

    function spanit(tn) {
        var col = document.getElementsByTagName(tn);
        for (var i = 0; i < col.length; i++) {
            var text = col[i].innerText;
            col[i].innerText = '';
            for (var j = 0; j < text.length; j++) {
                var s = document.createElement('span');
                s.className = 'colored';
                s.innerHTML = text[j];
                col[i].appendChild(s);
            }
        }
    }

    var x = 0;
    function rot(offset) {
        x++;
        if (x%offset== 0) {
            var col = document.getElementsByClassName('colored');
            for (var i = 0; i < col.length; i++) {
                col[i].style.color = randcolor();
            }
        }
    }

    spanit('h1');
    spanit('h2');
    window.addEventListener('mousemove', function() {rot(20)});
    window.addEventListener('scroll',    function() {rot(5) });
});
</script>

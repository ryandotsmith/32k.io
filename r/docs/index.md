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

<h1>Ryan R. Smith</h1>

Waddup 🤙 Im a son of God, I want to be like Jesus, I am a lovesick worshiper, I live in Frisco, I ride bikes, I derp computers, and other things that are not listed on this webpage...

In 2013 I co-founded [Chain](https://chain.com) and serve as CTO. In 2010 I joined the Heroku engineering team to help build the world's greatest PAAS. Before joining Heroku I earned a B.S. in Mathematics from the University of Missouri Kansas City.

## Contact

* [r@32k.io](mailto:r@32k.io)
* [github.com/ryandotsmith](https://github.com/ryandotsmith)

## Talks

* [Predictable Failure](http://vimeo.com/75304752) - Monitorama 2013 [s](http://cl.ly/1o1o243O0z2A/Predictable%20Failure%20Monitorama.pdf)
* [Fewer Constraints More Concurrency](http://vimeo.com/68850147) - RubyKaigi 2013 [s](http://cl.ly/2812472J073R/Ruby%20Kaigi%202013%20-%20Fewer%20Constraints%20More%20Concurrency.pdf)

## Spirituality

* [The Narrow Way - A Personal Story](/gospel-v1)
* [The Gospel - Version 1](/gospel-v1)
* [Continuing Forgiveness](/forgiveness)
* [Rest for my soul](/rest-for-my-soul)
* [God](/God)
* [Friendships and Oak Trees](/friendships-and-oak-trees)
* [No Greater Miracle](/no-greater-miracle)
* [Remind Me of Your Name](/remind-me-of-your-name)

## Technical Articles

* [Building an App Platform on AWS](/app-platforms-on-aws)
* [L2met Introduction](/l2met-introduction)

## Notes From My Career

* [Engineering Values](/eng-vals)
* [Personal Accountability at Chain](/personal-accountability-at-chain)
* [Farewell, Heroku](/farewell-heroku)

## Favorite Quotes

"A human being should be able to change a diaper, plan an invasion, butcher a hog, conn a ship, design a building, write a sonnet, balance accounts, build a wall, set a bone, comfort the dying, take orders, give orders, cooperate, act alone, solve equations, analyze a new problem, pitch manure, program a computer, cook a tasty meal, fight efficiently, die gallantly. Specialization is for insects."

*--Robert Heinlein, Time Enough for Love*

"If you want to go fast, go alone. If you want to go far, go together."

*--African Proverb*

"The business world –where the majority of American men live and die– requires a man to be efficient and punctual. Corporate policies and procedures are designed with one aim: to harness a man to the plow and make him produce. But the soul refuses to be harnessed; it knows nothing of Day Timers and deadlines and P&L statements. The soul longs for passion, for freedom, for life."

*--John Eldredge, Wild at Heart*

"It's the possibility of having a dream come true that makes life interesting,"

*--Paulo Coelho, The Alchemist*

"And don’t worry about losing. If it is right, it happens—The main thing is not to hurry. Nothing good gets away."

*--John Steinbeck, A Letter to His Son*

When Rabbi Zusha was on his deathbed, his students found him in uncontrollable tears. They tried to comfort him by telling him that he was almost as wise as Moses and as kind as Abraham, so he was sure to be judged positively in Heaven. He replied, "When I get to Heaven, I will not be asked Why weren't you like Moses, or Why weren't you like Abraham. They will ask, Why weren't you like Zusha?"

*--Rabbi Meshulam Zusha of Hanipol*

<hr />

[public key](/pk), [odds](/odds), [travel](/travel), [photos](/photos), [chopper links](/chopper-links), [gifs](/gifs)

<hr />
<img src="https://d.32k.io/earth.gif" style="margin: 0 auto; display: block; max-width: 50px;" alt="world wide web">

<p style="text-align: center">
<a href="https://hotlinewebring.club/ryandotsmith/previous">←</a>
Hotline Webring
<a href="https://hotlinewebring.club/ryandotsmith/next">→</a>
</p>

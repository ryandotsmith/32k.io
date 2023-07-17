<h1>Ryan R. Smith</h1>

Welcome to my web page. My interests include: mythology, theology, history of western civilization, trees, architecture, decentralization, other similar things, and all thing related to computer programming.


October 2022 I started a new company called [Index Supply, Co.](https://github.com/indexsupply) I'm going to build a new kind of ethereum node.

My work history in a nutshell:
- 2021 - 2022 - Built the Ethereum indexing system at [mint.fun](https://mint.fun).
- 2018 - 2019 - Started and stopped a cross border wallet called Pogo with co-founders from Chain.
- 2013 - 2018 - Co-founder / CTO at Chain. Acquired by Interstellar –the comercial arm of Stellar. Pivoted from Bitcoin API to private blockchain node to decentralized ledger.
- 2010 - 2013 - Heroku engineer. Joined as the 18th employee and grew with the company till we hit 150 people and millions of apps. Loved Heroku a lot!

Before that:
- 2006 - 2010 - Undergrad Math at the University of Missouri KC.
- 1986 - I’ve been programming computers for as long as I can remember.

Listed in [30u30](https://www.forbes.com/pictures/mdg45ejdik/ryan-smith-29/?sh=2aed631ac70b) and not convicted of crimes

### Books I've read recently

- [Anti-Intellectualism in American Life, The Paranoid Style in American Politics](https://www.goodreads.com/book/show/51005822-richard-hofstadter)
- [Khomeinism: Essays on the Islamic Republic](https://www.goodreads.com/book/show/238095.Khomeinism)
- [Six Easy Pieces: Essentials of Physics Explained by Its Most Brilliant Teacher](https://www.goodreads.com/book/show/10025702-six-easy-pieces)

## Spirituality

* [Work Life Balance](/work-life-balance)
* [Where is God?](/where-is-god)
* [Friendships and Oak Trees](/friendships-and-oak-trees)
* [Remind Me of Your Name](/remind-me-of-your-name)

## Finance

* [Money is a Reflection - Investing](/money-reflect-invest)

## Talks

* [Predictable Failure](http://vimeo.com/75304752) - Monitorama 2013 [s](http://cl.ly/1o1o243O0z2A/Predictable%20Failure%20Monitorama.pdf)
* [Fewer Constraints More Concurrency](http://vimeo.com/68850147) - RubyKaigi 2013 [s](http://cl.ly/2812472J073R/Ruby%20Kaigi%202013%20-%20Fewer%20Constraints%20More%20Concurrency.pdf)

## Nerd Stuff

* [KN6LLA](/kn6lla)
* [Production Checklist](/production-checklist)
* [Building an App Platform on AWS](/app-platforms-on-aws)
* [L2met Introduction](/l2met-introduction)

## Notes From My Career

* [Engineering Inspiration](/eng-inspiration)
* [Engineering Values](/eng-vals)
* [Personal Accountability at Chain](/personal-accountability-at-chain)
* [Farewell, Heroku](/farewell-heroku)

## Contact

* [r@32k.io](mailto:r@32k.io)
* [github.com/ryandotsmith](https://github.com/ryandotsmith)

<hr />

[public key](/pk), [random](/random), [travel](/travel), [photos](/photos), [chopper links](/chopper-links), [gifs](/gifs)

<hr />
<img src="https://d.32k.io/earth.gif" style="margin: 0 auto; display: block; max-width: 50px;" alt="world wide web">

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

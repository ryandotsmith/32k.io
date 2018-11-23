# chopper links

![img](https://d.32k.io/Harley-Davidson-Engines-7--3184-default-large.jpeg)

Here is a collection of resources for all things Harley Chopper

1. [Technical Notes](#technical-notes)
2. [Small Batch Manufacturing](#small-batch-manufacturing)
3. [Repop Manufacturing](#repop-manufacturing)
4. [Retailers](#retailers)
5. [Garages](#garages)
6. [Apparel](#apparel)

## Sick Vids

* [2 Days In Watsonville](https://vimeo.com/148416019)
* [Ramble On, Brother](https://vimeo.com/254214424)

## Technical Notes

* Shovelhead Points/Timing Adjustment
* Shovelhead CB Carb Adjustment
* [Sportster Chain Conversion](http://www.chopcult.com/forum/showthread.php?t=15790)

## Small Batch Manufacturing

* [American Prime](https://americanprimemfginc.com/) - clutch
* [After Hours Parts](http://afterhourschoppers.com) - tail lights
* [Baker](https://bakerdrivetrain.com/) - transmission
* **[BCM](https://prismmotorcycles.com/)** - air cleaners
* [Bone Orchard](https://www.theboneorchardcycles.com/) - shift konb, cleaners
* [B&C Cycles](https://bnccycles.com/) - hi mid controls
* [Conflict Machine](https://www.conflictmachine.org/) - fork caps
* [Counter Balance Cycles](https://counterbalancecycles.com/) - seats
* [Cro's](https://crocustoms.bigcartel.com/product/cro-bars) - cro bars
* [Deep Six Cycles](http://www.deepsixcycles.com) - custom fab
* [Denvers Choppers](https://denverschoppers.com/) - springers
* [Fab Kevin](http://www.fabkevin.com) - brakes
* [Fagerberg Machine](http://www.fagerbergmachine.com/) - linkert carbs
* [Frank's Forks](http://franksforks.com/) - fork tubes
* [Gasbox](https://www.thegasbox.com/) - period exhaust
* [Haifley Brothers](http://www.haifleybrothers.com/) - hard tail kits
* [King Mfg.](https://www.kingmfg.co) - wires
* [Kustom Tech](http://www.kustomtech.eu/en/) - wheels, brakes
* [Main Drive](http://www.maindrivecycle.com/) - frisco mid controls
* [Morris Magneto](http://shop.morrismagneto.com/) - magneto
* [Mullins Chain Drive](http://mullinschaindrive.bigcartel.com) - triple tree
* [Nash Motorcycle Co](https://www.nashmotorcycle.com) - evo stuff
* [No School Choppers](http://www.noschoolchoppers.com/) - tail lights
* [Pangea Speed](https://pangeaspeed.com) - bars
* [Phares Cycle Parts](https://pharescycleparts.com/) - springers
* **[Prism Supply Co](https://prismmotorcycles.com/)** - everything!
* [Psycho Resin](http://www.psychoresin.bigcartel.com/) - taillight lenses
* [Regatta Garage](http://regattagarage.com/shop/) - risers
* [Sprucest Fabrication](https://sprucestfabrication.bigcartel.com/) - coil mount
* [Twin City Cycle Parts](http://www.twincitycycleparts.com/) - bungs
* [Visionary Cycle Products](https://visionarycycleproducts.com/) - springer brakes
* [Wargasser Speed Shop](https://www.instagram.com/wargasserspeedshop) - Juice drum brake brackets
* [Zombie Performance](http://zombieperformance.com/) - handlebars
* [0 Given](http://www.0given.com/) - tank decals

## Repop Manufacturing

* [Colony Machine](http://www.colonymachine.com/ColonyCatalog2018.pdf)
* [Eastern Motorcycle Parts](http://www.easternmotorcycleparts.com/)
* [V-Twin Manufacturing](https://www.vtwinmfg.com/)

## Retailers

* [Bison Motor Sports](http://www.bisonmotorsports.com/)
* [Deadbeat Customs](https://www.deadbeatcustoms.com/)
* [Delux HD Restoration](* https://deluxehdrestorations.com/)
* [Dennis Kirk](* https://www.denniskirk.com)
* [Lick's Cycles](https://www.lickscycles.com)
* [Road Side Repair Shop](https://www.roadsiderepairshop.com/)
* [TC Bros](http://www.tcbroschoppers.com) - combo of retailer / manufacturing
* [Throttle Addiction](http://www.throttleaddiction.com) - combo of retailer / manufacturing

## Garages

* [Lucky Wheels](https://www.luckywheelsgarage.com/) - Los Angeles
* [Steady Rolling](http://www.steadyrollingmotorcycles.com/) - Oakland

## Helmets

* [Custom Destruction](https://www.helmetrestoration.com/)
* [Head Candy](http://headcandyfactory.com/index.html)
* [Joe King Speed Shop](https://joekingspeedshop.bigcartel.com/)

## Apparel

* [Abel Brown](http://www.psychoresin.bigcartel.com/) - tent
* [All Good Days](https://allgooddays.bigcartel.com/products) - shirts
* [Chopper Supply Co.](http://www.choppersupplyco.com/)
* [Cycle Zombies](http://cyclezombies.com)
* [Love Cycles](http://lovecycles.bigcartel.com/)
* [Sweatshop Industries](https://sweatshopind.myshopify.com/) - sissy bar bag

<hr />
send me your suggestions

<script>
window.addEventListener("load", function () {
    function post(f) {
        var req = new XMLHttpRequest();
        var body = new FormData(f);
        req.addEventListener("load", function(event) {
            alert(event.target.responseText);
            document.getElementsByName("suggestion")[0].value = "";
        });
        req.addEventListener("error", function(event) {
            alert('Oops! Something went wrong.');
        });
        req.open("POST", f.action);
        req.send(body);
    }
    var form = document.getElementById("suggest");
    form.addEventListener("submit", function (event) {
        event.preventDefault();
        post(form);
    });
});
</script>

<textarea name="suggestion" form="suggest" style="width: 96%; height: 75px;"></textarea>
<form id="suggest" action="https://server.32k.io/c/suggest" method="post">
  <input type="submit">
</form>

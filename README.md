# Using the Theorem project idea to explore project structure in golang.
Sometime around the end of 2018  I was approached by a company out of Santa Monica called Theorem about a fully remote senior dev consulting position. Back in 2018, that was still kind of unusual, so I pursued it with a passion.  
After a lengthy and thorough interview schedule which included code samples, gists for small technical problems, and calls with some very interesting people all over the world, up to and including the CTO in Santa Monica, they still wanted more.
At this point I've spent maybe 1.5 months talking to these guys, and I'm still very interested.  My last project had finished up, and I was just itching to end the downtime.  Anyway, they came to me with a final task:  spend two days coding a solution for some given scenario.  
They offered to pay me small fee for my two days of effort and said at the end they'd be able to make a final decision.  The task they presented was for a Smart Air Conditioning unit which would update the mother-ship with periodic telemetry data.  They needed a device API and an admin web app.  There were a number of standard requirements like search capability, some basic alerting reqs, etc.  
Anyway you can read more about it in the [project spec](https://github.com/brettman/SmartAcGo/blob/master/docs/%5BBackend%5D%20Smart%20AC.pdf) they gave me.  

At the time, I wasn't comfortable enough with Golang to build this in two days, so I reached for C# and gave it my best shot.  

Because it's a nice, self-contained project with clear requirements, I later decided to pick it up again and build it out in Golang, as I had originally wanted to do.  This repo is the result of those ongoing efforts.
After a long hiatus, I've decided to pick this back up and see how far I can get it in my "free" time.  



## project structure
[Go API project tutorial](https://tutorialedge.net/courses/go-rest-api-course/)

[Medium: go-project-layout](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2)

[Golang Standards: project-layout](https://github.com/golang-standards/project-layout)

[Medium: standard-package-layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)

[Medium: structuring applications in go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)

[Applying clean architecture to golang](https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)


## database
[Installing Mongodb on a mac](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/)

[Short and handy mgo gist](https://gist.github.com/345161974/4f2048f90584a64891cf07997bfd9e23)

I'm starting to feel like maybe mongoDb is a waste of time...  In any case, I think I'd rather dig deeper into PostGRE implementations in golang.
Of course this will open a whole 'nuther can o' worms. Which of the million-and-one drivers out there to choose... Use an ORM or not... 

[PGX: A pure go driver for postgres](https://github.com/jackc/pgx)

[All the documentation for PostGRE](https://www.postgresql.org/docs/current/index.html)

## logging
[Zerolog](https://github.com/rs/zerolog)
# Places distance

Find distance between 2 addresses located in France. This tool relies on [government open api](https://geo.api.gouv.fr/adresse) to provide a distance computation between 2 textual addresses

Distance is a spherical distance computation. It gives the same result than the distance computation on gmap (right click for tests) in my own few tests.

I found the distance computation function here https://www.geodatasource.com/developers/go under LGPL v3 as of 18-04-2020, so is this project

## Usage

All commands are located in [Makefile](Makefile)

For quick test you can do

`make docker-run`

then

`make test`

This should display the computed distance between the two adresses given in parameter in the curl call

Result looks like

```javascript
    //HTTP 200
    {"result": 23.19747198781676} // Result is in kilometers
```

## Licence

[LGPL v3](LICENCE.md)
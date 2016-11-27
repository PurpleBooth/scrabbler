# Scrabbler

Library that returns what points you can get back from a scrabble game.

## Getting Started

These instructions will get you a copy of the project up and running on your
local machine for development. I recommend running this on a docker based
environment, so your prod is the same as your dev.

### Prerequisities

You'll need [docker].

### Trying it out

Then run the application with
```
docker-compose up
```

You can then see the application running by

```
curl http://$(boot2docker ip):8080/cats
```

### Environment Variables

* `NEW_RELIC_APP_NAME` - Optional New Relic Application Name
* `NEW_RELIC_LICENSE_KEY` - Optional New Relic License Key

## Contributing
Feel free to submit pull requests and issues. If it's a particularly large PR,
you may wish to discuss it in an issue first.

Please note that this project is released with a
[Contributor Code of Conduct](CONDUCT.md). By participating in this project you
agree to abide by its terms.

## Versioning

We use [SemVer] for the version tags available See the tags
on this repository.

## Authors

See the list of [contributors] who participated in this project.

[contributors]: https://github.com/PurpleBooth/scrabbler/contributors
[SemVer]: http://semver.org/
[docker]: https://www.docker.com/products/overview

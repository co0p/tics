TICS - Thumbnail Image Caching Service
=====================================

A proof of concept implementation of a clean-architecture (see [uncle bob's clean architecture description](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)) with 
an image caching service as the main use-case.

 * The *domain* is the thumbnail. 
 * The *usecase* is the image resizing and storing part. 
 * The *interface* of the system is a webservice and storage abstraction.
 * The *infrastructure* implements the concrete resizer, storage etc systems.

how to run
----------

 * Go to ```$ cd $GO/src/github.com/co0p/tics/cmd/tics-webserver``` and then install the program with ```$ go install```. 
 * Run it with: ```$ /pth/to/go/bin/tics-websever```.

how to use
----------

The tics service is using a base64 encoded url including width and height resizing information:
 
 * url : https://www.wikipedia.org/portal/wikipedia.org/assets/img/Wikipedia-logo-v2.png
 * width: 20
 * height: 20

should be converted into the following pattern:```<url>?w=<width>&h=<height>```. Finally run all this through a base64 encoder and you are good to go:

  * https://www.wikipedia.org/portal/wikipedia.org/assets/img/Wikipedia-logo-v2.png?w=20&h=20
  * aHR0cHM6Ly93d3cud2lraXBlZGlhLm9yZy9wb3J0YWwvd2lraXBlZGlhLm9yZy9hc3NldHMvaW1nL1dpa2lwZWRpYS1sb2dvLXYyLnBuZz93PTIwJmg9MjA=

Now you have a nice base64 encoded image path and some resizing information. Call the service with
 
 * ```localhost:8080/?i=<base64encoded...>```
 * localhost:8080/?i=aHR0cHM6Ly93d3cud2lraXBlZGlhLm9yZy9wb3J0YWwvd2lraXBlZGlhLm9yZy9hc3NldHMvaW1nL1dpa2lwZWRpYS1sb2dvLXYyLnBuZz93PTIwJmg9MjA=

and you get a nice small wikipedia image.


run tests
---------

```cd /path/to/tics``` and then ```go test ./...```.

License
-------

This project is licensed under MIT license. Use at your own risk and will!
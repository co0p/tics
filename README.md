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

Install the program with ```$ go install``` and then run it with: ```$ tics```.

run tests
---------

```cd /path/to/tics``` and then ```go test```.

License
-------

This project is licensed under MIT license. Use at your own risk and will!
# Go APNs [![Build Status](https://travis-ci.org/mattprice/Go-APNs.png)](https://travis-ci.org/mattprice/Go-APNs)

**Work In Progress:** A Golang package for easily connecting and sending notifications to the Apple Push Notification Service. We abstract all the quirks and difficulties of the APNs so that you can focus on your application.

## Features
* **Based on a [queue system](http://redth.info/the-problem-with-apples-push-notification-ser/ "The Problem with Apple's Push Notification Service") instead of timeouts.** All your notifications will be delivered without sacrificing speed.
* **Certificates can be passed as a string or file path.** This allows you to store your certificates anywhere you want—in a database, environmental variable, or on disk.
* **Simplifies creating the notification payload.** You can start focusing on what you need your notifications to do instead of Apple's API.
* **Creates a persistent connection to Apple's server.** This is recommended by Apple, and it allows you to send notifications faster by avoiding multiple, costly SSL handshakes.

## License (MIT)
Copyright (c) 2013 Matthew Price, http://mattprice.me/

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
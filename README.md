# Go APNs [![Build Status](https://travis-ci.org/mattprice/Go-APNs.png)](https://travis-ci.org/mattprice/Go-APNs)

**Work In Progress:** A Golang package for easily connecting and sending notifications to the Apple Push Notification Service. We abstract all the quirks and difficulties of the APNs so that you can focus on your application.

## Features
* [Based on queues](http://redth.info/the-problem-with-apples-push-notification-ser/ "The Problem with Apple's Push Notification Service") instead of timeouts. Ensure that all your messages get delivered without sacrificing any speed or worrying about socket disconnections.
* [Certificates](http://developer.apple.com/library/mac/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/ProvisioningDevelopment.html#//apple_ref/doc/uid/TP40008194-CH104-SW6) can be passed as either a file or a string, allowing you to store the certificates in a database or environmental variable.
* Completely handles the creation of the [JSON payload](http://developer.apple.com/library/mac/#documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/ApplePushService.html). Stop worrying about how and when to use the alert dictionary format required for complex notifications.
* Allows you to create a persistent connection to Apple's server, [as recommended](http://developer.apple.com/library/ios/#documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/CommunicatingWIthAPS.html). Perfect for applications that need to continuously send notifications throughtout the day, or for applications that only need to send notifications through in small batches.

## License (MIT)
Copyright (c) 2013 Matthew Price, http://mattprice.me/

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
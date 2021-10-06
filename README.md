# SFV-API


[![AGPL License](https://img.shields.io/badge/license-MIT-brown.svg?style=flat-square)](https://opensource.org/licenses/MIT) 

## What is it?
`SFV-API` is an open source Street Fighter V API that exposes the profile page under CFN as a JSON API. It is currently work-in-progress, under heavy development and is not yet recommended for production use.

## Installation
As `SFV-API` is a work-in-progress there is no set way of installaing and running, while we're developing the first version... the day-life developing is made through `make run` (locally).

At cmd/api/api.go there is a reference to `token`. This value must be extract from the CFN page after the Sign In process. Inspect your browser and find the cookie labed as **scirid**. Once this cookie is provided, each following request will store all response headers in a file (cookies.json). Sometimes after some requests those values change regarding you aren't using a bot (¬¬).

I experienced that after some days without a single request the cookie expires. But every time you do a new request and the token is refreshed, the expiration date is increased. No problem whether the API receives some requests along the time.
## Licensing
All code is licensed under the [MIT](https://github.com/taunti/sfv-api/blob/master/LICENSE).

All CAPCOM references: SFV franchise, CFN, and Characters are registered trademarks or trademarks of CAPCOM CO., LTD.

### MIT License

    Copyright 2021 (taunti)

    Permission is hereby granted, free of charge, to any person obtaining a copy 
    of this software and associated documentation files (the "Software"), to deal 
    in the Software without restriction, including without limitation the rights 
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies 
    of the Software, and to permit persons to whom the Software is furnished to do 
    so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all 
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, 
    INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
    PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT 
    HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION 
    OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE 
    OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.



## Contributing and Feedback
Currently, you can contribute to the SFV-API project by:
* Submitting a detailed [issue](https://github.com/taunti/sfv-api/issues/new).
* [Forking the project](https://github.com/taunti/sfv-api/fork), and sending a pull request back to for review.

There is a Discord Server channel `#sfv-api` on (), for talking directly with testers and developers.

### Core Maintainers

* "taunti" <https://github.com/taunti>

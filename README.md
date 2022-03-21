# Exercise: Write a simple fizz-buzz REST server.

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing
all multiples of 3 by `fizz`, all multiples of 5 by
`buzz`, and all multiples of 15 by `fizzbuzz`. The output would look like this:
> 1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...

Your goal is to implement a web server that will expose a REST API endpoint that:

1. Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and
   str2.
2. Returns a list of strings with numbers from 1 to limit, where:
    - all multiples of int1 are replaced by str1
    - all multiples of int2 are replaced by str2,
    - all multiples of int1 and int2 are replaced by str1str2.

3. The server needs to be:
    - Ready for production
    - Easy to maintain by other developers

4. Bonus: Add a statistics endpoint allowing users to know what the most frequent request
   has been. This endpoint should:
    - Accept no parameter
    - Return the parameters corresponding to the most used request, as well as the number
      of hits for this request

---

# How to use


To run the application:
```sh
make build
make run-localy //(to run the server on local machine)
or 
make build-image (to build a docker image)
make run-container (to run the docker image)
```

To launch the test:

```sh
  make test
```

##### The Route:

 ```
 
 Get  /api/compute/fizzbuzz (query params are "fstM" for the first multiple,"secM" for the second multiple, "limit" for the limit,"label1" for the first label,"label2" for the second label)
 Get  /api/metrics/besthits (no query params)
 Get  /health (or just "/")  (no query params)  
 ```



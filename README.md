# DoublePrime_GO
Exposes a rest end point which returns true if a number is a double prime and false if it isn't  

### Pre requisites:
1. Install golang 1.13.x
2. Install dependencies
In the project root, run the following command to install all the dependencies
`
$ go install -v
`

### Steps to run:
1. Run the following command to start the rest api
`
$ go run index.go
`
2. To test, use the following curl command, where ```<your-number>``` is the sample number you want to test for
`
$ curl -v http://localhost:8010/isDoublePrime/<your-number>
`
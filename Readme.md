### Welcome to json_checker, a simple app to check if there are specified Resources in Resource field in your "AWS::IAM::Role Policy"

json file.

To use this tool, make sure you have installed Go language. In case you haven't, follow steps listed here:

https://go.dev/doc/install

To check your json file, run:

```bash
go build ./src/main
```

and after that:

```bash
./main <filepath to your file>
```

To run tests, do:

```bash
go test ./src/policy
```

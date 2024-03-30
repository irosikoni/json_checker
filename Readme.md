<h3>Welcome to json_checker, a simple app to check if there are specified Resources in Resource field in your "AWS::IAM::Role Policy"
json file.</h3></br>

To use this tool, make sure you have installed Go language. In case you haven't, follow steps listed here:</br>
https://go.dev/doc/install </br>

To check your json file, run: </br>
```
go build ./src/iamrp ./src/main
```
and after that: </br>
```
./main <filepath to your file>
```
</br>

To run tests, do: </br>

```
go test ./src/iamrp
```
</br>
In case you would like to use the validating function in your project, you can import it into your module, add </br>

```
import (
github.com/irosikoni/json_checker/src/iamrp
)
```
to your package, and then use a function like that: </br>

```
iamrp.IAMRolePolicyValidator(<your filepath>)
```

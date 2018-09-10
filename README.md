# go-mvn-artifacts-prop

### Go Utility to parse version from pom.xml

Example

````
$ curl -L https://github.com/muhilan/go-mvn-artifacts-prop/releases/download/0.1.0/go-mvn-artifacts-prop -o version

$ chmod +x version

````

In a folder where pom.xml is present run the following

````
$ ./version
$ 0.1.0-SNAPSHOT
````

````
$ helm package --dependency-update --version "$(./version)" ./myhelmfolder
````

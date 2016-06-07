# gapp
Dummy web application written in Go

# Content

1. [Prerequisites](#prerequisites)
2. [Getting the application](#getting-the-application)
3. [Building](#building)
4. [Testing](#testing)
5. [Running](#running)

## Prerequisites
<table>
  <tbody>
    <tr>
      <td>GNU Make</td>
      <td>(>= 4.1)</td>
    </tr>
    <tr>
      <td>git</td>
      <td>(>= 2.0)</td>
    </tr>
    <tr>
      <td>Go</td>
      <td>(>= 1.5)</td>
    </tr>
  </tbody>
</table>

### Getting the application
```
$ git clone https://github.com/ContainerSolutions/gapp.git
```

### Building
```
$ make build
```
This will compile the application and place the generated binary under `bin/app`.

### Testing
```
$ make test
```
will run all the tests.

### Running
To run the application directly with Go:
```
$ make run
```

By default, the application will try to bind to port `8000` and will fail if the port
is not available.

If you wish to run the application using the compiled version, run it like this:

```
$ VERSION=1.0 ./bin/app
```

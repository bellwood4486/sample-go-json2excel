# sample-go-json2excel

This is a sample project to convert json to excel.

## Sample Cases

This sample loads a big json file and convert it to excel.

This project tries 3 cases for combination of loading and converting methods.

| Case#  | Loading JSON | Converting JSON to Excel | Generating intermediate objects |
|--------|--------------|--------------------------|---------------------------------|
| Case 1 | batch        | stream                   | yes                             |
| Case 2 | stream       | stream                   | yes                             |
| Case 3 | stream       | stream                   | no                              |

The data in the JSON file is as follows.

* a user has the following fields.
    * Name(string)
    * Age(int)
    * Profile(string)
* 1000000 users

## Results

The pprof results in the three cases are as follows.

| Case 1                                                                                                             | Case 2                                                                                                             | Case 3                                                                                                             |
|--------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| 238.05MB                                                                                                           | 257.08MB                                                                                                           | 20MB                                                                                                               |
| ![mem1 prof](https://user-images.githubusercontent.com/2452581/229333094-922bc58e-4578-4e85-b105-70a8ff07aaf2.png) | ![mem2 prof](https://user-images.githubusercontent.com/2452581/229333098-7c81a37e-ea61-4e84-a117-09871e2424a1.png) | ![mem3 prof](https://user-images.githubusercontent.com/2452581/229333099-3a4e067d-9154-41e0-b7a4-5938286f3218.png) |

## How to run

### Prerequisites

Generate the input JSON file by running the following command.
```bash
make gen-userlist
````

#### Run

`make memp{case#}` runs the sample.

For example, to run Case 1, run the following command.
```bash
make memp1
```

To run all cases, run the following command.
```bash
make memp_all
```

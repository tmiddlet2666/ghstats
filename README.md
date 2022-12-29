# GHSTATS

## Overview

Thus CLI displays download metrics for GitHub repositories.

* [Installing](#installing)
* [Building](#building)
* [Usage](#usage)

## Installing

## Mac AMD

```bash
curl -Lo ghstats "https://github.com/tmiddlet2666/ghstats/releases/download/0.9.0/ghstats-0.9.0-darwin-amd64"
chmod u+x ./ghstats
```

## Mac ARM

```bash
curl -Lo ghstats "https://github.com/tmiddlet2666/ghstats/releases/download/0.9.0/ghstats-0.9.0-darwin-arm64"
chmod u+x ./ghstats
```

> Note: You will need to give the `ghstats` executable permissions in Security in Settings.

### Windows AMD

```cmd
curl -Lo ghstats.exe "https://github.com/tmiddlet2666/ghstatsreleases/download/0.9.0/ghstats-0.9.0-windows-amd64.exe"
```

### Windows ARM

```cmd
curl -Lo ghstats.exe "https://github.com/tmiddlet2666/ghstatsreleases/download/0.9.0/ghstats-0.9.0-windows-arm.exe"
```

### Linux AMD

```bash
curl -Lo ghstats "https://github.com/tmiddlet2666/ghstats/releases/download/0.9.0/ghstats-0.9.0-linux-amd64"
chmod u+x ./ghstats
```

### Linux ARM

```bash
curl -Lo ghstats "https://github.com/tmiddlet2666/ghstats/releases/download/0.9.0/ghstats-0.9.0-linux-arm64"
chmod u+x ./ghstats
```


## Building

```bash
make ghstats
```

## Usage

### Help

```bash
$ ghstats get --help
The 'get' command displays one or more resources.

Usage:
  ghstats get [command]

Available Commands:
  downloads   Display the downloads for assets for a user and repository
  releases    Display the releases for a user and repository
  repo        Display the repository details for a user and repository

Flags:
  -h, --help   help for get

Global Flags:
  -r, --repository string   GitHub repository
  -u, --username string     GitHub username

Use "ghstats get [command] --help" for more information about a command.
```


### 1. Get repository details

```bash
$ ghstats -u oracle -r coherence-visualvm get repo

Repository URL: https://github.com/oracle/coherence-visualvm
Name:           coherence-visualvm
Full Name:      oracle/coherence-visualvm
Description:    The Coherence-VisualVM Plugin provides management and monitoring of a single Coherence cluster using the VisualVM management utility.
Language:       Java
Stars:          5
Subscribers:    5
Forks:          1
```

### 2. Get all releases

```bash
$ ghstats -u oracle -r coherence-visualvm get releases

Repository: https://github.com/oracle/coherence-visualvm
TAG    Name                             Pre-release?  Published                 Assets   Downloads  
1.2.0  Coherence VisualVM Plugin 1.2.0  false         2021-12-21T01:16:16Z           1         231  
1.1.0  Coherence VisualVM Plugin 1.1.0  false         2021-07-13T07:12:09Z           1       1,792  
1.0.1  Coherence VisualVM Plugin 1.0.1  false         2021-03-26T03:04:54Z           1       1,005  
1.0.0  Coherence VisualVM Plugin 1.0.0  false         2020-12-03T04:27:34Z           1       1,256  
                                                                                        ----------  
                                                                            TOTAL            4,284 
```

### 3. Getting all downloads

```bash
$ ghstats -u oracle -r coherence-operator  get downloads

Repository: https://github.com/oracle/coherence-operator
TAG     Name                                       Created                       Size   Downloads  
v3.2.4  coherence-dashboards.tar.gz                2021-09-30T17:53:54Z       49.21KB           3  
        coherence-operator-manifests.tar.gz        2021-09-30T17:53:54Z       84.39KB           2  
        coherence-operator.yaml                    2021-09-30T17:53:54Z        7.89KB          87  
        coherence.oracle.com_coherence.yaml        2021-09-30T17:53:54Z      605.61KB           2  
v3.2.3  coherence-dashboards.tar.gz                2021-09-24T13:42:30Z       49.21KB           4  
        coherence-operator-manifests.tar.gz        2021-09-24T13:42:30Z       85.28KB           4  
        coherence-operator.yaml                    2021-09-24T13:42:30Z        7.89KB           5  
        coherence.oracle.com_coherence.yaml        2021-09-24T13:42:30Z      607.81KB           1  
v3.2.2  coherence-dashboards.tar.gz                2021-09-03T09:48:55Z       49.19KB           1  
        coherence-operator-manifests.tar.gz        2021-09-03T09:48:55Z       99.89KB           7  
        coherence-operator.yaml                    2021-09-03T09:48:55Z        7.92KB          41  
        coherence.oracle.com_coherence.yaml        2021-09-03T09:48:55Z      745.30KB           3  
v3.2.1  coherence-operator-manifests.tar.gz        2021-07-27T13:31:35Z       98.33KB           1  
        coherence-operator.yaml                    2021-07-27T13:31:35Z        7.96KB          24  
        coherence.oracle.com_coherence.yaml        2021-07-27T13:31:35Z      737.07KB           2  
v3.2.0  coherence-operator-manifests.tar.gz        2021-07-01T13:03:11Z       98.36KB           3  
        coherence-operator.yaml                    2021-07-01T13:03:11Z        8.06KB          29  
v3.1.5  coherence-operator-manifests.tar           2021-04-09T09:03:11Z        1.29MB           5  
        coherence-operator.yaml                    2021-04-09T09:03:11Z        6.16KB          71  
v3.1.4  coherence-operator-manifests.tar.gz        2021-03-04T21:40:14Z      160.97KB          13  
        coherence-operator.yaml                    2021-03-04T21:40:14Z        5.44KB          27  
v3.1.3  coherence-operator-manifests-3.1.3.tar.gz  2021-01-28T12:53:08Z        1.85KB           3  
v3.1.1  coherence-operator-manifests-3.1.1.tar.gz  2020-10-13T07:59:53Z        1.85KB          10  
v3.1.0  coherence-operator-manifests-3.1.0.tar.gz  2020-09-28T09:33:52Z        1.85KB           3  
v3.0.0  converter-darwin-amd64                     2020-06-22T17:08:43Z       36.89MB           7  
        converter-linux-amd64                      2020-06-22T17:08:43Z       37.23MB           6  
        converter-windows-amd64                    2020-06-22T17:08:43Z       36.64MB           4  
                                                                                       ----------  
                                                                         TOTAL                368 
```

### 4. Getting all downloads for a specific tag

```bash
$ ghstats -u oracle -r coherence-operator -t v3.2.4 get downloads 

Repository: https://github.com/oracle/coherence-operator
TAG     Name                                 Created                       Size   Downloads  
v3.2.4  coherence-dashboards.tar.gz          2021-09-30T17:53:54Z       49.21KB           3  
        coherence-operator-manifests.tar.gz  2021-09-30T17:53:54Z       84.39KB           2  
        coherence-operator.yaml              2021-09-30T17:53:54Z        7.89KB          85  
        coherence.oracle.com_coherence.yaml  2021-09-30T17:53:54Z      605.61KB           2  
                                                                                 ----------  
                                                                   TOTAL                 92  
```

### 5. Getting downloads for a repository for a specific file

```bash
$ ghstats -u oracle -r coherence-operator -f coherence-operator.yaml get downloads

Repository: https://github.com/oracle/coherence-operator
TAG     Name                     Created                       Size   Downloads  
v3.2.4  coherence-operator.yaml  2021-09-30T17:53:54Z        7.89KB          85  
v3.2.3  coherence-operator.yaml  2021-09-24T13:42:30Z        7.89KB           5  
v3.2.2  coherence-operator.yaml  2021-09-03T09:48:55Z        7.92KB          41  
v3.2.1  coherence-operator.yaml  2021-07-27T13:31:35Z        7.96KB          24  
v3.2.0  coherence-operator.yaml  2021-07-01T13:03:11Z        8.06KB          29  
v3.1.5  coherence-operator.yaml  2021-04-09T09:03:11Z        6.16KB          71  
v3.1.4  coherence-operator.yaml  2021-03-04T21:40:14Z        5.44KB          27  
                                                                     ----------  
                                                       TOTAL                282  
```

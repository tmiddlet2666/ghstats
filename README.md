# GHSTATS

## Overview

Displays download metrics for GitHub repositories.

## Building

```bash
make ghstats
```

## Usage

### Help

```bash
$ ghstats get downloads --help
The 'get downloads' command displays the downloads for asssets for a GitHub repository.

Usage:
  ghstats get downloads [flags]

Flags:
  -f, --file-name string   file name to show downloads for
  -h, --help               help for downloads
  -t, --tag string         tag to show downloads for (default "all")

Global Flags:
  -r, --repository string   GitHub repository
  -u, --username string     GitHub username
```

### 1. Getting all downloads for a repository

```bash
$ ghstats -u oracle -r coherence-visualvm get downloads

Repository: https://github.com/oracle/coherence-visualvm
TAG    Name                                 Created                       Size   Downloads  
1.2.0  coherence-visualvm-plugin-1.2.0.nbm  2021-12-17T01:29:31Z        2.02MB         123  
1.1.0  coherence-visualvm-plugin-1.1.0.nbm  2021-07-13T02:53:34Z        2.00MB       1,792  
1.0.1  coherence-visualvm-plugin-1.0.1.nbm  2021-03-26T03:02:50Z        1.99MB       1,005  
1.0.0  coherence-visualvm-plugin-1.0.0.nbm  2020-12-02T06:24:48Z        1.95MB       1,256  
                                                                                ----------  
                                                                  TOTAL              4,176 
```

### 2. Getting downloads for a repository for a specific tag

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

### 3. Getting downloads for a repository for a specific file

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
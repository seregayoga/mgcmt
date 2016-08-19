# mgcmt
Magic commit utility for JIRA-like commit naming
Suppose you're working on task with JIRA name MYPRJCT-1234 and want to create fast commit based on your branch name.
Great, with this utility it's pretty straightforward.
# Installation
```
go get github.com/seregayoga/mgcmt
```
# Usage
```
$ git checkout -b MYPRJCT-1234-very-long-description
$ git add -A
$ mgcmt
$ git log -1
commit c8a50ffaff10c479f6cadf12ae0b673e69084b33
Author: Your Name <your@email.com>
Date:   Fri Aug 19 17:51:25 2016 +0000

    MYPRJCT-1234 very long description
```
**Magic!**

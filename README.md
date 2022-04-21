# cmdlineargs

Output all given command line arguments to STDOUT unless the `CMDLINEARGS_FILE` environment variable is set.
In that case, append the command line arguments to the file stored in this value.

## Example

```
PS R:\> .\cmdlineargs.exe aaa "with spaces" bbb ccc 'another with spaces' ddd eee

========================================================================
2022-04-21T07:20:25-04:00
[ 0] C:\Users\john\go\src\github.com\jftuga\cmdlineargs\cmdlineargs.exe
[ 1] aaa
[ 2] with spaces
[ 3] bbb
[ 4] ccc
[ 5] another with spaces
[ 6] ddd
[ 7] eee
```

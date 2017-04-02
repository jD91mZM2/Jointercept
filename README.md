# Jointercept
Jointercept stands for Join Intercept and is a replacements for the buggy EventGhost to instead launch your own programs!

Runs on:
- Linux (Tested)
- Windows (Fixed, thanks to [Kenneth Pang](https://plus.google.com/u/0/+KennethPang))
- Mac (Untested)

## Pros using Jointercept rather than a custom program:
1. Compatibility with multiple programs!!
2. Cleaner and easier code!
3. See pushes log on a website.
4. Programmed in Go. Speed is a must.
## Pros using Jointercept rather than a URL protocol:
1. Intercept all messages rather than specific ones.
2. Cross platform.

## Command line
Jointercept [port]

Examples:
```
Jointercept
```
```
Jointercept 12345
```

# Usage:
Just place your programs or scripts _(yes, you can invoke a program from anywhere on the computer using shell scripts)_ under the Join-AutoStart folder! It will now be called with the message as command line argument!

Also, make sure to set the "EventGhost Port" in Join to whatever port you start Jointercept with!

# Collatz Conjecture Iterator

## Overview
The *Collatz Conjecture* was conceived in 1937 by Lothar Collatz.  It is also known as the 3n + 1 problem.  For more detailed descriptions, see [Wikipedia](https://en.wikipedia.org/wiki/Collatz_conjecture) and a [host of other sources](https://www.google.co.uk/search?q=collatz+conjecture) that describe it in great detail.  It's frequently described as the most complex unsolved problem in mathematics.  Your opinion may differ on this; there are a lot of unsolved problems!

## Justification
There are loads of implementations of Collatz programs that iterate over ranges of numbers in the futile hope of finding the magic number that disproves the conjecture.  No doubt many of them are faster and more comprehensive than this one.  I wrote this as an exercise in learning [Google Go's](https://golang.org/) support for [big numbers](https://pkg.go.dev/math/big).  If you're aware of ways to make it faster or have suggestions for code improvements, I'd be delighted to hear from you.

## Where to start?
Obviously there are quite a lot of numbers to choose from!  If you copy the default YAML file supplied in this repo, you'll begin iterating at 295,147,905,179,352,825,856.  This equates to 2^68.  Why start there?  Why not?  Feel free to change it to anything you fancy.  Maybe you'll hit upon the crock of gold at the end of the rainbow and solve the problem.  There's not much point in starting at a lower number than this though as they've been exhaustively tested.

### A word of caution
Please don't enter into this with a sense of anticipation.  Much computing time has been devoted to the problem, including time on computers that fill large rooms.  Think of this as a bit of fun and a good excuse for writing/refining some **Go** code.

## Getting Started
* Grab a copy of [Go](https://golang.org/dl/) and follow the installation instructions.
* Clone this repo and compile the code into a **collatz** binary.
* In a directory of your choosing, create a copy of the **collatz.yml** example file and modify it in accordance with [the instructions](#Configuration).
* Run the binary!
* **Note:** the config file doubles-up as a state file.  It will get periodically overwritten as the code condems more numbers to the vast stack of numbers that fail to disprove the conjecture.

## Configuration
### Flags
By default the program will look for a configuration file called **collatz.yml** in the current directory.  This behaviour can be modified using the config flag.  E.g.:-

`/usr/local/bin/collatz --config /etc/collatz/myconfig.yml`

### YAML file
The YAML file contains a number of options that can be split into those defined by the operator and those considered stateful that are overwritten by the program.
* **start_integer** - The operator-defined number from which iteration will begin.
* **write_interval** - The operator-defined interval (in seconds) after which the program will update the state file.
* **high_integer** - The integer that has taken the most iterative steps to resolve to 1.
* **high_steps** - The number of steps it took to resolve **high_integer** to 1.
* **restart_integer** - The last resolved integer prior to updating the state file.  The program will resume from this number (if it's greater than **start_integer**).

# Good luck!
# Jitter
Keep your computer alive for a specified amount of time by jittering the mouse.

# Why?
There are many like it but this one is mine. I was bored on a Saturday during quarantine, okay?
Also, my wife's Zoom client refuses to identify itself with respect to the Linux power management
APIs. 

# Usage
Fine tune the total duration with the `-hours` and `-mins` arguments, and change
how frequently within that duration you'd like to jitter the mouse with the `-freq` argument.
By default it'll keep the computer alive for an hour and jitter the mouse every minute of that 
hour.

```
Usage of jitter:
  -freq int
  	move the mouse every N seconds (default 60)
  -hours int
  	hours (default 1)
  -mins int
  	mins (default -1)
```

So, to jitter the mouse every 5 seconds for an hour, this will suffice:

`./jitter -freq 5`

# Author
Ant Zucaro AKA Antibody AKA dfdashh in various places on the internets. 

# g3n-helpme-1
Why doesn't this work?

## Background
The demo [hellog3n](https://github.com/g3n/demos/hellog3n) works flawlessly, but I wanted +Z to be up rather than +Y.

## What I changed
I made three changes:

1.  The camera was placed at (0, 0, 5): level with the origin and 5 units toward the viewer when +Y is up.
    I changed this to (0, -5, 0): level with the origin and 5 units toward the viewer when +Z is up.

2.  The up direction is (0, 1, 0) (aka +Y) by default.  I changed it to (0, 0, 1) (aka +Z).

3.  A sphere doesn't have a visible "up" direction, so I changed it to a cylinder which tapers inward toward the top.

## What I expected
-   I expected to see a red axis helper leading from the center to the right: **as expected**
-   I expected to see a blue axis helper leading from the center to the top the same size as the red axis helper: **not as expected**, it did lead upward from the center but it was "shorter" than the red one indicating the camera was not level
-   I expected to see no green axis helper as it should be leading from the center away from the viewer: **not as expected**, it was visible pointing downward
-   I expected to see a tapered cylinder standing upright with the small end up: **not as expected**, it was upside-down and angled with the higher part (the botttom) closer to the viewer than the lower part (the top)

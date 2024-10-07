# horizon

## What is this?

## Want to contribute? Here are some things that need to be done.

- Make sure that there is no essential data loss. The youtube data api returns a lot of data. Investigate to find out if there is any important data that snapshots should be made of(that currently is not).
  
- I'm relatively new to golang so if there are any best practices that should be followed that currently are not then put up a PR and explain your changes.

- Currently no tests, there is a lot of stuff that can break so we needs tests

- Right now everything is in a single package. Ideally there should be some logical separation.
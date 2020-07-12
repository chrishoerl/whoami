# whoami

## Purpose
Simple HTTP docker service that prints its container ID
with a unique background color.

Makes it easier to tell the difference between containers
if they are scaled to many running instances on Kubernetes.

Simply press Ctrl + F5 to refresh your browser frequently
to trigger a container- and color-change.

## How to use this container image
```
# Start container
$ docker run -d -p 8000:8000 --name whoami -t chrishoerl/whoami

# Browse
http://localhost:8000
```

## Credits
Based on jwilder/whoami https://github.com/jwilder/whoami
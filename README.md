# Message in a bottle

## Getting started

> clone the repo: 
``` bash
git clone https://github.com/david-brt/miab.git
```

> start the application:
``` bash
docker-compose up -d
```
- the website should now be visible on localhost port 5173

## Designideas

### Naming of chatrooms
- chat rooms should not be named after countries but rather after known international waters, as these are topic-related.
- e.g. North Sea, Baltic Sea, Pacific or Indian Ocean

### avoidance of message accumulation 

- mark a message, that has been sent to a user as Sent/seen
- if a user sends a new Message into the system, send the id of the last obtained Message along with it, this Message can now be safely deleted(if it has not been deleted before)
- this alone will not result in a decreasion..but 
- one could set an arbitrary threshold of messages that could be kept all the time in the System (e.g. 5 messages)
- if there are a Messages above this threshold, then it will be possible to delete all the messages that have already been read
- due to the fact that there mostly will be multiple messages in the system, the possibility of sending the same message to multiple persons will be reduced

### ANIMATIONSSSSS
- when sending or receiving a new Message an animation of a guys fishing a bottle out of the sea could play(sum really crappy pixel art, I know a guy... it is me, the guy is me)

### Passwort anzeigen lassen
- beim login optional das passwort anzeigen
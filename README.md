# game-server
This is an experiment with go/grpc based game server, and electron client.
The initial version will just be a chat bot, but the intent is to have versioned and namespaced resources down the road.

## Getting Started
if you've pulled this project, and have docker installed, just run `./scripts/run`.

### Environment Variables ###
These are secret, so they are stored in a git-ignored scripts/.env

All of these are pulled in to the project in the config, look there to see what variables you can change. They should all have sane defaults.

### Running Locally ###
add you your .env `RUN_LOCAL=1`

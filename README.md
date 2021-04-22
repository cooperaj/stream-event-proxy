Stream Event Proxy
======

Stream event proxy allows you to set up stream alerts for your channel without having to use something like Streamlabs, StreamElements or Muxy.

## Todo
Currently Stream Event Proxy supports one service, Twitch. Additionally, it only alerts for new followers.

### Twitch
 - [x] Enable platform
 - [x] Followers
 - [ ] Subscriptions
 - [ ] Cheers
 - [ ] Channel Points Redemption
 - [ ] Raids
 - [ ] Hype Trains

### YouTube
 - [ ] Enable platform
 - [ ] Things YouTube viewers can do

### Other stuff
 - [ ] Make it easier to customise the alerts

## Requirements
 * You need a VPS (or similar) running with a publicly routable and HTTPS protected domain.
 * If you're wanting to change the alerts (which, to be honest, given they're used on my stream I would) you need to be some sort of web capable developer.

## Usage
You'll need to create a developer account (or similar) on your platforms of choice and create Apps that will let you use those platforms API's

Ensure that you've specified values for the following environment variables.

#### SEP_CLIENTID
The client ID for Twitch. Provided as a part of the 'app' creation process
#### SEP_CLIENTSECRET
The client secret for Twitch. Provided as a part of the 'app' creation process
#### SEP_BROADCASTERID
The unique ID of your user broadcast user on Twitch. Can be found using the twitch-cli client using the above credentials (id and secret)
#### SEP_SERVICEURL
Where your service will be living, e.g. 'https://alerts.example.com'
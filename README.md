<h1>Memorization App</h1>

<p> This Memorization apps is designed to help people retain information more effectively than traditional methods. </p>
<p> Users can sign up into the app and save the words they intend on memorizing </p>

Tools used for the app
**Traefik**: Used to implement dynamic routing.<br>
**Docker**: containerized the Golang application and set up postgresql.<br>
**Docker-compose:** :
- Setting up both Traefik container and the application container service
- Enabling Traefik to find and route traffic to the app ( "account") service based on URL patterns.
- Automatically rebuilding and restarting the app ( "account" service) on Go code changes




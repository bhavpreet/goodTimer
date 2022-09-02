heroku:
	# DOCKER_DEFAULT_PLATFORM=linux/amd64 docker buildx build --platform "linux/amd64" -t goodtimer .
	# docker tag goodtimer registryoku.com/warm-shelf-72300/goodtimer
	# docker push registry.heroku.com/warm-shelf-72300/goodtimer
	DOCKER_DEFAULT_PLATFORM=linux/amd64 heroku container:push goodtimer
	heroku container:release goodtimer
	heroku ps:scale goodtimer=1

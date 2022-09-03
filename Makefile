heroku:
	# DOCKER_DEFAULT_PLATFORM=linux/amd64 docker buildx build --platform "linux/amd64" -t goodtimer .
	# docker tag goodtimer registryoku.com/warm-shelf-72300/goodtimer
	# docker push registry.heroku.com/warm-shelf-72300/goodtimer
	DOCKER_DEFAULT_PLATFORM=linux/amd64 heroku container:push web
	heroku container:release web
	heroku ps:scale web=1

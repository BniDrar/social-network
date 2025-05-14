install:
	@sudo npm install -g json-server

serve:
	@json-server ./fake-db.json -p 3100

install2:
	@npm install json-server --save-dev

run:
	@npx json-server ./fake-db.json -p 3100
config:
	curl -fsSL https://get.docker.com/rootless | sh
	export

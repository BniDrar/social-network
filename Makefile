install:
	@sudo npm install -g json-server

fakedb:
	@json-server ./fake-db.json -p 3100

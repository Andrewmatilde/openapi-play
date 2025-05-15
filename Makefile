OPENAPI_FILE=api/openapi.yaml
GENERATOR=openapi-generator-cli
GEN_LANG=javascript
OUT_DIR=api/js-client

.PHONY: all update build clean

all: update build

update:
	$(GENERATOR) generate -i $(OPENAPI_FILE) -g $(GEN_LANG) -o $(OUT_DIR) --skip-validate-spec

build:
	cd api/js-client && npm install

clean:
	rm -rf api/js-client/dist api/js-client/src/api api/js-client/src/model api/js-client/docs api/js-client/test
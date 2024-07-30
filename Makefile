.PHONY: generate-rules-version
generate-rules-version:
	git clone --branch $(VERSION) https://github.com/open-telemetry/semantic-conventions.git ./tmp/semantic-conventions
	weaver registry generate \
		--registry ./tmp/semantic-conventions/model \
		--param version=$(VERSION) \
		go ./
	rm -rf ./tmp

.PHONY: generate-rules-all-versions
generate-rules-all-versions:
	$(MAKE) generate-rules-version VERSION=v1.26.0
	$(MAKE) generate-rules-version VERSION=v1.25.0
	$(MAKE) generate-rules-version VERSION=v1.24.0
	# $(MAKE) generate-rules-version VERSION=v1.23.1
	# $(MAKE) generate-rules-version VERSION=v1.23.0
	# $(MAKE) generate-rules-version VERSION=v1.22.0
	# $(MAKE) generate-rules-version VERSION=v1.21.0

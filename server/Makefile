.PHONY: clean

C_DIR   = dist
C_FLAGS = GOARCH=amd64

clean:
	rm -rf $(C_DIR)

server:
	$(call build_os,linux)
	$(call build_os,darwin)
	$(call build_os,windows,.exe)

define build_os
	GOOS=$(1) $(C_FLAGS) go build -o $(C_DIR)/$(1)/flux_start_server$(2)
endef

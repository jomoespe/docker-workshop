# ==========================================================================
# Copyright 2019 Jomoespe. All rights reserved.
#
# Use of this source code is governed by a The Unlicense license that can be
# found in the LICENSE file.
# ==========================================================================

# ----------------------------------------------------------------------
# First stage: build the binary program
# ----------------------------------------------------------------------

FROM golang:1.17.2 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /
COPY hello-world.go .
RUN go build -a -installsuffix cgo hello-world.go 

# ----------------------------------------------------------------------
# Second stage: build the container from binary built in previous stage
# ----------------------------------------------------------------------

FROM scratch
WORKDIR /
COPY --from=builder hello-world .
ENTRYPOINT [ "/hello-world" ]

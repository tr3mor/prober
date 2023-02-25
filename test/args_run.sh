#!/bin/bash
docker pull stasky745/prober:latest
docker run -it --rm --name prober stasky745/prober ./config/config.yaml

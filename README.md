# SynBioHub Docker Setups
## General
The docker-compose files in this repository represent various configurations for deploying SynBioHub.
The files can be layered with Docker Compose's [multiple file](https://docs.docker.com/compose/reference/overview/#specifying-multiple-compose-file) capabilities. 

The base configuration, described with `docker-compose.yml`, is simply SynBioHub, its graph database Virtuoso, and an autohealer.

To run the base configuration:
1. Open terminal
2. `git clone https://github.com/synbiohub/synbiohub-docker`
3. `docker-compose --f ./synbiohub-docker/docker-compose.yml up`

### With Explorer
To add [SBOLExplorer](https://github.com/michael13162/SBOLExplorer), add the `docker-compose.explorer.yml` to the main docker-compose, i.e. for step 3 run `docker-compose --f ./synbiohub-docker/docker-compose.yml -f ./synbiohub-docker/docker-compose.explorer.yml up`

### With Plugins
To add plugins to the configuration change step 3 to: `docker-compose --f ./synbiohub-docker/docker-compose.yml -f ./synbiohub-docker/docker-compose.explorer.yml -f ./synbiohub-docker/docker-compose.<Plugin 1 File Name>.yml -f ./synbiohub-docker/docker-compose.<Plugin 2 File Name>.yml up`

Note that all plugins are added before the `up` and each is preceeded by `-f `. For example, to run the configuration with the VisualIgem plugins and the VisualSeqviz plugin run: `docker-compose --f ./synbiohub-docker/docker-compose.yml -f ./synbiohub-docker/docker-compose.explorer.yml -f ./synbiohub-docker/docker-compose.pluginVisualIgem.yml -f ./synbiohub-docker/docker-compose.pluginVisualSeqviz.yml up`

### With Developmental SynBioHub
The `docker-compose.version.yml` can be added to another configuration, and simply contains the latest version of the SynBioHub docker image. 
This version does not even contain the Virtuoso image, so it should only be used by someone who knows what they are doing. 

## Plugins
For full plugin information please see this [table](https://synbiohub.github.io/synbiohub-docker/) which contains a sortable table with plugin overview.

## Plugin Ports Used
Note that thee are other ports that are already in use: synbiohub:7777, virtuoso:8890, elasitcsearch:9200, and sbolexplorer:13162.
 - docker-compose.pluginVisualIgem.yml : 3000
 - docker-compose.pluginVisualComponentUse.yml : 8080
 - docker-compose.pluginVisualTest.yml : 8081
 - docker-compose.pluginVisualTestJS.yml : 8082
 - docker-compose.pluginDownloadSnapgene.yml : 8083
 - docker-compose.pluginSubmitSnapgene.yml : 8084
 - docker-compose.pluginVisualSeqviz.yml : 8085
 - docker-compose.pluginVisualServelet.yml : 8086
 - docker-compose.pluginSubmitTest.yml : 8087

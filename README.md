# SynBioHub Docker Setups
## General

The docker-compose files in this repository represent various configurations for deploying SynBioHub.
The files can be layered with Docker Compose's [multiple file](https://docs.docker.com/compose/reference/overview/#specifying-multiple-compose-file) capabilities. 

The base configuration, described with `docker-compose.yml`, is simply SynBioHub, its graph database Virtuoso, and an autohealer.

To run the base configuration, simply use `docker-compose up`.

To add [SBOLExplorer](https://github.com/michael13162/SBOLExplorer), add the `docker-compose.explorer.yml` to the main docker-compose.

To run the configuration with SBOLExplorer, use `docker-compose -f docker-compose.yml -f docker-compose.explorer.yml up`.

To add plugins to the configuration use `docker-compose -f docker-compose.yml -f <plugin 1 file name> -f <plugin 2 file name> -f <plugin 3 file name> up`. Note that all plugins are added before the `up` and each is preceeded by `-f `. For example, to run the configuration with the VisualIgem plugins and the VisualSeqviz plugin run: `docker-compose -f docker-compose.yml -f docker-compose.pluginVisualIgem.yml -f docker-compose.pluginVisualSeqviz up`.

The `docker-compose.version.yml` can be added to another configuration, and simply contains the latest version of the SynBioHub docker image. 
This version does not even contain the Virtuoso image, so it should only be used by someone who knows what they are doing. 

## Description of available plugins
 - **DownloadSnapgene** : For more information see [this](https://github.com/SynBioHub/Plugin-Download-Snapgene)
 - **SubmitSnapgene** : For more information see: [this](https://github.com/SynBioHub/Plugin-Submit-Snapgene)
 - **SubmitTest** : A python TEST submit plugin which simply indicates that submit plugins are working and provides a framework to play with for plugin developers. For more information see: [this](https://github.com/SynBioHub/Plugin-Submit-Test)
 - **VisualComponentUse** : The visualisation plugin(s) containing a co-use component sankey diagram, and the most used components bar graph endpoints. For more information see: [this](https://github.com/SynBioHub/Plugin-Visual-Component-Use)
 - **VisualIgem** : The visualisation plugin(s) containing endpoints for iGEM Main Page, iGEM Design Page, and iGEM Experience Page. For more information see: 
 - **VisualSeqviz** : The visualisation plugin which shows the plasmid view and sequence view of components. For more information see: [this](https://github.com/alicelh/sequence-view-plugin)
 - **VisualServelet** : A javascript TEST visualisation plugin which allows testing of file serving and provides a framework to play with for plugin developers. For more information see: [this](https://github.com/SynBioHub/Plugin-Visual-Serve-Test-js)
 - **VisualTest** : A python TEST visualisation plugin which simply indicates that visualisation plugins are working and provides a framework to play with for plugin developers. For more information see: [this](https://github.com/SynBioHub/Plugin-Visual-Test)
 - **VisualTestJS** : A javascript TEST visualisation plugin which simply indicates that submit plugins are working and provides a framework to play with for plugin developers. For more information see: [this](https://github.com/SynBioHub/Plugin-Visual-Test-js)

## Plugin Ports Used
Note that thee are other ports that are already in use: synbiohub:7777, virtuoso:8890, elasitcsearch:9200, and sbolexplorer:13162.
### In alphabetical order
 - docker-compose.pluginDownloadSnapgene.yml : 8083
 - docker-compose.pluginSubmitSnapgene.yml : 8084
 - docker-compose.pluginSubmitTest.yml : 8087
 - docker-compose.pluginVisualComponentUse.yml : 8080
 - docker-compose.pluginVisualIgem.yml : 3000
 - docker-compose.pluginVisualSeqviz.yml : 8085
 - docker-compose.pluginVisualServelet.yml : 8086
 - docker-compose.pluginVisualTest.yml : 8081
 - docker-compose.pluginVisualTestJS.yml : 8082
 ### In port order
 - docker-compose.pluginVisualIgem.yml : 3000
 - docker-compose.pluginVisualComponentUse.yml : 8080
 - docker-compose.pluginVisualTest.yml : 8081
 - docker-compose.pluginVisualTestJS.yml : 8082
 - docker-compose.pluginDownloadSnapgene.yml : 8083
 - docker-compose.pluginSubmitSnapgene.yml : 8084
 - docker-compose.pluginVisualSeqviz.yml : 8085
 - docker-compose.pluginVisualServelet.yml : 8086
 - docker-compose.pluginSubmitTest.yml : 8087

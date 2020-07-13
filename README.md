# SynBioHub Docker Setups
## Updated March 20, 2020

The docker-compose files in this repository represent various configurations for deploying SynBioHub.
The files can be layered with Docker Compose's [multiple file](https://docs.docker.com/compose/reference/overview/#specifying-multiple-compose-file) capabilities. 

The base configuration, described with `docker-compose.yml`, is simply SynBioHub, its graph database Virtuoso, and an autohealer.

To run the base configuration, simply use `docker-compose up`.

To add [SBOLExplorer](https://github.com/michael13162/SBOLExplorer), add the `docker-compose.explorer.yml` to the main docker-compose.

To run the configuration with SBOLExplorer, use `docker-compose -f docker-compose.yml -f docker-compose.explorer.yml up`.

To add iGEM rendering plugins, add the `docker-compose.igem.yml`.

To run the configuration with iGEM, use `docker-compose -f docker-compose.yml -f docker-compose.pluginIgem.yml up`.

The `docker-compose.version.yml` can be added to another configuration, and simply contains the latest version of the SynBioHub docker image. 
This version does not even contain the Virtuoso image, so it should only be used by someone who knows what they are doing. 

## Description of available plugins
 - docker-compose.pluginVisualComponentUse.yml : The visualisation plugin(s) containing a co-use component sankey diagram, and the most used components bar graph endpoints. For more information see: https://github.com/SynBioHub/Plugin-Visual-Component-Use
 - docker-compose.pluginVisualIgem.yml : The visualisation plugin(s) containing endpoints for iGEM Main Page, iGEM Design Page, and iGEM Experience Page. For more information see: 
 - docker-compose.pluginVisualSeqviz.yml : The visualisation plugin which shows the plasmid view and sequence view of components. For more information see: https://github.com/alicelh/sequence-view-plugin
 - docker-compose.pluginVisualServelet.yml : A javascript TEST visualisation plugin which allows testing of file serving and provides a framework to play with for plugin developers. For more information see: https://github.com/SynBioHub/Plugin-Visual-Serve-Test-js
 - docker-compose.pluginDownloadSnapgene.yml : For more information see: https://github.com/SynBioHub/Plugin-Download-Snapgene
 - docker-compose.pluginSubmitSnapgene.yml : For more information see: https://github.com/SynBioHub/Plugin-Submit-Snapgene
 - docker-compose.pluginSubmitTest.yml : A python TEST submit plugin which simply indicates that submit plugins are working and provides a framework to play with for plugin developers. For more information see: https://github.com/SynBioHub/Plugin-Submit-Test
 - docker-compose.pluginVisualTest.yml : A python TEST visualisation plugin which simply indicates that visualisation plugins are working and provides a framework to play with for plugin developers. For more information see: https://github.com/SynBioHub/Plugin-Visual-Test
 - docker-compose.pluginVisualTestJS.yml : A javascript TEST visualisation plugin which simply indicates that submit plugins are working and provides a framework to play with for plugin developers. For more information see: https://github.com/SynBioHub/Plugin-Visual-Test-js

## Plugin Ports Used
Note that thee are other ports that are already in use: synbiohug:7777, virtuoso:8890, elasitcsearch:9200, and sbolexplorer:13162.
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

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

## Plugins
<html>
	<body>
		<div style="overflow-x:auto;">
			<table id="PluginTable">
				<tr>
					<!--When a header is clicked, run the sortTable function, with a parameter,
					0 for sorting by names, 1 for sorting by country: -->
					<th onclick="sortTable(0)">Name</th>
					<th onclick="sortTable(1)">Type</th>
					<th onclick="sortTable(2)">Port</th>
					<th onclick="sortTable(3)">Language</th>
					<th onclick="sortTable(4)">Test</th>
					<th onclick="sortTable(5)">Description</th>
					<th onclick="sortTable(6)">Multi-Endpoint</th>
					<th onclick="sortTable(7)">Repository</th>
					<th onclick="sortTable(8)">File</th>
				</tr>
				<tr>
					<td>DownloadSnapgene</td>
					<td>Download</td>
					<td>8083</td>
					<td>Python</td>
					<td>No</td>
					<td></td>
					<td></td>
					<td><a href="https://github.com/SynBioHub/Plugin-Download-Snapgene">Plugin-Download-Snapgene</a></td>
					<td>docker-compose.pluginDownloadSnapgene.yml</td>
				</tr>
				<tr>
					<td>SubmitSnapgene</td>
					<td>Submit</td>
					<td>8084</td>
					<td>Python</td>
					<td>No</td>
					<td></td>
					<td></td>
					<td><a href="https://github.com/SynBioHub/Plugin-Submit-Snapgene">Plugin-Submit-Snapgene</a></td>
					<td>docker-compose.pluginSubmitSnapgene.yml</td>
				</tr>
				<tr>
					<td>SubmitTest</td>
					<td>Submit</td>
					<td>8087</td>
					<td>Python</td>
					<td>Yes</td>
					<td>Simply indicates that submit plugins are working and provides a framework to play with for plugin developers</td>
					<td>No</td>
					<td><a href="https://github.com/SynBioHub/Plugin-Submit-Test">Plugin-Submit-Test</a></td>
					<td>docker-compose.pluginSubmitTest.yml</td>
				</tr>
				<tr>
					<td>VisualComponentUse</td>
					<td>Visual</td>
					<td>8080</td>
					<td>Python</td>
					<td>No</td>
					<td>Containing a co-use component sankey diagram, and the most used components bar graph endpoints</td>
					<td>Yes</td>
					<td><a href="https://github.com/SynBioHub/Plugin-Visual-Component-Use">Plugin-Visual_component-Use</a></td>
					<td>docker-compose.pluginVisualComponentUse.yml</td>
				</tr>
				<tr>
					<td>VisualIgem</td>
					<td>Visual</td>
					<td>3000</td>
					<td></td>
					<td>No</td>
					<td>Containing endpoints for iGEM Main Page, iGEM Design Page, and iGEM Experience Page</td>
					<td>Yes</td>
					<td><a href=""></a></td>
					<td>docker-compose.pluginVisualIgem.yml</td>
				</tr>
				<tr>
					<td>VisualSeqviz</td>
					<td>Visual</td>
					<td>8085</td>
					<td>Javascript</td>
					<td>No</td>
					<td>Shows the plasmid view and sequence view of components</td>
					<td>No</td>
					<td><a href="https://github.com/alicelh/sequence-view-plugin">sequence-view-plugin</a></td>
					<td>docker-compose.pluginVisualSeqviz.yml</td>
				</tr>
				<tr>
					<td>VisualServelet</td>
					<td>Visual</td>
					<td>8086</td>
					<td>Javascript</td>
					<td>Yes</td>
					<td>Allows testing of file serving and provides a framework to play with for plugin developers</td>
					<td>No</td>
					<td><a href="https://github.com/SynBioHub/Plugin-Visual-Serve-Test-js">Plugin-Visual-Serve-Test-js</a></td>
					<td>docker-compose.pluginVisualServelet.yml</td>
				</tr>
				<tr>
					<td>VisualTest</td>
					<td>Visual</td>
					<td>8081</td>
					<td>Python</td>
					<td>Yes</td>
					<td>Smply indicates that visualisation plugins are working and provides a framework to play with for plugin developers</td>
					<td>No</td>
					<td><a href="https://github.com/SynBioHub/Plugin-Visual-Test">Plugin-Visual-Test</a></td>
					<td>docker-compose.pluginVisualTest.yml</td>
				</tr>
				<tr>
					<td>VisualTestJS</td>
					<td>Visual</td>
					<td>8082</td>
					<td>Javascript</td>
					<td>Yes</td>
					<td>Aimply indicates that submit plugins are working and provides a framework to play with for plugin developers</td>
					<td>No</td>
					<td><a href="https://github.com/SynBioHub/Plugin-Visual-Test-js">Plugin-Visual-Test-js</a></td>
					<td>docker-compose.pluginVisualTestJS.yml</td>
				</tr>

			</table>
		</div>
		<script>
			function sortTable(n) {
			  var table, rows, switching, i, x, y, shouldSwitch, dir, switchcount = 0;
			  table = document.getElementById("PluginTable");
			  switching = true;
			  // Set the sorting direction to ascending:
			  dir = "asc";
			  /* Make a loop that will continue until
			  no switching has been done: */
			  while (switching) {
				// Start by saying: no switching is done:
				switching = false;
				rows = table.rows;
				/* Loop through all table rows (except the
				first, which contains table headers): */
				for (i = 1; i < (rows.length - 1); i++) {
				  // Start by saying there should be no switching:
				  shouldSwitch = false;
				  /* Get the two elements you want to compare,
				  one from current row and one from the next: */
				  x = rows[i].getElementsByTagName("TD")[n];
				  y = rows[i + 1].getElementsByTagName("TD")[n];
				  /* Check if the two rows should switch place,
				  based on the direction, asc or desc: */
				  if (dir == "asc") {
					if (x.innerHTML.toLowerCase() > y.innerHTML.toLowerCase()) {
					  // If so, mark as a switch and break the loop:
					  shouldSwitch = true;
					  break;
					}
				  } else if (dir == "desc") {
					if (x.innerHTML.toLowerCase() < y.innerHTML.toLowerCase()) {
					  // If so, mark as a switch and break the loop:
					  shouldSwitch = true;
					  break;
					}
				  }
				}
				if (shouldSwitch) {
				  /* If a switch has been marked, make the switch
				  and mark that a switch has been done: */
				  rows[i].parentNode.insertBefore(rows[i + 1], rows[i]);
				  switching = true;
				  // Each time a switch is done, increase this count by 1:
				  switchcount ++;
				} else {
				  /* If no switching has been done AND the direction is "asc",
				  set the direction to "desc" and run the while loop again. */
				  if (switchcount == 0 && dir == "asc") {
					dir = "desc";
					switching = true;
				  }
				}
			  }
			}
		</script>
	</body>
</html>

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

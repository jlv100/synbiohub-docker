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

Note that all plugins are added before the `up` and each is preceeded by `-f `. For example, to run the configuration with the VisualIgem plugins and the VisualSeqviz plugin run:

`docker-compose --f ./synbiohub-docker/docker-compose.yml -f ./synbiohub-docker/docker-compose.explorer.yml -f ./synbiohub-docker/docker-compose.pluginVisualIgem.yml -f ./synbiohub-docker/docker-compose.pluginVisualSeqviz.yml up`

### With Developmental SynBioHub
The `docker-compose.version.yml` can be added to another configuration, and simply contains the latest version of the SynBioHub docker image. 
This version does not even contain the Virtuoso image, so it should only be used by someone who knows what they are doing. 

## Plugins
This is a sortable table with plugin information. Click on the header to sort by that column.
<html>
	<head>
		<title>Plugin Table</title>
		<link rel="stylesheet" type="text/css" href="https://www.jqueryscript.net/demo/DataTables-Jquery-Table-Plugin/media/css/jquery.dataTables.css">
	</head>
	<body>
		<table id="example" class="display" cellspacing="0" width="100%">
			<thead>
				<tr>
					<th style="visibility: hidden;">Name</th>
					<th>Type</th>
					<th style="visibility: hidden;">Port</th>
					<th>Language</th>
					<th>Test</th>
					<th style="visibility: hidden;">Description</th>
					<th>Multi-Endpoint</th>
					<th style="visibility: hidden;">Repository</th>
					<th style="visibility: hidden;">File</th>
				</tr>
				<tr>
					<th>Name</th>
					<th>Type</th>
					<th>Port</th>
					<th>Language</th>
					<th>Test</th>
					<th>Description</th>
					<th>Multi-Endpoint</th>
					<th>Repository</th>
					<th>File</th>
				</tr>
			</thead>
			<tbody>
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
			</tbody>
		</table>
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.0/jquery.min.js"></script>
		<script type="text/javascript" src="https://cdn.datatables.net/1.10.16/js/jquery.dataTables.min.js"></script>
		<script>
			$(document).ready(function() {
				var table = $('#example').DataTable({
					paging:false,
					orderCellsTop:false,
					fixedHeader: true,
					scrollX: false
				});
				$('#example thead tr:eq(0) th').each( function (i) {
					var select = $('<select><option value=""></option></select>')
						.appendTo( $(this).empty() )
						.on( 'change', function () {
							table.column( i )
								.search( $(this).val() )
								.draw();
						} );
					table.column( i ).data().unique().sort().each( function ( d, j ) {
						select.append( '<option value="'+d+'">'+d+'</option>' )
					} );
				} );
			} );
		</script>
	</body>
</html>

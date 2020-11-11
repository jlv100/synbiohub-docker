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
					<td><a href="https://github.com/SynBioHub/Plugin-Visual-Component-Use">Plugin-Visual-Component-Use</a></td>
					<td>docker-compose.pluginVisualComponentUse.yml</td>
				</tr>
				<tr>
					<td>VisualIgem</td>
					<td>Visual</td>
					<td>3000</td>
					<td>TypeScript</td>
					<td>No</td>
					<td>Containing endpoints for iGEM Main Page, iGEM Design Page, and iGEM Experience Page</td>
					<td>Yes</td>
					<td><a href="https://github.com/SynBioHub/Plugin-Visual-Igem">Plugin-Visual-Igem</a></td>
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

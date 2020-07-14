## Plugins
This is a sortable table with plugin information. Click on the header to sort by that column.
<html>
	<style>
		table {
		  border-spacing: 0;
		  width: 100%;
		  border: 1px solid #ddd;
		}
		th {
		  cursor: pointer;
		}
		th, td {
		  text-align: left;
		  padding: 14px;
		}
		tr:nth-child(even) {
		  background-color: #f2f2f2
		}
	</style>
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

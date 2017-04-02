package main

const TEMPLATE_CODE = `
<!DOCTYPE html>
<html>
	<head>
		<title>Jointercept</title>
		<link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">
		<style>
body{
	background: #0097C9;
}
#contents{
	margin: 40vh 30% 0 30%;
	background: white;
	padding: 20px;
	min-height: 100vh;
	font-size: 20px;
	font-family: 'Roboto', sans-serif;
}
table{
	font-size: 12px;
	width: 100%;
	margin-top: 40px;
	border-top: 1px solid black;
	table-layout: fixed;
}
td:first-child{ width: 80%; overflow: hidden; }
td:last-child{  width: 20%; text-align: right; }
tr:hover{
	background: #BFBFBF;
}
#contents img{
	width: 20px;
	height: 20px;
}
		</style>
		<script>
setInterval(function(){
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if(xhttp.readyState == 4 && this.status == 200) {
			document.getElementsByTagName("table")[0].innerHTML = this.responseText;
		}
	}
	xhttp.open("GET", "/table", true);
	xhttp.send();
}, 3000);
		</script>
		<meta charset="UTF-8" />
	</head>
	<body>
		<div id="contents">
			Hello! <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAABmJLR0QA/QD/AGZTrt8TAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH4QMUCCU4R/XhXgAAABl0RVh0Q29tbWVudABDcmVhdGVkIHdpdGggR0lNUFeBDhcAAAIISURBVFjD7Zc9TxtBEIafsa0oHXDQpImgoDMSHTRpUqWhcEHlgsbQ5A9EobAoSJWWxj5FSkHlP4AUiQZFQNqkQUSKRLp8rNJEYJC9KdgkeG93bw8jPiSmufPO7L2zM+/OjOFeblgk1rCnV6aBRaABTHnMvgAp0ClL6/BKHDDATaBe8HCbwFqeI5ID3gA2gAeXjPAp8LwsrdRnUAqAvwTaQ4Bj9rbNt+IjYE7evmK+LbsiIZ6cf8qe/BnCUwA028AWxfScAlWbE64UNLPgc5SoIYwgjFCiBswV0P9LRzPIAXP6ejZMTwit5ekvSN1geCOw6E7f4/OPiiAiA2txej+G7UDjGopfI+SAp8Id5awdRe7JYpRiXNbsnD+1Rms9sBajD0klLmr79Bmzrtl+AX1kKe7pFX0dHbAsLSkUAaVOBn4nycOh7AqlQKkTJpK3A2s/1JLT1mWX50TF0c+nsiT8iDBj3n8xnrwAekDXWHSBCpqvwCOECpp3BGYGrwMpsO7I2gXSjAKjEYk78ynS0DXsuPd8vwTVej5Fx+uA6VSbNpF+qjf0eY3mM3F1o4vmm3NKsrthdDv+y/AkmUWoIUxaoH3gAM0eSu0CXZuAznZceCD578g8wgJwjOYDSr0HfoeuX9xAYo1k68PWBiOrZWm9ultDqSFlClRtYhYYy6sh8Fvxx+Reblz+AA+bxckQbH28AAAAAElFTkSuQmCC" /><br />
			This page is currently being used by Jointercept.<br />
			You can always change ports in arguments!

			<table>
` + TEMPLATE_CODE_TABLE + `
			</table>
		</div>
	</body>
</html>
`

const TEMPLATE_CODE_TABLE = `
{{range .}}
<tr>
	<td>{{.Msg}}</td>
	<td>{{.Time}}</td>
</tr>
{{end}}
`

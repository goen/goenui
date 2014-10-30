package main

const (
	bingraph = `
<html>
<head>
<style type="text/css">
  #container1 {
	display: block;
    max-width:400px;
    height: 400px;

  }
  #container2 {
	display: block;
    max-width:400px;
    height: 400px;

  }
</style>
</head>
<body>
<div id="container1"></div>
<div id="container2"></div>
<script src="v/sigmajs/release-v1.0.3/sigma.min.js"></script>
<script src="v/sigmajs/release-v1.0.3/plugins/sigma.parsers.json.min.js"></script>
<script>
  sigma.parsers.json('d1/o.json', {
    container: 'container1',
    settings: {
      defaultNodeColor: '#ec5148'
    }
  });

  sigma.parsers.json('d1/o.json', {
    container: 'container2',
    settings: {
      defaultNodeColor: '#ec5148'
    }
  });
</script>
</body>
</html>`

	stepsgraph = `
<html>
<head>
</head>
<body>
<table>
<tr>

</tr>


</table>
</body>
</html>`

)

package main

const (
	bingraph = `
<html>
<head>
<style type="text/css">
  #container {
    max-width: 400px;
    height: 400px;
    margin: auto;
  }
</style>
</head>
<body>
<div id="container"></div>
<script src="v/sigmajs/release-v1.0.3/sigma.min.js"></script>
<script src="v/sigmajs/release-v1.0.3/plugins/sigma.parsers.json.min.js"></script>
<script>
  sigma.parsers.json('data.json', {
    container: 'container',
    settings: {
      defaultNodeColor: '#ec5148'
    }
  });
</script>
</body>
</html>`
)

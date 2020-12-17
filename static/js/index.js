mapboxgl.accessToken = 'pk.eyJ1IjoibWVsY2hzZWUiLCJhIjoiY2tpaGZmeHRxMDNlODJ0bXJpOXE3MmxnZSJ9.wpw5RxLlikYAFbHs7ljHlQ';
  var map = new mapboxgl.Map({
    container: 'mapid', // container id
    style: 'mapbox://styles/melchsee/ckihfibig6i9z19k0nxc6fpcs',
    center: [103.851959, 1.340270], // set starting position
    zoom: 12, // starting zoom
  });

  // Add geolocate control to the map.
  map.addControl(
    new mapboxgl.GeolocateControl({
      positionOptions: {
        enableHighAccuracy: true
      },
      trackUserLocation: true
    })
  );

  // Add zoom and rotation controls to the map.
  map.addControl(new mapboxgl.NavigationControl());

  var size = 80;

// implementation of CustomLayerInterface to draw a pulsing dot icon on the map
// see https://docs.mapbox.com/mapbox-gl-js/api/#customlayerinterface for more info
var pulsingDot = {
  width: size,
  height: size,
  data: new Uint8Array(size * size * 4),
   
  // get rendering context for the map canvas when layer is added to the map
  onAdd: function () {
  var canvas = document.createElement('canvas');
  canvas.width = this.width;
  canvas.height = this.height;
  this.context = canvas.getContext('2d');
  },
   
  // called once before every frame where the icon will be used
  render: function () {
  var duration = 1000;
  var t = (performance.now() % duration) / duration;
   
  var radius = (size / 2) * 0.3;
  var outerRadius = (size / 2) * 0.7 * t + radius;
  var context = this.context;
   
  // draw outer circle
  context.clearRect(0, 0, this.width, this.height);
  context.beginPath();
  context.arc(
  this.width / 2,
  this.height / 2,
  outerRadius,
  0,
  Math.PI * 2
  );
  context.fillStyle = 'rgba(255, 200, 200,' + (1 - t) + ')';
  context.fill();
   
  // draw inner circle
  context.beginPath();
  context.arc(
  this.width / 2,
  this.height / 2,
  radius,
  0,
  Math.PI * 2
  );
  context.fillStyle = 'rgba(255, 100, 100, 1)';
  context.strokeStyle = 'white';
  context.lineWidth = 2 + 4 * (1 - t);
  context.fill();
  context.stroke();
   
  // update this image's data with data from the canvas
  this.data = context.getImageData(
  0,
  0,
  this.width,
  this.height
  ).data;
   
  // continuously repaint the map, resulting in the smooth animation of the dot
  map.triggerRepaint();
   
  // return `true` to let the map know that the image was updated
  return true;
  }
  };

  map.on('load', function () {
    map.addImage('pulsing-dot', pulsingDot, { pixelRatio: 2 });
     
    map.addSource('points', {
      type: 'geojson',
      data: '/json/pointer.geojson',
      cluster: true,
      clusterMaxZoom: 14, // Max zoom to cluster points on
      clusterRadius: 20, // Radius of each cluster when clustering points (defaults to 50)
    });
    
    // add pulsing dot layer
    map.addLayer({
    'id': 'points',
    'type': 'symbol',
    'source': 'points',
    'layout': {
    'icon-image': 'pulsing-dot'
    }
    });

  

  });

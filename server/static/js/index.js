mapboxgl.accessToken = 'pk.eyJ1IjoibWVsY2hzZWUiLCJhIjoiY2tpaGZmeHRxMDNlODJ0bXJpOXE3MmxnZSJ9.wpw5RxLlikYAFbHs7ljHlQ';
  var map = new mapboxgl.Map({
    container: 'mapid', // container id
    style: 'mapbox://styles/melchsee/ckihfibig6i9z19k0nxc6fpcs',
    center: [103.851959, 1.340270], // set starting position
    zoom: 12 // starting zoom
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


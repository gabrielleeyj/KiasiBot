let mymap = L.map('mapid').fitWorld();

L.tileLayer('https://api.mapbox.com/styles/v1/melchsee/ckihfibig6i9z19k0nxc6fpcs.html?fresh=true&title=view&access_token=pk.eyJ1IjoibWVsY2hzZWUiLCJhIjoiY2tpaGZmeHRxMDNlODJ0bXJpOXE3MmxnZSJ9.wpw5RxLlikYAFbHs7ljHlQ',{
attribution:'Map data &copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
maxZoom: 18,
tileSize: 512,
zoomOffset: -1,
}).addTo(map);

map.locate({setView: true, maxZoom: 16});

function onLocationFound(e) {
    var radius = e.accuracy;

    L.marker(e.latlng).addTo(map)
        .bindPopup("You are within " + radius + " meters from this point").openPopup();

    L.circle(e.latlng, radius).addTo(map);
}

map.on('locationfound', onLocationFound);

function onLocationError(e) {
    alert(e.message);
}

map.on('locationerror', onLocationError);


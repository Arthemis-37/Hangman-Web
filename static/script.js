var button = document.getElementById('imageButton');
var buttonImage = document.getElementById('buttonImage');

button.addEventListener('mouseover', function() {
    buttonImage.src = '../static/Image/play_button400.png';
 });

button.addEventListener('mouseout', function() {
    buttonImage.src = '../static/Image/play_button2_400.png';
});

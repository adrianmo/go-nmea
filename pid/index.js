
var GetGraph = function() {

  // Populate image query attributes. Look for all 'slider' class elements
  // and query those for the value.
  var imgAttrs = {}
  $('.slider').each(function(idx, element) {
    var name = element.id.split('-')[0];
    imgAttrs[name] = $('#' + element.id).slider("value");
  });

  // Create the tentative img tag, set callback to put it in the page.
  var img = new Image();
  img.onload = function() {
      var graph = $('.graph')[0];
      graph.removeChild(graph.children[0]);
      graph.appendChild(img);
  };
  img.src = '/graph?' + $.param(imgAttrs)
}

// Add a slider to the page.
addSlider = function(element, name, title, min, max, step, value, unit) {
  var sliderBox = document.createElement('div');
  sliderBox.className = 'slider-box';
  element.appendChild(sliderBox);

  var sliderHeader = document.createElement('div');
  sliderHeader.innerText = title;
  sliderBox.appendChild(sliderHeader);

  var slider = document.createElement('div');
  slider.className = 'slider';
  slider.id = name + '-slider';
  sliderBox.appendChild(slider);

  var sliderValue = document.createElement('div');
  sliderValue.className = 'slider-value';
  sliderValue.id = name + '-value';
  sliderValue.innerText = value + unit;
  sliderBox.appendChild(sliderValue);

  $('#' + name + '-slider').slider({min: min, max: max, step: step, value: value});
  $('#' + name + '-slider').on("slide", function(event, ui) {
    $('#' + name + '-value').html(ui.value + unit);
    GetGraph();
  });

}

$(document).ready(function() {
  console.log("Ready!");
  $('.graph').html("<img src='/graph'/>");
  addSlider($('#tabs-1')[0], "setpoint", "Set Point", 30, 100, 1, 80, " deg");
  addSlider($('#tabs-1')[0], "maxpower", "Max Power", 100, 2500, 100, 2000, "W");
  addSlider($('#tabs-1')[0], "fluctuation", "Fluctuation", 0, 200, 2, 20, "%");
  addSlider($('#tabs-1')[0], "inertia", "Thermal Inertia", 0, 1000, 10, 1000, "W/s");

  addSlider($('#tabs-2')[0], "kp", "Kp", 1, 10000, 100, 6000, "");
  addSlider($('#tabs-2')[0], "ki", "Ki", 0, 500, 10, 25, "");
  addSlider($('#tabs-2')[0], "kd", "Kd", 0, 10000, 100, 5000, "");

  addSlider($('#tabs-3')[0], "volume", "Volume", 0, 30, 1, 10, "L");
  addSlider($('#tabs-3')[0], "loss", "Thermal Loss", 0, 30, 1, 13, "W/deg");

  $('#control-tabs').tabs();
  GetGraph();
});


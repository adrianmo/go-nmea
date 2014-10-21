
var GetGraph = function() {

  // Populate image query attributes. Look for all 'slider' class elements
  // and query those for the value.
  var imgAttrs = {systemName: $('#system-selector')[0].value};
  $('.slider').each(function(idx, element) {
    imgAttrs[element.formName] = $('#' + element.id).slider("value");
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
  slider.formName = element.id + '_' + name;
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
  $('#control-tabs').tabs();
  $('#system-selector').change(function(event) {
    loadSystem(event.target.value);
  });
  loadSystem('');
  GetGraph();
});

loadSystem = function(name) {
  var allSystems = $.getJSON('systems', function(data) {
    clearSystem();
    // Add all the received ones.
    for (var system in data) {
      displaySystem(system, data[system], name == '' || system == name);
    }
    $('#control-tabs').tabs();
  });
}

clearSystem = function() {
  // Delete existing options.
  $('#control-tabs').tabs('destroy');
  $('.system-selector-option').remove()
  $('.control-tab').remove()
  $('.slider').slider('destroy');
  $('.sliders').remove()
}


displaySystem = function(name, system, selected) {
  var selector = new Option(system.description, name);
  selector.className = 'system-selector-option';
  $('#system-selector')[0].add(selector);
  if (selected) {
    selector.selected = true;
    for (var param in system) {
      if (param.indexOf('description') < 0) {
        console.log('Add ' + param);
        addTabAndParameters(param.split('_')[0], system[param]);
      }
    }
  }
}

addTabAndParameters = function(name, parameters) {
  var tabList = $('#control-tabs-list')[0];
  var newTab = document.createElement('li');
  newTab.className = 'control-tab';
  var newTabA = document.createElement('a');
  newTabA.href = '#' + name;
  newTabA.innerText = name;
  newTab.appendChild(newTabA);
  tabList.appendChild(newTab);

  var tabTop = $('#control-tabs')[0];
  var tabDiv = document.createElement('div');
  tabTop.appendChild(tabDiv);
  tabDiv.className = 'sliders';
  tabDiv.id = name;
  for (var i = 0 ; i < parameters.length ; i++) {
    var p = parameters[i];
    addSlider(tabDiv, p.Name, p.Title, p.Minimum, p.Maximum, p.Step, p.Default, p.Unit);
  }
}

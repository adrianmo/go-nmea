
// This variable tracks inflight requests:
// 'null' indicates no request is in flight or pending.
// 'false' indicates a request is inflight but no more are pending.
// 'true' indicates a request is inflight, but another request is pending.
var outstandingRequests = null;

var autoRefresh = null;

var updating = false;

var GetGraph = function() {
  if (outstandingRequests == null) {
    // No current inflight/outstanding requests.
    outstandingRequests = false;
  } else if (outstandingRequests == false) {
    // Current inflight, but none outstanding.
    outstandingRequests = true;
    return;
  } else {
    // Already inflight and outstanding;
    return;
  }
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
      if (outstandingRequests == true) {
        outstandingRequests = null;
        GetGraph();
      } else {
        outstandingRequests = null;
      }
      updateParameters($('#system-selector')[0].value);
      loadValues($('#system-selector')[0].value);
  };
  imgAttrs['x'] = (new Date).getTime();
  img.src = '/graph?' + $.param(imgAttrs)
  console.log('graph!');
}

updateOnChange = function() {
  if (!updating) {
    GetGraph();
  }
}

// Change a slider on mousewheel scroll event.
onMouseScroll = function(element, e) {
  var ev = e.originalEvent;
  var step = element.slider('option', 'step');
  var value = element.slider('value');
  if (ev.wheelDelta > 0) {
    value += step;
  } else {
    value -= step;
  }
  element.slider('option', 'value', value);
}

// Add a slider to the page.
addSlider = function(element, name, title, min, max, step, value, unit) {
  var sliderBox = document.createElement('div');
  sliderBox.className = 'slider-box';
  element.appendChild(sliderBox);

  var sliderHeader = document.createElement('div');
  sliderHeader.innerText = title;
  sliderHeader.className = 'slider-header';
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
    updateOnChange();
  });
  $('#' + name + '-slider').on("slidechange", function(event, ui) {
    $('#' + name + '-value').html(ui.value + unit);
    updateOnChange();
  });
  $('#' + name + '-slider').bind('mousewheel DOMMouseScroll', function(e) {
    onMouseScroll($(this), e);
  });

}

$(document).ready(function() {
  console.log("Ready!");
  $('.graph').html("<img src='/graph'/>");
  $('#control-tabs').tabs();
  $('#refresh').change(function(event) {
    updateRefresh(event.target.value);
  });
  $('#system-selector').change(function(event) {
    loadSystem(event.target.value);
  });
  loadSystem('');
  $('#reset').button();
  $('#reset').on('click', function() {
    if (confirm('Reset controller?')) {
      ResetAll();
    }
  });
  GetGraph();
});

loadSystem = function(name) {
  var allSystems = $.getJSON('/config', function(data) {
    clearSystem();
    // Add all the received ones.
    for (sys in data) {
      var system = data[sys];
      displaySystem(system, sys, name == '' || sys == name);
    }
    $('#control-tabs').tabs();
  });
}

ResetAll = function() {
  console.log('Resetting..');
  $.getJSON('/reset', function(data) {
    console.log('refresh..');
    location.reload();
  });
}

loadValues = function(name) {
  var allSystems = $.getJSON('/config', function(data) {
    for (sys in data) {
      var system = data[sys];
      displayValues(system, sys, name == '' || sys == name);
    }
  });
}

displayValues = function(system, name, selected) {
  var valBox = $('#values-box')[0];
  for (var i = valBox.children.length-1 ; i > -1 ; i--) {
    valBox.removeChild(valBox.children[i]);
  }
  for (var component in system.Values) {
    var vals  = system.Values[component];
    var compDiv = document.createElement('div');
    compDiv.className = 'value-component';
    valBox.appendChild(compDiv);

    var compLabel = document.createElement('div');
    compLabel.innerText = component;
    compLabel.className = 'value-label';
    compDiv.appendChild(compLabel);

    for (var value in vals) {
      var div = document.createElement('div');
      div.innerText = vals[value].Title;
      div.className = 'value-title';
      compDiv.appendChild(div);

      var div = document.createElement('div');
      div.innerText = vals[value].Value.toFixed(2) + vals[value].Unit;
      div.className = 'value-value';
      compDiv.appendChild(div);

    }
  }
}


clearSystem = function() {
  // Delete existing options.
  $('#control-tabs').tabs('destroy');
  $('.system-selector-option').remove()
  $('.control-tab').remove()
  $('.slider').slider('destroy');
  $('.sliders').remove()
}


displaySystem = function(system, name, selected) {
  var selector = new Option(system.Description, name);
  selector.className = 'system-selector-option';
  $('#system-selector')[0].add(selector);
  if (selected) {
    selector.selected = true;
    for (compName in system.Components) {
      var component = system.Components[compName];
      addTabAndParameters(compName, component);
    }
  }
}

updateRefresh = function(value) {
  console.log(value);
  if (value > 0) {
    autoRefresh = setInterval(GetGraph, value * 1000);
  } else {
    clearInterval(autoRefresh);
    autoRefresh = null;
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
    addSlider(tabDiv, p.Name, p.Title, p.Minimum, p.Maximum, p.Step, p.Value, p.Unit);
  }
}

updateSystemParameters = function(data) {
  updating = true;
  for (var c in data['Components']) {
    var component = data['Components'][c];
    for (var p in component) {
      var slider = '#' + component[p].Name + '-slider';
      var value = component[p].Value;
      $(slider).slider("value", value);
    }
  }
  updating = false;
}

updateParameters = function(system) {
  var allSystems = $.getJSON('/config', function(data) {
    updateSystemParameters(data[system]);
  });
}

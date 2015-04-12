var addButtons = function() {
  $.getJSON('/status', setupButtons);
}

var removeButtons = function() {
	$('.button-div').remove();
}

var updateButtons = function(data, status, jqxhdr) {
	for (var i in data.Regions) {
		var region = data.Regions[i]
		var btn = $('.' + region + '-button')[0]
		btn.setAttribute('current', data.Region == region)
	}
}

var buttonState = function(disable) {
	var btns = document.getElementsByClassName('country-button')
	for (var i = 0 ; i < btns.length ; i++) {
		btns[i].disabled = disable;
	}
}

var setupButtons = function(data, status, jqxhdr) {
	var clientInfo = $('#clientinfo')[0];
	clientInfo.innerText = 'Client address: ' + data.Client;
	var buttonBox = $('#buttons')[0];
	for (var i in data.Regions) {
		var region = data.Regions[i]
		var buttonDiv = document.createElement('div')
		buttonDiv.classList.add('button-div');
		buttonDiv.classList.add(region + '-button-div');
		buttonBox.appendChild(buttonDiv);
		var button = document.createElement('button')
		button.classList.add('country-button');
		button.classList.add(region + '-button');
		button.innerHTML = region;
		button.setAttribute('country', region);
		button.setAttribute('current', data.Region == region);
		buttonDiv.appendChild(button);
		$('.' + region + '-button').button({
			icons: {
				primary: region + "-flag"
			}
		})
		$('.' + region + '-button').on('click', function(ev) {
			buttonState(true);
			var btn = ev.currentTarget
			var country = btn.getAttribute('country')
			console.log('Clicked: ' + country);
			$('.status').css('visibility', 'visible');
			$.get('set', {country: country}, function(data, status, jqxhdr) {
				buttonState(false);
				$('.status').css('visibility', 'hidden');
				$.getJSON('/status', updateButtons);
			});
		});
	}
}

$(document).ready(function() {
  addButtons();
});

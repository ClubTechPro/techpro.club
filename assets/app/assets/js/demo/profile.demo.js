/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleGoogleMapSetting = function() {
	"use strict";
	var mapOptions = {
		zoom: 4,
		center: new google.maps.LatLng(37.09024, -95.712891),
		mapTypeId: google.maps.MapTypeId.ROADMAP,
		disableDefaultUI: true,
	};
	var mapDefault = new google.maps.Map(document.getElementById('google-map'), mapOptions);
};

var handleSuperboxSetting = function() {
	$(document).on('shown.bs.tab', function (e) {
		e.preventDefault();
		
		if ($(e.target).attr('href') == '#profile-photos') {
			$('#profile-photos').attr('data-init', true);
			$('.superbox').SuperBox({
				background : '#242a30',
				border : 'rgba(0,0,0,0.1)',
				xColor : '#a8acb1',
				xShadow : 'embed'
			});
		}
	});
};

var Profile = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleGoogleMapSetting();
			handleSuperboxSetting();
		}
	};
}();

$(document).ready(function() {
	Profile.init();
});
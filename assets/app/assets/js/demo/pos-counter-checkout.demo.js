/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleCheckTime = function(i) {
	"use strict";
	
	if (i < 10) {i = "0" + i};
	return i;
};
		
var handleStartTime = function() {
	"use strict";
	
	var today = new Date();
	var h = today.getHours();
	var m = today.getMinutes();
	var s = today.getSeconds();
	var a;
	m = handleCheckTime(m);
	s = handleCheckTime(s);
	a = (h > 11) ? 'pm' : 'am';
	h = (h > 12) ? h - 12 : h;
	document.getElementById('time').innerHTML = h + ":" + m + a;
	var t = setTimeout(handleStartTime, 500);
};

var handleSelectTable = function() {
	"use strict";
	
	$(document).on('click', '[data-toggle="select-table"]', function(e) {
		e.preventDefault();
		
		var targetTable = $(this).closest('.table');
		
		if ($(targetTable).hasClass('in-use')) {
			$('[data-toggle="select-table"]').not(this).closest('.table').removeClass('selected');
			$(targetTable).toggleClass('selected');
			$('#pos-counter').toggleClass('pos-mobile-sidebar-toggled');
		}
	});
};

var handleFilter = function() {
	"use strict";
	
	$(document).on('click', '.pos-menu [data-filter]', function(e) {
		e.preventDefault();
		
		var targetType = $(this).attr('data-filter');
		
		$(this).addClass('active');
		$('.pos-menu [data-filter]').not(this).removeClass('active');
		if (targetType == 'all') {
			$('.pos-content [data-type]').removeClass('d-none');
		} else {
			$('.pos-content [data-type="'+ targetType +'"]').removeClass('d-none');
			$('.pos-content [data-type]').not('.pos-content [data-type="'+ targetType +'"]').addClass('d-none');
		}
	});
};

$(document).ready(function() {
	handleStartTime();
	handleSelectTable();
	handleFilter();
});
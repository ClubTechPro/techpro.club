/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleDataTableSelect = function() {
	"use strict";
    
	if ($('#data-table-select').length !== 0) {
		$('#data-table-select').DataTable({
			select: true,
			responsive: true
		});
	}
};

var TableManageSelect = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleDataTableSelect();
		}
	};
}();

$(document).ready(function() {
	TableManageSelect.init();
});
/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleDataTableColReorder = function() {
	"use strict";
    
	if ($('#data-table-colreorder').length !== 0) {
		$('#data-table-colreorder').DataTable({
			colReorder: true,
			responsive: true
		});
	}
};

var TableManageColReorder = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleDataTableColReorder();
		}
	};
}();

$(document).ready(function() {
	TableManageColReorder.init();
});
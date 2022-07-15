/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleDataTableFixedHeader = function() {
	"use strict";
    
	if ($('#data-table-fixed-header').length !== 0) {
		$('#data-table-fixed-header').DataTable({
			lengthMenu: [20, 40, 60],
			fixedHeader: {
				header: true,
				headerOffset: $('#header').height()
			},
			responsive: true
		});
	}
};

var TableManageFixedHeader = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleDataTableFixedHeader();
		}
	};
}();

$(document).ready(function() {
	TableManageFixedHeader.init();
});
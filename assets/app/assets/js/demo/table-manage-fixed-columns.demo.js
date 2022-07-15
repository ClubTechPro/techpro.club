/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleDataTableFixedColumns = function() {
	"use strict";
    
	if ($('#data-table-fixed-columns').length !== 0) {
		$('#data-table-fixed-columns').DataTable({
			scrollY:        300,
			scrollX:        true,
			scrollCollapse: true,
			paging:         false,
			fixedColumns:   true
		});
	}
};

var TableManageFixedColumns = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleDataTableFixedColumns();
		}
	};
}();

$(document).ready(function() {
	TableManageFixedColumns.init();
});
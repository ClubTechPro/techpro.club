/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleJstreeDefault = function() {
	$("#jstree-default").jstree({
		"plugins": ["types"],
		"core": {
			"themes": {
				"responsive": false
			}            
		},
		"types": {
			"default": {
				"icon": "fa fa-folder text-warning fa-lg"
			},
			"file": {
				"icon": "fa fa-file text-inverse fa-lg"
			}
		}
	});

	$("#jstree-default").on("select_node.jstree", function(e,data) { 
		var link = $("#" + data.selected).find("a");
		if (link.attr("href") != "#" && link.attr("href") != "javascript:;" && link.attr("href") != "") {
			if (link.attr("target") == "_blank") {
				link.attr("href").target = "_blank";
			}
			document.location.href = link.attr("href");
			return false;
		}
	});
};


var handleJstreeCheckable = function() {
	$("#jstree-checkable").jstree({
		"plugins": ["wholerow", "checkbox", "types"],
		"core": {
			"themes": {
				"responsive": false
			},
			"data": [{
					"text": "Same but with checkboxes",
					"children": [{
						"text": "initially selected",
						"state": {
							"selected": true
						}
					}, {
						"text": "Folder 1"
					}, {
						"text": "Folder 2"
					}, {
						"text": "Folder 3"
					}, {
						"text": "initially open",
						"icon": "fa fa-folder fa-lg",
						"state": {
							"opened": true
						},
						"children": [{
							"text": "Another node"
						}, {
							"text": "disabled node",
							"state": {
								"disabled": true
							}
						}]
					}, {
						"text": "custom icon",
						"icon": "fa fa-cloud fa-lg text-blue"
					}, {
						"text": "disabled node",
						"state": {
							"disabled": true
						}
					}]
				},
				"Root node 2"
			]
		},
		"types": {
			"default": {
				"icon": "fa fa-folder text-primary fa-lg"
			},
			"file": {
				"icon": "fa fa-file text-success fa-lg"
			}
		}
	});
};


var handleJstreeDragAndDrop = function() {
  $("#jstree-drag-and-drop").jstree({
    "plugins": ["contextmenu", "dnd", "state", "types"],
    "core": {
      "themes": {
        "responsive": false
      },
      "check_callback": true,
      "data": [{
          "text": "Parent Node",
          "children": [{
            "text": "Initially selected",
            "state": {
              "selected": true
            }
          }, {
            "text": "Folder 1"
          }, {
            "text": "Folder 2"
          }, {
            "text": "Folder 3"
          }, {
            "text": "Initially open",
            "icon": "fa fa-folder text-success fa-lg",
            "state": {
              "opened": true
            },
            "children": [{
                "text": "Disabled node",
                "disabled": true
              },
              {
                "text": "Another node"
              }
            ]
          }, {
            "text": "Another Custom Icon",
            "icon": "fa fa-cog text-inverse fa-lg"
          }, {
            "text": "Disabled Node",
            "state": {
              "disabled": true
            }
          }, {
            "text": "Sub Nodes",
            "icon": "fa fa-folder text-danger fa-lg",
            "children": [{
                "text": "Item 1",
                "icon": "fa fa-file fa-lg"
              },
              {
                "text": "Item 2",
                "icon": "fa fa-file fa-lg"
              },
              {
                "text": "Item 3",
                "icon": "fa fa-file fa-lg"
              },
              {
                "text": "Item 4",
                "icon": "fa fa-file fa-lg"
              },
              {
                "text": "Item 5",
                "icon": "fa fa-file fa-lg"
              }
            ]
          }]
        },
        "Another Node"
      ]
    },
    "types": {
      "default": {
        "icon": "fa fa-folder text-warning fa-lg"
      },
      "file": {
        "icon": "fa fa-file text-warning fa-lg"
      }
    },
    "state": {
      "key": "demo2"
    }
  });
};


var handleJstreeAjax = function() {
  $("#jstree-ajax").jstree({
    "plugins": ["dnd", "state", "types"],
    "core": {
      "themes": {
        "responsive": false
      },
      "check_callback": true,
      "data": {
        "url": function(node) {
          return node.id === "#" ? "../assets/js/demo/json/data_root.json" : "../assets/js/demo/json/" + node.original.file;
        },
        "data": function(node) {
          return {
            "id": node.id
          };
        },
        "dataType": "json"
      }
    },
    "types": {
      "default": {
        "icon": "fa fa-folder text-warning fa-lg"
      },
      "file": {
        "icon": "fa fa-file text-warning fa-lg"
      }
    }
  });
};


var TreeView = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleJstreeDefault();
			handleJstreeCheckable();
			handleJstreeDragAndDrop();
			handleJstreeAjax();
		}
	};
}();

$(document).ready(function() {
	TreeView.init();
});
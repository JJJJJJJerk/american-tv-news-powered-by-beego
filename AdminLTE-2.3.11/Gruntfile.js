// AdminLTE Gruntfile
module.exports = function (grunt) {

  'use strict';

  grunt.initConfig({
    watch: {
      // If any .less file changes in directory "build/less/" run the "less"-task.
      files: ["build/less/*.less", "build/less/skins/*.less", "dist/js/app.js"],
      tasks: ["less", "uglify"]
    },
    // "less"-task configuration 编译less文件
    // This task will compile all less files upon saving to create both AdminLTE.css and AdminLTE.min.css
    less: {
      // Development not compressed
      pro: {
        options: {
          // Whether to compress or not
          compress: false
        },
        files: {
          // compilation.css  :  source.less
          "dist/css/AdminLTE.css": "build/less/AdminLTE.less",
          // AdminLTE without plugins
          "dist/css/alt/AdminLTE-without-plugins.css": "build/less/AdminLTE-without-plugins.less",
          // Separate plugins
          "dist/css/alt/AdminLTE-select2.css": "build/less/select2.less",
          "dist/css/alt/AdminLTE-fullcalendar.css": "build/less/fullcalendar.less",
          "dist/css/alt/AdminLTE-bootstrap-social.css": "build/less/bootstrap-social.less",
          //Non minified skin files
          "dist/css/skins/skin-blue.css": "build/less/skins/skin-blue.less",
          "dist/css/skins/skin-black.css": "build/less/skins/skin-black.less",
          "dist/css/skins/skin-yellow.css": "build/less/skins/skin-yellow.less",
          "dist/css/skins/skin-green.css": "build/less/skins/skin-green.less",
          "dist/css/skins/skin-red.css": "build/less/skins/skin-red.less",
          "dist/css/skins/skin-purple.css": "build/less/skins/skin-purple.less",
          "dist/css/skins/skin-blue-light.css": "build/less/skins/skin-blue-light.less",
          "dist/css/skins/skin-black-light.css": "build/less/skins/skin-black-light.less",
          "dist/css/skins/skin-yellow-light.css": "build/less/skins/skin-yellow-light.less",
          "dist/css/skins/skin-green-light.css": "build/less/skins/skin-green-light.less",
          "dist/css/skins/skin-red-light.css": "build/less/skins/skin-red-light.less",
          "dist/css/skins/skin-purple-light.css": "build/less/skins/skin-purple-light.less",
          "dist/css/skins/_all-skins.css": "build/less/skins/_all-skins.less"
        }
      }
    },
    //合并package css
    cssmin: {
      options: {
        mergeIntoShorthands: false,
        roundingPrecision: -1,
        //banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n',
        //美化代码
        beautify: {
            //中文ascii化，非常有用！防止中文乱码的神配置
            ascii_only: true
        }
      },
      target: {
        files: {
          '../static/css/all.min.css': [
                  "bootstrap/css/bootstrap.css",
                  //font awesome 放在这里
                  "dist/css/font-awesome_4.7.0.2.css",
                  //font ionicons 放在这里
                  "dist/css/ionicons_2.0.1.css",


                  "plugins/jvectormap/jquery-jvectormap-1.2.2.css",

                  "dist/css/AdminLTE.css",
                  // AdminLTE without plugins
                  //"dist/css/alt/AdminLTE-without-plugins.css",
                  // Separate plugins
                  "dist/css/alt/AdminLTE-select2.css",
                  "dist/css/alt/AdminLTE-fullcalendar.css",
                  "dist/css/alt/AdminLTE-bootstrap-social.css",
                  //Non minified skin files
                  "dist/css/skins/_all-skins.css"
                ]
        }
      }
  },
    // Uglify task info. Compress the js files.
    uglify: {
      options: {
        mangle: false,
        preserveComments: 'some'
      },
      my_target: {
        files: {
          //合并 jq bootstrap fastclick slimscroll
          '../static/js/all-package.min.js': ['plugins/jQuery/jquery-2.2.3.min.js',
                                              'bootstrap/js/bootstrap.js',
                                              'plugins/fastclick/fastclick.js',
                                              'dist/js/app.js',
                                              'plugins/sparkline/jquery.sparkline.min.js',
                                              'plugins/jvectormap/jquery-jvectormap-1.2.2.min.js',
                                              'plugins/jvectormap/jquery-jvectormap-world-mill-en.js',
                                              'plugins/slimScroll/jquery.slimscroll.min.js',
                                              'plugins/chartjs/Chart.min.js'

                                              
                                              
                                              
                                              ]
        }
      }
    },
    // Build the documentation files
    includes: {
      build: {
        src: ['*.html'], // Source files
        dest: 'documentation/', // Destination directory
        flatten: true,
        cwd: 'documentation/build',
        options: {
          silent: true,
          includePath: 'documentation/build/include'
        }
      }
    },

    // Optimize images
    image: {
      dynamic: {
        files: [{
          expand: true,
          cwd: 'build/img/',
          src: ['**/*.{png,jpg,gif,svg,jpeg}'],
          dest: 'dist/img/'
        }]
      }
    },

    // Validate JS code
    jshint: {
      options: {
        jshintrc: '.jshintrc'
      },
      core: {
        src: 'dist/js/app.js'
      },
      demo: {
        src: 'dist/js/demo.js'
      },
      pages: {
        src: 'dist/js/pages/*.js'
      }
    },

    // Validate CSS files
    csslint: {
      options: {
        csslintrc: 'build/less/.csslintrc'
      },
      dist: [
        'dist/css/AdminLTE.css'
      ]
    },

    // Validate Bootstrap HTML
    bootlint: {
      options: {
        relaxerror: ['W005']
      },
      files: ['pages/**/*.html', '*.html']
    },

    // Delete images in build directory
    // After compressing the images in the build/img dir, there is no need
    // for them
    clean: {
      build: ["build/img/*"]
    }
  });

  // Load all grunt tasks

  // LESS Compiler
  grunt.loadNpmTasks('grunt-contrib-less');
  // Watch File Changes
  grunt.loadNpmTasks('grunt-contrib-watch');
  // Compress JS Files
  grunt.loadNpmTasks('grunt-contrib-uglify');
  // Include Files Within HTML
  grunt.loadNpmTasks('grunt-includes');
  // Optimize images
  grunt.loadNpmTasks('grunt-image');
  // Validate JS code
  grunt.loadNpmTasks('grunt-contrib-jshint');
  // Delete not needed files
  grunt.loadNpmTasks('grunt-contrib-clean');
  // Lint CSS
  grunt.loadNpmTasks('grunt-contrib-csslint');
  //concat 合并css
  grunt.loadNpmTasks('grunt-contrib-cssmin');
  // Lint Bootstrap
  grunt.loadNpmTasks('grunt-bootlint');

  // Linting task
  //grunt.registerTask('lint', ['jshint', 'csslint', 'bootlint']);

  // The default task (running "grunt" in console) is "watch"
  //grunt.registerTask('default', ['watch']);
  grunt.registerTask('all', ['less','cssmin','uglify']);
};

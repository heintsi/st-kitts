'use strict';

module.exports = function(grunt) {

  // Load grunt tasks automatically
  require('load-grunt-tasks')(grunt);

  require('time-grunt')(grunt);

  grunt.initConfig({

    app: {
      src: 'src',
      dist: 'dist',
      scripts: '<%= app.src %>/scripts',
      images: '<%= app.src %>/images',
      styles: '<%= app.src %>/styles',
      tests: '<%= app.src %>/test'
    },

    copy: {
      dist: {
        files: [{
          expand: true,
          dot: true,
          cwd: '<%= app.src %>',
          dest: '<%= app.dist %>',
          src: [
            'scripts/st-kitts.js',
            '*.{ico,png,txt}',
            'lib/*/*.js',
            '{,*/}*.html',
            'styles/*.css'
          ]
        }]
      }
    },

    // + MOCHA
    // + PHANTOMJS

    jshint: {
      files: [
        'Gruntfile.js',
        '<%= app.scripts %>/*.js',
        '<%= app.tests %>/*.js'
      ],
      options: {
        jshintrc: '.jshintrc',
        reporter: require('jshint-stylish')
      }
    },

    watch: {
      options: {
        livereload: true
      },
      files: ['<%= jshint.files %>'],
      tasks: ['jshint']
    },

    // Empties folders to start fresh
    clean: {
      dist: {
        files: [{
          dot: true,
          src: [
            '<%= app.dist %>/*',
            '!<%= app.dist %>/.git*'
          ]
        }]
      }
    },

    // The actual grunt server settings
    connect: {
      options: {
        port: 9000,
        livereload: 35729,
        // Change this to '0.0.0.0' to access the server from outside
        hostname: 'localhost'
      },
      livereload: {
        options: {
          open: true,
          base: [
            '<%= app.src %>'
          ]
        }
      },
      test: {
        options: {
          port: 9001,
          base: [
            '<%= app.tests %>',
            '<%= app.src %>'
          ]
        }
      },
      dist: {
        options: {
          open: true,
          base: '<%= app.dist %>',
          livereload: false
        }
      }
    }

  });

  // TASKS

  grunt.registerTask('inspect', [
    'jshint'
  ]);

  grunt.registerTask('test', [
    // TODO 'mocha'
  ]);

  grunt.registerTask('build', [
    'clean',
    'copy'
  ]);

  grunt.registerTask('run', [
    'connect:livereload',
    'watch'
  ]);

  grunt.registerTask('prod', [
    'inspect',
    'test',
    'build',
    'connect:dist:keepalive'
  ]);

  grunt.registerTask('default', [
    'inspect',
    'test',
    'build'
  ]);
};
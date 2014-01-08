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
      dev: {
        files: [{
          expand: true,
          dot: true,
          cwd: '<%= app.src %>',
          dest: '<%= app.dist %>',
          src: [
            'scripts/st-kitts.js',
            '{,*/}*.html',
            'styles/*.css'
          ]
        }]
      },
      lib: {
        files: [{
          expand: true,
          dot: false,
          cwd: '<%= app.src %>',
          dest: '<%= app.dist %>',
          src: [
            'lib/requirejs/require.js',
            'lib/jquery/jquery.js',
            'lib/underscore/underscore.js'
          ]
        }]
      }
    },

    sync: {
      dev: {
        files: [{
          cwd: 'src/',
          src: [
            'scripts/*.js',
            'index.html'
          ],
          dest: 'dist'
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
      sync: {
        files: [
          'scripts/*.js',
          '{,*/}*.html',
          'styles/*.css'
        ],
        tasks: ['sync:dev']
      }
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
        // Change this to '0.0.0.0' to access the server from outside
        hostname: 'localhost'
      },
      dev: {
        options: {
          open: true,
          base: [
            '<%= app.dist %>'
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
    'default',
    'connect:dev',
    'watch'
  ]);

  grunt.registerTask('prod', [
    'default',
    'connect:dev:keepalive'
  ]);

  grunt.registerTask('default', [
    'inspect',
    'test',
    'build'
  ]);
};
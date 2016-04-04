'use strict';

var gulp = require('gulp');
var $ = require('gulp-load-plugins')();
var del = require('del');
var runSequence = require('run-sequence');
var vulcanize = require('gulp-vulcanize');
var connect = require('gulp-connect');
var merge = require('merge-stream');
var path = require('path');

var DIST = 'www';

var dist = function(subpath) {
  return !subpath ? DIST : path.join(DIST, subpath);
};

gulp.task('devcopy', function() {
  var app = gulp.src([
    'app/**/*',
    '!app/components'
  ], {
    dot: true
  }).pipe(gulp.dest(dist()));

  // For dev build, copy all the items to enable debugging
  var bower = gulp.src([
    'app/components/**/*'
  ]).pipe(gulp.dest(dist('components')));

  return merge(app, bower)
    .pipe($.size({
      title: 'copy'
    }));
});

// Clean output directory
gulp.task('clean', function() {
  return del(['www', dist()]);
});


gulp.task('vulcanize', function() {
  return gulp.src('elements.html')
    .pipe(vulcanize({
      dest: 'dist',
      strip: true
    }))
    .pipe(gulp.dest('dist'));
});

gulp.task('connect', function() {
  connect.server({
    root: 'www',
    livereload: true
  });
});

gulp.task('serve', ['connect']);

// For Dev build no need to vulcanize. Just copy the required bower components to dist/www directory
gulp.task('devbuild', ['clean'], function() {
  runSequence(['devcopy']);
});


// TODO All items below this belongs to prod build

gulp.task('prodcopy', function() {
  var app = gulp.src([
    'app/**/*',
    '!app/components'
  ], {
    dot: true
  }).pipe(gulp.dest(dist()));

  // For prod, Copy over only the components we need
  // These are things which cannot be vulcanized
  // But for dev copy all the items to enable debugging
  var bower = gulp.src([
    'app/components/{webcomponentsjs,platinum-sw,sw-toolbox,promise-polyfill}/**/*'
  ]).pipe(gulp.dest(dist('components')));

  return merge(app, bower)
    .pipe($.size({
      title: 'copy'
    }));
});


gulp.task('prodbuild', ['vulcanize']);
